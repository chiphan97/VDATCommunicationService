package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var broker *TestBroker

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	To     string `json:"to"`
	From   string `json:"from"`
	Body   string `json:"body"`
	status string
}

type TestBroker struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	inbound chan Message

	// Outbound messages that need to send to clients.
	outbound chan Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	messageRepository []*Message
}

func newbroker() *TestBroker {
	return &TestBroker{
		inbound:    make(chan Message),
		outbound:   make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (b *TestBroker) run() {
	// polling "new" message from repository
	// and send to outbound channel to send to clients
	// finally, marked message that sent to outbound channel as "done"
	go func() {
		for {
			for idx, m := range b.messageRepository {
				if m.status != "done" {
					select {
					case b.outbound <- *m:
					default:
						close(b.outbound)
					}

					b.messageRepository[idx].status = "done"
				}
			}

			time.Sleep(200 * time.Millisecond)
		}
	}()

	for {
		select {
		case client := <-b.register:
			b.clients[client] = true
		case client := <-b.unregister:
			if _, ok := b.clients[client]; ok {
				delete(b.clients, client)
				close(client.send)
			}
		case message := <-b.inbound:

			b.messageRepository = append(b.messageRepository, &message)
			fmt.Printf("%+v, %d\n", message, len(b.messageRepository))

		case message := <-b.outbound:
			fmt.Println("send")
			for client := range b.clients {
				if client.userID == message.To {
					msg, _ := json.Marshal(message)
					select {
					case client.send <- msg:
					default:
						close(client.send)
						delete(b.clients, client)
					}
				}
			}
		}

	}
}

type Client struct {
	broker *TestBroker

	// user ID, this will be parse from Access Token in production
	userID string

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

func (c *Client) readPump() {
	defer func() {
		c.broker.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

		var messageJSON Message
		_ = json.Unmarshal(message, &messageJSON)
		messageJSON.From = c.userID

		c.broker.inbound <- messageJSON
	}
}

// writePump pumps messages from the broker to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The broker closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	if broker == nil {
		broker = &TestBroker{
			inbound:    make(chan Message),
			outbound:   make(chan Message),
			register:   make(chan *Client),
			unregister: make(chan *Client),
			clients:    make(map[*Client]bool),
		}

		go broker.run()
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(r.URL.Query()["userID"][0])

	client := &Client{userID: r.URL.Query()["userID"][0], broker: broker, conn: conn, send: make(chan []byte, 256)}
	client.broker.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
