package tcpclient

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type TCPClient struct {
	conn net.Conn
}

func New() *TCPClient {
	return &TCPClient{}
}

func (c *TCPClient) ReadMessages(address string) *bufio.Reader {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	fmt.Println("TCP connection successful.")

	return bufio.NewReader(conn)
}

func (c *TCPClient) Disconnect() {
	c.conn.Close()
}
