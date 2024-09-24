package server

import (
    "net"
    "net/pkg/common"
)

type Client struct {
    Conn net.Conn
    Name string
}

func NewClient(conn net.Conn, name string) *Client {
    return &Client{
        Conn: conn,
        Name: name,
    }
}

func (c *Client) SendMessage(msg string) error {
    _, err := c.Conn.Write([]byte(msg))
    return err
}

func (c *Client) Disconnect() error {
    return c.Conn.Close()
}

func (c *Client) String() string {
    return c.Name
}

func (c *Client) WelcomeMessage() string {
    return common.WelcomeMessage([]string{})
}