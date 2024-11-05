/*
Copyright © 2024 Steven A. Zaluk

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
