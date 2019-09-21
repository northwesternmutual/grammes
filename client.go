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
	"sync"

	"github.com/northwesternmutual/grammes/gremconnect"
	"github.com/northwesternmutual/grammes/gremerror"
	"github.com/northwesternmutual/grammes/logging"
	"github.com/northwesternmutual/grammes/manager"
)

// maxConCurrentMessages determines the size of the request channel.
const maxConCurrentMessages = 3

// Client is used to handle the graph, schema, connection,
// and basic debug logging when querying the graph database.
// When handling more than one queries to the graph database,
// it is better to use a client rather than the quick package.
type Client struct {
	// GraphManager is an interface used to have query functions that
	// handle all interactions to the graph database.
	manager.GraphManager
	// conn stores the connection dialer which controls the
	// process of sending and receiving messages to the TinkerPop server.
	conn gremconnect.Dialer
	// gremlinVersion determines the version of gremlin that is being used.
	// The gremlinVersion is defaulted to 2. Grammes supports 2 and 3.
	// Neptune: https://docs.aws.amazon.com/neptune/latest/userguide/access-graph-gremlin-differences.html
	gremlinVersion string
	// errs is a channel to pass errors that involve connection,
	// responses, and requests to and from the TinkerPop server.
	err chan error
	// request is a buffer for requests to be sent to the TinkerPop server.
	request chan []byte
	// results is a buffer for responses with their [ID] and Data.
	results *sync.Map
	// resultMessenger is used to store the ID and notifier when result is ready.
	resultMessenger *sync.Map
	// broken is used to determine if the client is not working properly.
	broken bool
	// logger is used to log out debug statements and errors from the client.
	logger logging.Logger
}

// setupClient default values some fields in the client.
func setupClient() *Client {
	return &Client{
		err:             make(chan error),
		request:         make(chan []byte, maxConCurrentMessages),
		results:         &sync.Map{},
		resultMessenger: &sync.Map{},
		logger:          logging.NewNilLogger(),
		gremlinVersion:  "3",
	}
}

// Dial returns a working client with the given dialer and configurations.
func Dial(conn gremconnect.Dialer, cfgs ...ClientConfiguration) (*Client, error) {
	c := setupClient()
	c.conn = conn
	// Go through the configurations to customize the client.
	for _, conf := range cfgs {
		conf(c)
	}

	// launch the connection to the TinkerPop server,
	// and spin up the read, write, and ping workers.
	if err := c.launchConnection(); err != nil {
		c.logger.Error("unable to launch connection",
			gremerror.NewGrammesError("Dial", err),
		)
		return c, err
	}

	// GraphManager should be set because it's after the connection is created.
	c.GraphManager = manager.NewGraphManager(c.conn, c.logger, c.executeRequest)

	return c, nil
}

// DialWithWebSocket returns a new client with a websocket dialer
// and possible client configurations.
func DialWithWebSocket(host string, cfgs ...ClientConfiguration) (*Client, error) {
	// Create the new client using the Dial function and the
	// new established connection using a websocket.
	return Dial(NewWebSocketDialer(host), cfgs...)
}

// SetLogger will switch out the old logger with
// a new one provided as a parameter.
func (c *Client) SetLogger(newLogger logging.Logger) {
	c.logger = newLogger
	c.GraphManager.SetLogger(newLogger)
}

// IsBroken returns whether the client is broken or not.
func (c *Client) IsBroken() bool {
	return c.broken
}

// Address returns the current host address from the dialer.
func (c *Client) Address() string {
	return c.conn.Address()
}

// Auth will get the authentication user and pass from the dialer.
func (c *Client) Auth() (*gremconnect.Auth, error) {
	return c.conn.Auth()
}
