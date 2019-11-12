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
	"encoding/json"

	"github.com/northwesternmutual/grammes/gremconnect"
)

var (
	jsonMarshalData     = json.Marshal
	gremMarshalResponse = gremconnect.MarshalResponse
)

// readWorker works on a loop and sorts messages as soon as it receives them
func (c *Client) readWorker(errs chan error, quit chan struct{}) {
	var (
		msg []byte
		err error
	)

	for {
		// attempt to read from the connection
		// and store the message back into a variable.
		if msg, err = c.conn.Read(); err != nil {
			errs <- err
			c.broken = true
			break
		}

		if msg != nil {
			// c.logger.Debug("container data", map[string]interface{}{"data": string(msg)})
			if err := c.handleResponse(msg); err != nil {
				errs <- err
			}
		}

		select {
		case <-quit:
			return
		default:
			continue
		}
	}
}

func (c *Client) retrieveResponse(id string) ([][]byte, error) {
	var (
		notifier, _ = c.resultMessenger.Load(id)
		err         error
		data        [][]byte
		dataPart    []byte
	)

	if n := <-notifier.(chan int); n == 1 {
		if dataI, ok := c.results.Load(id); ok {
			for _, d := range dataI.([]interface{}) {
				if err, ok = d.(error); ok {
					break
				}
				if dataPart, err = jsonMarshalData(d); err != nil {
					break
				}
				data = append(data, dataPart)
			}
			close(notifier.(chan int))
			c.resultMessenger.Delete(id)
			c.deleteResponse(id)
		}
	}

	return data, err
}

// deleteRespones deletes the response from the container. Used for cleanup purposes by requester.
func (c *Client) deleteResponse(id string) {
	c.results.Delete(id)
}

// saveResponse makes the response available for retrieval by the requester. Mutexes are used for thread safety.
func (c *Client) saveResponse(resp gremconnect.Response) {

	var container []interface{}

	// Retrieve the existing data (if there are multiple responses).
	if existingData, ok := c.results.Load(resp.RequestID); ok {
		container = existingData.([]interface{})
	}

	newData := append(container, resp.Data)  // Combine the old data with the new data.
	c.results.Store(resp.RequestID, newData) // Add data to buffer for future retrieval

	notifier, _ := c.resultMessenger.LoadOrStore(resp.RequestID, make(chan int, 1))

	if resp.Code != 206 {
		notifier.(chan int) <- 1
	}
}

func (c *Client) handleResponse(msg []byte) error {
	resp, err := gremMarshalResponse(msg)
	if err != nil {
		return err
	}

	if resp.Code == 407 { // Server request authentication
		return c.authenticate(resp.RequestID)
	}

	c.saveResponse(resp)
	return nil
}
