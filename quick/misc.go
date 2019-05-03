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

package quick

// DropAll drops everything from the graph.
func DropAll(host string) error {
	err := checkForClient(host)
	if err != nil {
		return err
	}

	mq := client.GraphManager.MiscQuerier()
	err = mq.DropAll()
	if err != nil {
		return err
	}

	return nil
}

// SetVertexProperty will search the graph for a vertex
// with the given ID and set the properties provided.
func SetVertexProperty(host string, id int64, properties ...interface{}) error {
	err := checkForClient(host)
	if err != nil {
		return err
	}

	mq := client.GraphManager.MiscQuerier()
	err = mq.SetVertexProperty(id, properties...)
	if err != nil {
		return err
	}

	return nil
}

// VertexCount retrieves the number of vertices
// that are currently on the graph as an int64.
func VertexCount(host string) (int64, error) {
	err := checkForClient(host)
	if err != nil {
		return 0, err
	}

	mq := client.GraphManager.MiscQuerier()
	res, err := mq.VertexCount()
	if err != nil {
		return 0, err
	}

	return res, nil
}
