package main

import (
	"github.com/faceplate-kleo/burrow/server"
	"sync"
)

func main() {
	var s server.Server

	var wg sync.WaitGroup
	wg.Add(1)
	s.Run(&wg)
}
