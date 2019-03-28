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
	"github.com/northwesternmutual/grammes"
	"github.com/northwesternmutual/grammes/logging"
	"github.com/northwesternmutual/grammes/query/graph"
	"github.com/northwesternmutual/grammes/query/traversal"
)

// CustomTraversal is have a custom prefix for your
// traversal.
func CustomTraversal(q string) traversal.String {
	return traversal.NewCustomTraversal(q)
}

// VerboseTraversal is used for when you need to access
// the schema or verbose functions when adding vertices.
func VerboseTraversal() graph.String {
	return graph.NewGraph()
}

// Traversal is the main graph traversing object. This has
// every step you will need in a shorter format.
func Traversal() traversal.String {
	return traversal.NewTraversal()
}

var (
	// nilVertex is used for returning nothing in
	// a vertex related function.
	nilVertex = grammes.Vertex{}
	logger    logging.Logger
	client    *grammes.Client
)

// executeQuery is used as a backend for all the functions
// in this package to create a client, run the query, and
// give the result from the given query.
var executeQuery = func(host string, query string) (res [][]byte, err error) {
	if client == nil {
		if logger == nil {
			client, err = grammes.DialWithWebSocket(host)
		} else {
			client, err = grammes.DialWithWebSocket(host, grammes.WithLogger(logger))
		}
		if err != nil {
			return
		}
	}

	res, err = client.ExecuteStringQuery(query)

	return
}

// checkForClient checks for the existence of a
// client that we can use to perform quick operations.
var checkForClient = func(host string) (err error) {
	if client == nil {
		if logger == nil {
			client, err = grammes.DialWithWebSocket(host)
		} else {
			client, err = grammes.DialWithWebSocket(host, grammes.WithLogger(logger))
		}
		if err != nil {
			return err
		}
	}

	return nil
}

// SetLogger will replace the function used
// to log the queries being executed in the quick package.
func SetLogger(newLogger logging.Logger) {
	if client == nil {
		logger = newLogger
	} else {
		client.SetLogger(newLogger)
	}
}
