package protocol

import (
	"fmt"

	"github.com/Rhisiart/battleship/internal/matchmaking"
)

type Hub struct {
	Queue        *matchmaking.Queue
	Clients      map[string]*Client
	Commands     chan Command
	Disconnected chan *Client
	Connected    chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Queue:        matchmaking.NewQueue(2),
		Connected:    make(chan *Client),
		Disconnected: make(chan *Client),
		Clients:      make(map[string]*Client),
		Commands:     make(chan Command),
	}
}

func (h *Hub) Run() {
	h.Queue.StartMatch()

	for {
		select {
		case client := <-h.Connected:
			h.connection(client)
		case client := <-h.Disconnected:
			h.disconnect(client)
		case cmd := <-h.Commands:
			switch cmd.id {
			case CLIENTS:
				h.numberOfClients(cmd.sender)
			case JOIN:
				h.join(cmd.sender)
			}
		}
	}
}

func (h *Hub) connection(c *Client) {
	if _, exist := h.Clients[c.Player.Username]; exist {
		c.Player.Username = ""
		c.conn.Write([]byte("Err username taken\n"))
	} else {
		h.Clients[c.Player.Username] = c
		c.conn.Write([]byte(fmt.Sprintf("%s connected", c.Player.Username)))
	}
}

func (h *Hub) disconnect(c *Client) {
	delete(h.Clients, c.Player.Username)
}

func (h *Hub) join(sender string) {
	fmt.Println("Trying to join")
	if c, ok := h.Clients[sender]; ok {
		h.Queue.Enqueue(c.Player)

		c.conn.Write([]byte(fmt.Sprintf("%s You join a queue", sender)))
	}
}

func (h *Hub) numberOfClients(sender string) {
	if c, ok := h.Clients[sender]; ok {
		c.conn.Write([]byte(fmt.Sprintf("%d Clients connected", len(h.Clients))))
	}
}
