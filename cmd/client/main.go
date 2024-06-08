package main

import (
	"github.com/faceplate-kleo/burrow/client"
	"sync"
)

func main() {
	var c client.Client

	var wg sync.WaitGroup
	wg.Add(1)
	c.Run(&wg)
}
