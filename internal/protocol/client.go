package protocol

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"

	"github.com/Rhisiart/battleship/internal/matchmaking"
)

type Client struct {
	*matchmaking.Player
	conn     net.Conn
	outbound chan<- Command
	login    chan<- *Client
	logout   chan<- *Client
}

func NewClient(
	conn net.Conn,
	o chan<- Command,
	lin chan<- *Client,
	lout chan<- *Client) *Client {
	return &Client{
		conn:     conn,
		outbound: o,
		login:    lin,
		logout:   lout,
	}
}

func (p *Client) Read() error {
	for {
		buffer, err := bufio.NewReader(p.conn).ReadBytes('\n')

		if err == io.EOF {
			p.logout <- p
			return nil
		}

		if err != nil {
			return err
		}

		p.handleMessage(buffer)
	}
}

func (c *Client) handleMessage(message []byte) {
	cmd := bytes.ToUpper(bytes.TrimSpace(bytes.Split(message, []byte(" "))[0]))
	args := bytes.TrimSpace(bytes.TrimPrefix(message, cmd))

	switch string(cmd) {
	case "LOGIN":
		if err := c.logIn(args); err != nil {
			c.err(err)
		}
	case "LOGOUT":
		if err := c.logOut(); err != nil {
			c.err(err)
		}
	case "CLIENTS":
		c.numberOfClients()
	case "JOIN":
		c.join()
	default:
		c.err(fmt.Errorf("unknown command %s", cmd))
	}
}

func (c *Client) logIn(args []byte) error {
	u := bytes.TrimSpace(args)

	if len(u) == 0 {
		return fmt.Errorf("username cannot be blank")
	}

	c.Player = matchmaking.NewPlayer(string(u))
	c.login <- c

	return nil
}

func (c *Client) logOut() error {
	c.logout <- c

	return nil
}

func (c *Client) join() {
	c.outbound <- Command{
		sender: c.Player.Username,
		id:     JOIN,
	}
}

func (c *Client) numberOfClients() {
	c.outbound <- Command{
		sender: c.Player.Username,
		id:     CLIENTS,
	}
}

func (c *Client) err(e error) {
	c.conn.Write([]byte(fmt.Sprintf("ERR %s \n", e.Error())))
}
