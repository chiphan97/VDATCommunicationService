package handler

type WsBroker struct {
	// Registered clients.
	Clients map[*WsClient]bool

	// Inbound messages from the clients.
	Inbound chan WsMessage

	// Outbound messages that need to send to clients.
	Outbound chan WsMessage

	// Register requests from the clients.
	Register chan *WsClient

	// Unregister requests from clients.
	Unregister chan *WsClient

	MessageRepository []*WsMessage
}
