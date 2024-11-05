package src

import (
	"fmt"
	"net"
)

type Client struct {
	Uri  string
	Conn *net.Conn
}

func (client *Client) Connect() {
	fmt.Println("[client - info] Attempting to connect to socket at: ", client.Uri)

	conn, err := net.Dial("tcp", client.Uri)
	if err != nil {
		panic(err)
	}

	client.Conn = &conn
}

func (client *Client) SendWelcome() {
	fmt.Println("[client - msg] Sending welcome message")

	msg := []byte("Hello from client")

	conn := *client.Conn

	_, err := conn.Write(msg)
	if err != nil {
		fmt.Println("[error] Failed to send welcome message: ", err.Error())
	}
}

func (client *Client) Disconnect() {
	fmt.Println("[client - info] Attempting to disconnect from socket")

	conn := *client.Conn
	conn.Close()
}
