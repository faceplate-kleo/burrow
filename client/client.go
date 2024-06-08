package client

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"time"
)

type Client struct{}

func (c *Client) Run(wg *sync.WaitGroup) {
	for {
		time.Sleep(2 * time.Second)
		conn, err := net.Dial("tcp", "localhost:8088")
		if err != nil {
			fmt.Printf("failed to dial: %s\n", err)
			continue
		}

		_, err = fmt.Fprintf(conn, "CLIENT REQUEST\n")
		if err != nil {
			fmt.Printf("failed to send: %s\n", err)
			continue
		}
		resp, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("failed to read server response: %s\n", err)
			continue
		}

		fmt.Printf("Client got message: %s\n", resp)
	}
	wg.Done()
}
