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
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"net"
	"strings"
)

type Client struct {
	Uri  string
	Conn *net.Conn

	PublicKey *rsa.PublicKey
	PemKey    string
}

func (client *Client) Connect() {
	fmt.Println("[client - info] Attempting to connect to socket at: ", client.Uri)

	conn, err := net.Dial("tcp", client.Uri)
	if err != nil {
		panic(err)
	}

	client.Conn = &conn
}

func (client *Client) NegotiateKeys() {
	conn := *client.Conn

	_, err := conn.Write([]byte("CONNECT"))
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 4096)
	n, _err := conn.Read(buffer)
	if _err != nil {
		panic(_err)
	}

	bufferSplit := strings.Split(string(buffer[:n]), "PUBKEY:")
	pemKey := bufferSplit[1]

	client.PemKey = pemKey

	pubKeyBlock, _ := pem.Decode([]byte(pemKey)) // split buffer
	pubKey, err := x509.ParsePKCS1PublicKey(pubKeyBlock.Bytes)
	if err != nil {
		panic(err)
	}

	client.PublicKey = pubKey

	fmt.Println("[info - client] ", string(buffer))
}

func (client *Client) ValidatePublicKey() {
	hash := sha256.Sum256([]byte(client.PemKey))
	hashStr := hex.EncodeToString(hash[:])

	fmt.Println("Key Pair", hashStr)
	message := "PUBKEY:ACK:" + hashStr

	conn := *client.Conn

	_, err := conn.Write([]byte(message))
	if err != nil {
		panic(err)
	}

}

func (client *Client) SendWelcome() {
	fmt.Println("[client - msg] Sending welcome message")

	conn := *client.Conn

	msg := []byte("Hello from client")

	cipherText, _err := rsa.EncryptOAEP(sha256.New(), rand.Reader, client.PublicKey, msg, nil)
	if _err != nil {
		panic(_err)
	}

	encodedText := base64.StdEncoding.WithPadding(base64.StdPadding).EncodeToString(cipherText)
	_, err := conn.Write([]byte(encodedText))
	if err != nil {
		fmt.Println("[error] Failed to send welcome message: ", err.Error())
	}
}

func (client *Client) Disconnect() {
	fmt.Println("[client - info] Attempting to disconnect from socket")

	conn := *client.Conn
	conn.Close()
}
