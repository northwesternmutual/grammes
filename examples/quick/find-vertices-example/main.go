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

	// remove any interfering vertices left on the graph.
	quick.DropAll(addr)

	// drop the testing vertices when we're done with them.
	defer quick.DropAll(addr)

	// Add some testing vertices with the
	// same labels but different properties.
	quick.AddVertex(addr, "Person", "name", "damien")
	quick.AddVertex(addr, "Person", "name", "bahram")

	// Retrieve the IDs of the vertices I'm trying to find.
	damienID, _ := quick.VertexIDs(addr, "Person", "name", "damien")
	bahramID, _ := quick.VertexIDs(addr, "Person", "name", "bahram")

	// Use the IDs to retrieve the Vertex structures.
	damienV, _ := quick.VertexByID(addr, damienID[0])
	bahramV, _ := quick.VertexByID(addr, bahramID[0])

	// Log out the data retrieved.
	logger.Info("Vertex", zap.Any("id", damienV.ID()))
	logger.Info("Vertex", zap.Any("name", damienV.PropertyValue("name", 0)))

	logger.Info("Vertex", zap.Any("id", bahramV.ID()))
	logger.Info("Vertex", zap.Any("name", bahramV.PropertyValue("name", 0)))
}
