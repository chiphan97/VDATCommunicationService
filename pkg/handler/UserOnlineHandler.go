package handler

import (
	"bytes"
	_ "bytes"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/service"
	"log"
	"net/http"
	"strings"
	"time"
)

var userOnlineBroker *WsBroker

func (broker *WsBroker) run() {
	// polling "new" message from repository
	// and send to outbound channel to send to clients
	// finally, marked message that sent to outbound channel as "done"
	go func() {
		for {
			for idx, message := range broker.MessageRepository {
				if message.Status != "done" {
					select {
					case broker.Outbound <- *message:
					default:
						close(broker.Outbound)
					}

					broker.MessageRepository[idx].Status = "done"
				}
			}

			time.Sleep(200 * time.Millisecond)
		}
	}()

	for {
		select {
		case client := <-broker.Register:
			broker.Clients[client] = true
			//add in database when client on
			_ = service.AddUserOnlineService(client.User)
		case client := <-broker.Unregister:
			if _, ok := broker.Clients[client]; ok {
				//delete in database when client off
				_ = service.DeleteUserOnlineService(client.User.SocketID)
				delete(broker.Clients, client)
				close(client.Send)
			}
		case message := <-broker.Inbound:
			broker.MessageRepository = append(broker.MessageRepository, &message)
			fmt.Printf("%+v, %d\n", message, len(broker.MessageRepository))
		case message := <-broker.Outbound:
			fmt.Println("send")
			for client := range broker.Clients {
				msg, _ := json.Marshal(message)
				select {
				case client.Send <- msg:
				default:
					close(client.Send)
					delete(broker.Clients, client)
				}
			}
		}

	}
}

func (client *WsClient) readPump() {
	defer func() {
		client.Broker.Unregister <- client
		_ = client.Conn.Close()
	}()
	client.Conn.SetReadLimit(MaxMessageSize)
	_ = client.Conn.SetReadDeadline(time.Now().Add(PongWait))
	client.Conn.SetPongHandler(func(string) error { _ = client.Conn.SetReadDeadline(time.Now().Add(PongWait)); return nil })
	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, Newline, Space, -1))

		var messageJSON WsMessage
		_ = json.Unmarshal(message, &messageJSON)
		messageJSON.From = client.User.UserID

		client.Broker.Inbound <- messageJSON
	}
}

func (client *WsClient) checkUserOnlinePump() {
	defer func() {
		client.Broker.Unregister <- client
		_ = client.Conn.Close()
	}()

	for {
		usersOnline, _ := service.GetListUSerOnlineService()

		message := WsMessage{
			From:   "VDAT-SERVICE",
			To:     nil,
			Body:   usersOnline,
			Status: "",
		}

		client.Broker.Inbound <- message
		time.Sleep(10000 * time.Millisecond)
	}
}

func (client *WsClient) writePump() {
	ticker := time.NewTicker(PingPeriod)

	defer func() {
		ticker.Stop()
		_ = client.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-client.Send:
			_ = client.Conn.SetWriteDeadline(time.Now().Add(WriteWait))
			if !ok {
				// The broker closed the channel.
				_ = client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := client.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			_, _ = w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(client.Send)
			for i := 0; i < n; i++ {
				_, _ = w.Write(Newline)
				_, _ = w.Write(<-client.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			_ = client.Conn.SetWriteDeadline(time.Now().Add(WriteWait))
			if err := client.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func UserOnlineHandler(w http.ResponseWriter, r *http.Request) {
	// authenticate
	var (
		claims model.UserClaims

		// https://accounts.vdatlab.com/auth/realms/vdatlab.com/protocol/openid-connect/certs
		publicKey = "MIICpTCCAY0CBgFrPLdvYjANBgkqhkiG9w0BAQsFADAWMRQwEgYDVQQDDAt2ZGF0bGFiLmNvbTAeFw0xOTA2MDkxNDQ4MDNaFw0yOTA2MDkxNDQ5NDNaMBYxFDASBgNVBAMMC3ZkYXRsYWIuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAleIkHoO6Q0GRQ4POIAKmN5Ev3zfAm8raTJQ1e/CbTXW4FQ0kDS9YPhLXPwcdnbxiL3rSGgz7+iWcq/Ix7yExuNbSyqDUjLUJSU6I9JvB1YP8GSaO8d996+TVCDC8E/VSID6wmfWbMNb5Ns6Y7YY/HAhj9zc73ObErvi0NV0BjeYAVOBqJKKgl9cHfyBshr+kpC/7nrbTRnAP7JQhKrQF6wBTKQiuJlEyYqvi1ugCRBYg2BZLPtTry+Kineb1DT8ynmxJjKMtr9hU0dsLPJpqW/4DWwNOarLOBP/K9WkfR2LUxbrm41goSTjJbz6s7f/Mvn/gDLjGjIsdlFP3Y7I2lwIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQBXU5Awwhv/cYJKCdSUzmtXpXty8KrdrHaDNa8potDXlEc2JrK3wHyFRwwfpBhkaicP0LllxRHUGUNsWFnggae1fudc75fysZ16NPH7VJlUuyV96K06K4v1aM5VCSWl5djky7rtyfi2W9iH2ddWZvCeSyFsSgCD4P5GjgYpsLy27g/cvdJJAdp/b7bweVDI1grlBtnInxLUPhJ4cnoNw3crh7twqKgG6F3GmZc2Hjl45LdlxBFfftDUYH66D1X0mdoipQCbg4JWlIxUZHVjJDIrSIlwnRMwjzCm7MUYv0ySmvsxgoNVI2NuFU6A/F7zlyVkDkmO4ilp4BueRtBKb7yR"
	)
	{
		var tok string
		queryToken := r.URL.Query()["token"]
		if len(queryToken) > 0 {
			tok = queryToken[0]
		}

		var accessToken string
		h := r.Header.Get("Authorization")
		a := strings.Split(h, " ")
		if len(a) == 2 {
			accessToken = a[1]
		} else if tok != "" {
			accessToken = tok
		} else {
			w.WriteHeader(400)
			return
		}

		token, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
			return jwt.ParseRSAPublicKeyFromPEM([]byte("-----BEGIN CERTIFICATE-----\n" + publicKey + "\n-----END CERTIFICATE-----"))
		})

		if err != nil {
			fmt.Println("CANNOT parse token: ", err)
			w.WriteHeader(401)
			return
		}

		if token.Valid {
			fmt.Println("Token is VALID")
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				fmt.Println("That's not even a token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				fmt.Println("Timing is everything")
			} else {
				fmt.Println("Couldn't handle this token:", err)
			}
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	}

	if userOnlineBroker == nil {
		userOnlineBroker = &WsBroker{
			Inbound:    make(chan WsMessage),
			Outbound:   make(chan WsMessage),
			Register:   make(chan *WsClient),
			Unregister: make(chan *WsClient),
			Clients:    make(map[*WsClient]bool),
		}

		go userOnlineBroker.run()
	}

	conn, err := WsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	user := model.UserOnline{
		HostName: r.URL.Path,
		SocketID: claims.Subject,
		UserID:   claims.Subject,
		Username: claims.UserName,
		First:    claims.GivenName,
		Last:     claims.FamilyName,
	}

	client := &WsClient{User: user, Broker: userOnlineBroker, Conn: conn, Send: make(chan []byte, 256)}
	client.Broker.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
	go client.checkUserOnlinePump()
}
