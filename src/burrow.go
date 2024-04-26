package main

import (
	"github.com/faceplate-kleo/burrow/client"
	"github.com/faceplate-kleo/burrow/server"
)

func main() {

	c := client.Client{}
	s := server.Server{}

	c.Run()
	s.Run()

	println("Hello world")
}
