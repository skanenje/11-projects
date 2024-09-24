package client

import (
	"bufio"
	"fmt"
	"os"
	"net/pkg/chat"

	"net"
)

type Client struct {
	conn   net.Conn
	name   string
	reader *bufio.Reader
	writer *bufio.Writer
}

func NewClient(ip, port string) *Client {
	return &Client{}
}

func (c *Client) Run() error {
	conn, err := net.Dial("tcp", c.getAddress())
	if err != nil {
		return err
	}
	c.conn = conn

	c.reader = bufio.NewReader(conn)
	c.writer = bufio.NewWriter(conn)

	welcome, _ := c.reader.ReadString('\n')
	fmt.Print(welcome)

	c.name = c.getName()
	fmt.Fprintf(c.writer, "%s\n", c.name)
	c.writer.Flush()

	go c.receiveMessages()
	c.sendMessages()

	return nil
}

func (c *Client) receiveMessages() {
	for {
		msg, err := c.reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error receiving message:", err)
			return
		}
		fmt.Print(msg)
	}
}

func (c *Client) sendMessages() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		fmt.Fprintf(c.writer, "[%s]: %s\n", c.name, text)
		c.writer.Flush()
	}
}

func (c *Client) getName() string {
	fmt.Print("Enter your name: ")
	name, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return chat.TrimName(name)
}

func (c *Client) getAddress() string {
	return fmt.Sprintf("%s:%s", os.Args[2], os.Args[1])
}
