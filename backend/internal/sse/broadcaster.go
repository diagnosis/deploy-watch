package sse

import "github.com/google/uuid"

type Client struct {
	UserID uuid.UUID   `json:"user_id"`
	Send   chan string `json:"send"`
}

type UserMessage struct {
	UserID  uuid.UUID
	Message string
}

type Broadcaster struct {
	clients         map[*Client]bool
	register        chan *Client
	unregister      chan *Client
	broadcast       chan string
	broadcastToUser chan UserMessage
}

func NewBroadcaster() *Broadcaster {
	return &Broadcaster{
		clients:         make(map[*Client]bool),
		register:        make(chan *Client),
		unregister:      make(chan *Client),
		broadcast:       make(chan string, 100),
		broadcastToUser: make(chan UserMessage, 100),
	}
}

func (b *Broadcaster) Run() {
	for {
		select {
		case client := <-b.register:
			b.clients[client] = true
		case client := <-b.unregister:
			if _, ok := b.clients[client]; ok {
				delete(b.clients, client)
				close(client.Send)
			}
		case msg := <-b.broadcast:
			for client := range b.clients {
				select {
				case client.Send <- msg:
				default:
				}
			}
		case userMsg := <-b.broadcastToUser:
			for client := range b.clients {
				if client.UserID == userMsg.UserID {
					select {
					case client.Send <- userMsg.Message:
					default:

					}
				}
			}
		}

	}
}

func (b *Broadcaster) Register(client *Client) {
	b.register <- client
}

func (b *Broadcaster) Unregister(client *Client) {
	b.unregister <- client
}

func (b *Broadcaster) Broadcast(msg string) {
	b.broadcast <- msg
}

func (b *Broadcaster) BroadcastToUser(userID uuid.UUID, msg string) {
	b.broadcastToUser <- UserMessage{
		UserID:  userID,
		Message: msg,
	}
}
