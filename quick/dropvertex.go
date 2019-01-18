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

import (
	"github.com/northwesternmutual/grammes/query"
)

// DropVertexLabel will search for a vertex with the
// provided label and drop it if such a vertex exists.
func DropVertexLabel(host, label string) error {
	err := checkForClient(host)
	if err != nil {
		return err
	}

	dq := client.GraphManager.DropQuerier()
	err = dq.DropVertexLabel(label)
	if err != nil {
		return err
	}

	return nil
}

// DropVertexByID will search for vertices with the
// provided IDs and drop them if such vertices exist.
func DropVertexByID(host string, ids ...int64) error {
	err := checkForClient(host)
	if err != nil {
		return err
	}

	dq := client.GraphManager.DropQuerier()
	err = dq.DropVertexByID(ids...)
	if err != nil {
		return err
	}

	return nil
}

// DropVerticesByQuery will consume the given query
// and drop the corresponding vertices.
func DropVerticesByQuery(host string, q query.Query) error {
	err := checkForClient(host)
	if err != nil {
		return err
	}

	dq := client.GraphManager.DropQuerier()
	err = dq.DropVerticesByQuery(q)
	if err != nil {
		return err
	}

	return nil
}
