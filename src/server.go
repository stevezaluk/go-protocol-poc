package src

import (
	"fmt"
	"net"
)

type Server struct {
	Sock            *net.Listener
	ConnectionCount int
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

func (server *Server) AcceptConnections() {
	fmt.Println("[server - info] Socket now open for connections")
	for {
		sock := *server.Sock
		conn, err := sock.Accept()
		if err != nil {
			fmt.Println("[server - error] Issue accepting connection: ", err.Error())
			continue
		}

		go server.ConnectionHandler(conn)
	}
}

func (server *Server) ConnectionHandler(conn net.Conn) {
	buf := make([]byte, 4096)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("[server - error] Issue receiving data: ", err.Error())
			return
		}

		fmt.Println("[server - msg] Message from Client: ", string(buf))
	}
}

func (server *Server) Stop() {
	fmt.Println("[server - info] Attempting to stop socket")
	serv := *server.Sock
	serv.Close()
}
