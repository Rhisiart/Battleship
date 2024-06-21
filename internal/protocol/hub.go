package protocol

import (
	"fmt"
)

type Hub struct {
	Clients      map[string]*Client
	Commands     chan Command
	Disconnected chan *Client
	Connected    chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Connected:    make(chan *Client),
		Disconnected: make(chan *Client),
		Clients:      make(map[string]*Client),
		Commands:     make(chan Command),
	}
}

func (h *Hub) Run() {
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
			}
		}
	}
}

func (h *Hub) connection(c *Client) {
	if _, exist := h.Clients[c.username]; exist {
		c.username = ""
		c.conn.Write([]byte("Err username taken\n"))
	} else {
		h.Clients[c.username] = c
		c.conn.Write([]byte(fmt.Sprintf("%s connected", c.username)))
	}
}

func (h *Hub) disconnect(c *Client) {
	delete(h.Clients, c.username)
}

func (h *Hub) numberOfClients(sender string) {
	if c, ok := h.Clients[sender]; ok {
		c.conn.Write([]byte(fmt.Sprintf("%d Clients connected", len(h.Clients))))
	}
}
