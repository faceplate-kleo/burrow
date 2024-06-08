package room

import (
	"fmt"
	"net"
)

type Room struct {
	connections []*net.Conn
	users       []string
}

func (r *Room) AddConnection(conn *net.Conn) {
	r.connections = append(r.connections, conn)
}

func (r *Room) AddUser(usr string) {
	r.users = append(r.users, usr)
}

func (r *Room) GetUserList() string {
	out := ""
	delim := "\n"
	for i, usr := range r.users {
		if i == len(r.users)-1 {
			delim = ""
		}
		out += fmt.Sprintf("%s%s", usr, delim)
	}
	return out
}
