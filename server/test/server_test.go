package server

import (
	"github.com/faceplate-kleo/burrow/server"
	"testing"
)

func TestBasicRun(t *testing.T) {
	var s server.Server

	s.Run(nil)
}
