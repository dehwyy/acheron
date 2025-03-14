package xdp

import (
	"crypto/tls"
	"net"
)

// config := &tls.Config{
// 	MinVersion:         tls.VersionTLS13,
// }

type ClientParams struct {
	TLS *tls.Config
}
type Client struct {
	conn net.Conn
}

func NewXDPClient(addr string, params ClientParams) (*Client, error) {
	conn, err := tls.Dial("tcp", addr, params.TLS)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
	}, nil
}

// TODO
func (cl *Client) Request() error { return nil }

// TODO
func (cl *Client) RequestStream() error { return nil }

func (cl *Client) Close() error { return cl.conn.Close() }
