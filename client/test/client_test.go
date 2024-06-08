package test

import (
	"github.com/faceplate-kleo/burrow/client"
	"testing"
)

func TestBasicRun(t *testing.T) {
	c := client.Client{}

	c.Run(nil)
}
