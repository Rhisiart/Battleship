package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Rhisiart/battleship/internal/protocol"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Printf("%v", err)
	}

	lobby := protocol.NewHub()

	go lobby.Run()

	fmt.Printf("Lobby created\n")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("%v", err)
			continue
		}

		c := protocol.NewClient(
			conn,
			lobby.Commands,
			lobby.Connected,
			lobby.Disconnected,
		)

		fmt.Printf("New Player connected\n")

		go c.Read()
	}
}
