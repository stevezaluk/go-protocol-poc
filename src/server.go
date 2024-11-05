package src

import (
	"fmt"
	"net"
)

type Server struct {
	Sock *net.Listener
}

func (server *Server) Start() {
	fmt.Println("[server - info] Attempting to start socket")
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("[server - error] Starting connection failed")
		panic(err)
	}

	server.Sock = &listener
}

func (server *Server) Stop() {
	fmt.Println("[server - info] Attempting to stop socket")
	serv := *server.Sock
	serv.Close()
}
