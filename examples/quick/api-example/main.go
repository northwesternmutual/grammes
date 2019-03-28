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

package main

import (
	"flag"

	"go.uber.org/zap"

	"github.com/northwesternmutual/grammes/examples/exampleutil"
	"github.com/northwesternmutual/grammes/quick"
)

var (
	// addr is used for holding the connection IP address.
	// for example this could be, "ws://127.0.0.1:8182"
	addr string
)

func main() {
	flag.StringVar(&addr, "h", "", "Connection IP")
	flag.Parse()

	logger := exampleutil.SetupLogger()
	defer logger.Sync()

	if addr == "" {
		logger.Fatal("No host address provided. Please run: go run main.go -h <host address>")
		return
	}

	g := quick.Traversal()

	// ------------------------------------- Adding Vertex using AddVertex
	logger.Info("Adding Vertices...")

	// Drop the vertices from the graph beforehand for no interference.
	quick.ExecuteQuery(addr, g.V().Drop())

	logger.Info("All vertices dropped from graph...")

	// Add all the testing vertices to the graph using QuickAddVertex.
	quick.AddVertex(addr, "testingVertex1")
	quick.AddVertex(addr, "testingVertex2")
	quick.AddVertex(addr, "testingVertex3")

	// Gather all of the vertices in the graph by label.
	responses, err := quick.ExecuteQuery(addr, g.V().Label())
	if err != nil {
		logger.Fatal("Error executing query", zap.Error(err))
	}

	// Log the resulting vertices in the graph.
	for _, res := range responses {
		logger.Info("executed query", zap.ByteString("result", res))
	}
}
