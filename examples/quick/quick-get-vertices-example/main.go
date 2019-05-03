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

	"github.com/northwesternmutual/grammes/examples/exampleutil"
	"github.com/northwesternmutual/grammes/quick"

	"go.uber.org/zap"
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

	if addr == "" {
		logger.Fatal("No host address provided. Please run: go run main.go -h <host address>")
		return
	}

	defer func() {
		logger.Sync()
		quick.DropAll(addr)
	}()

	quick.DropAll(addr)

	// ------------------------------------- Gathering the Vertices from Graph
	logger.Info("Gathering Vertices...")

	quick.AddVertex(addr, "testvertex1")
	quick.AddVertex(addr, "testvertex2")
	quick.AddVertex(addr, "testvertex3")

	// Using this command it takes all of the vertices from the graph
	// by using `g.V()` and unmarshals them into a structured format.
	vertices, err := quick.AllVertices(addr)
	if err != nil {
		logger.Fatal("Error getting all vertices", zap.Error(err))
	}

	// Log out the resulting vertices in the structured format.
	for _, v := range vertices {
		logger.Info("Vertex",
			zap.String("Label", v.Value.Label),
		)
	}
}
