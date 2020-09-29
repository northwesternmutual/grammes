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
	"github.com/northwesternmutual/grammes/gremconnect"
	"strconv"
	"time"

	"github.com/northwesternmutual/grammes/logging"
)

// ClientConfiguration is the type used for configuring
// and changing values in the client and the dialer.
type ClientConfiguration func(*Client)

// WithErrorChannel will assign an error channel to send
// connection errors through for the user to handle.
func WithErrorChannel(err chan error) ClientConfiguration {
	return func(c *Client) {
		c.err = err
	}
}

// WithLogger will replace the default zap.Logger with a
// custom logger that implements the logging.Logger interface.
func WithLogger(newLogger logging.Logger) ClientConfiguration {
	return func(c *Client) {
		c.logger = newLogger
	}
}

// WithGremlinVersion sets the version of the gremlin traversal
// language being used by the client.
func WithGremlinVersion(versionNumber int) ClientConfiguration {
	return func(c *Client) {
		c.gremlinVersion = strconv.Itoa(versionNumber)
	}
}

// WithMaxConcurrentMessages sets the limit as to how many
// requests can be stored in the requests buffer.
func WithMaxConcurrentMessages(limit int) ClientConfiguration {
	return func(c *Client) {
		c.request = make(chan []byte, limit)
	}
}

// WithAuthUserPass sets the authentication credentials
// within the dialer. (This includes the username and password)
func WithAuthUserPass(user, pass string) ClientConfiguration {
	return func(c *Client) {
		c.conn.SetAuth(user, pass)
	}
}

// WithHTTPAuth sets the authentication provider
// within the dialer
func WithHTTPAuth(provider gremconnect.AuthProvider) ClientConfiguration {
	return func(c *Client) {
		c.conn.SetHTTPAuth(provider)
	}
}

// WithTimeout sets the timeout to wait when dialing
// with the dialer in seconds.
func WithTimeout(interval time.Duration) ClientConfiguration {
	return func(c *Client) {
		c.conn.SetTimeout(interval)
	}
}

// WithPingInterval sets the interval of ping sending for know is
// connection is alive and in consequence the client is connected.
func WithPingInterval(interval time.Duration) ClientConfiguration {
	return func(c *Client) {
		c.conn.SetPingInterval(interval)
	}
}

// WithWritingWait sets the time to wait when
// writing with the dialer in seconds.
func WithWritingWait(interval time.Duration) ClientConfiguration {
	return func(c *Client) {
		c.conn.SetWritingWait(interval)
	}
}

// WithReadingWait sets the time to wait when
// reading with the dialer in seconds.
func WithReadingWait(interval time.Duration) ClientConfiguration {
	return func(c *Client) {
		c.conn.SetReadingWait(interval)
	}
}

// WithRequestTimeout sets the timeout when
// reading a request from the gremlin server
func WithRequestTimeout(interval time.Duration) ClientConfiguration {
	return func(c *Client) {
		c.requestTimeout = interval
	}
}

// WithWriteBufferSize sets the max write buffer size
// for the websocket frame
func WithWriteBufferSize(writeBufferSize int) ClientConfiguration {
	return func(c *Client) {
		c.conn.SetWriteBufferSize(writeBufferSize)
	}
}

// WithReadBufferSize sets the max read buffer size
// for the websocket frame
func WithReadBufferSize(readBufferSize int) ClientConfiguration {
	return func(c *Client) {
		c.conn.SetReadBufferSize(readBufferSize)
	}
}

// WithHandshakeTimeout sets the websocket handshake timeout
func WithHandshakeTimeout(handshakeTimeout time.Duration) ClientConfiguration {
	return func(c *Client) {
		c.conn.SetHandshakeTimeout(handshakeTimeout)
	}
}

// WithCompression sets the compression
// flag for websocket connections
func WithCompression(enableCompression bool) ClientConfiguration {
	return func(c *Client) {
		c.conn.SetCompression(enableCompression)
	}
}
