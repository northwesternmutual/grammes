// Copyright (c) 2018 Northwestern Mutual.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package grammes

import (
	"errors"

	"github.com/northwesternmutual/grammes/gremconnect"
	"github.com/northwesternmutual/grammes/gremerror"
)

// launchConnection will establish a connection to
// the Gremlin-server and launch the concurrent functions
// to handle requests, responses, and server pings.
func (c *Client) launchConnection() error {
	// Connect to the Gremlin-Server.
	if err := c.conn.Connect(); err != nil {
		c.broken = true
		c.logger.Error("cannot establish connection with dialer",
			gremerror.NewGrammesError("launchConnection", err),
		)
		return err
	}

	quit := c.conn.GetQuit()

	// Launch processes to keep track of connection & data
	go c.writeWorker(c.err, quit) // Initiates message writing to the Gremlin-server
	go c.readWorker(c.err, quit)  // Initiates message reading from the Gremlin-server
	go c.conn.Ping(c.err)         // Manages pinging and connection to the Gremlin-server

	return nil
}

// Close the connection to the Gremlin-server.
func (c *Client) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

// IsConnected returns if the client currently
// has an established connection to any servers.
func (c *Client) IsConnected() bool {
	return c.conn.IsConnected()
}

// Redial will dial a new connecting using the provided.
func (c *Client) Redial(dialer gremconnect.Dialer) error {
	if c.conn != nil {
		if !c.conn.IsDisposed() {
			c.Close()
		}
	}
	c.conn = dialer

	return c.launchConnection()
}

// Connect will connect to the configured host address.
// If this is reconnecting then make sure that your
// errs channel has a new handler function set up,
// because it won't set up automatically.
func (c *Client) Connect() error {
	if c.conn == nil {
		return errors.New("client's connection is currently nil")
	}
	if c.conn.IsDisposed() {
		// Create a new connection using the old address.
		// If you want to create a connection to a new address
		// then you have to create a new client.
		c.conn = NewWebSocketDialer(c.conn.Address())

		if err := c.launchConnection(); err != nil {
			c.logger.Error("unable to launch connection",
				gremerror.NewGrammesError("Connect", err),
			)
			return err
		}
	}
	return nil
}
