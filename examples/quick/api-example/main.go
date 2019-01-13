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
	"go.uber.org/zap"

	"github.com/northwesternmutual/grammes/examples/exampleutil"
	"github.com/northwesternmutual/grammes/quick"
)

func main() {
	// Setup the logger using zap.
	localhost := "ws://127.0.0.1:8182"
	logger := exampleutil.SetupLogger()
	defer logger.Sync()

	g := quick.G

	// ------------------------------------- Adding Vertex using AddVertex
	logger.Info("Adding Vertices...")

	// Drop the vertices from the graph beforehand for no interference.
	quick.ExecuteQuery(localhost, g.V().Drop())

	logger.Info("All vertices dropped from graph...")

	// Add all the testing vertices to the graph using QuickAddVertex.
	quick.AddVertex(localhost, "testingVertex1")
	quick.AddVertex(localhost, "testingVertex2")
	quick.AddVertex(localhost, "testingVertex3")

	// Gather all of the vertices in the graph by label.
	res, err := quick.ExecuteQuery(localhost, g.V().Label())
	if err != nil {
		logger.Fatal("Error executing query", zap.Error(err))
	}

	// Log the resulting vertices in the graph.
	logger.Info("Added Vertices",
		zap.ByteString("Result", res),
	)
}
