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
	"github.com/northwesternmutual/grammes/gremerror"
)

var (
	gremPrepareRequest     = gremconnect.PrepareRequest
	gremPackageRequest     = gremconnect.PackageRequest
	gremPrepareAuthRequest = gremconnect.PrepareAuthRequest
)

func (c *Client) executeRequest(query string, bindings, rebindings map[string]string) ([][]byte, error) {
	// Construct a map containing the values along
	// with a randomly generated id to fetch the response.
	req, id, err := gremPrepareRequest(query, bindings, rebindings)
	if err != nil {
		c.logger.Error("uuid generation when preparing request",
			gremerror.NewGrammesError("executeRequest", err),
		)
		return nil, err
	}
	// Marshal the map and add on the
	// mimetype to the header of the request.
	msg, err := gremPackageRequest(req, c.gremlinVersion)
	if err != nil {
		c.logger.Error("unmarshal when packaging request",
			gremerror.NewGrammesError("executeRequest", err),
		)
		return nil, err
	}

	c.resultMessenger.Store(id, make(chan int, 1))
	c.dispatchRequest(msg)              // send the request.
	resp, err := c.retrieveResponse(id) // retrieve the response from the gremlin server
	if err != nil {
		c.logger.Error("retrieving response",
			gremerror.NewGrammesError("executeRequest", err),
		)
		return nil, err
	}

	return resp, nil
}

// writeWorker works on a loop and dispatches messages as soon as it receives them
func (c *Client) writeWorker(errs chan error, quit chan struct{}) {
	for {
		select {
		// Wait for a response from the Request
		// channel and store its value in a variable.
		case msg := <-c.request:
			// Write the message to the connection
			// and check for any errors.
			err := c.conn.Write(msg)
			if err != nil {
				errs <- err
				c.broken = true
				break
			}
		// Wait for a response from the quit
		// channel and break out of the loop.
		case <-quit:
			return
		}
	}
}

func (c *Client) authenticate(requestID string) error {
	auth, err := c.conn.Auth()
	if err != nil {
		return err
	}

	var req gremconnect.Request

	if auth != nil {
		req, err = gremPrepareAuthRequest(requestID, auth.Username, auth.Password)
		if err != nil {
			c.logger.Error("preparing authentication request",
				gremerror.NewGrammesError("authenticate", err),
			)
			return err
		}
		c.logger.Debug("authenticate: Prepared authentication request", map[string]interface{}{})
	}

	// Marshal the map and add on the
	// mimetype to the header of the request.
	msg, err := gremPackageRequest(req, c.gremlinVersion)
	if err != nil {
		c.logger.Error("packaging request",
			gremerror.NewGrammesError("authenticate", err),
		)
		return err
	}

	c.dispatchRequest(msg) // Send the request.

	return nil
}

func (c *Client) dispatchRequest(msg []byte) {
	// Send the message through a channel
	// for the writing worker to pickup and
	// write to the connection.
	c.request <- msg
}
