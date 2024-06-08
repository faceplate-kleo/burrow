package main

import (
	"github.com/faceplate-kleo/burrow/client"
	"github.com/faceplate-kleo/burrow/server"
	"sync"
)

func main() {

	c := client.Client{}
	s := server.Server{}

	var wg sync.WaitGroup

	wg.Add(2)

	go c.Run(&wg)
	go s.Run(&wg)

	wg.Wait()

	println("Hello world")
}
