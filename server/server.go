package server

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

type Server struct{}

func (c *Server) Run(wg *sync.WaitGroup) {
	ln, err := net.Listen("tcp", ":8088")
	if err != nil {
		fmt.Printf("failed to listen: %s", err)
		wg.Done()
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("failed to accept: %s\n", err)
		}

		fmt.Printf("Accepted connection from %s | %s\n", conn.RemoteAddr(), conn.LocalAddr())

		resp, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("server failed to read: %s\n", err)
			continue
		}

		fmt.Println(resp)

		fmt.Printf("Server got message: %s\n", resp)

		_, err = fmt.Fprintf(conn, "SERVER ACK\n")
		if err != nil {
			fmt.Printf("failed to send ack: %s", err)
		}
	}
}
