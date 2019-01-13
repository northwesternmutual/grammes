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
	localhost := "ws://127.0.0.1:8182/gremlin"
	// Setup the logger using zap.
	logger := exampleutil.SetupLogger()
	defer logger.Sync()

	// ------------------------------------- Adding Property to Vertex
	logger.Info("Adding Property to Vertex...")

	// Drop all of the vertices already in the graph for no interference.
	quick.ExecuteQuery(localhost, quick.G.V().Drop())

	logger.Info("All vertices dropped from the graph...")

	// Add a testing vertex to add a property to and saving the ID.
	v, err := quick.AddVertex(localhost, "person")
	if err != nil {
		logger.Fatal("Error adding vertex", zap.Error(err))
	}
	id := v.ID()

	// Add the testing properties to testing vertex.
	quick.SetVertexProperty(localhost, id, "first", "damien")
	quick.SetVertexProperty(localhost, id, "last", "stamates")

	// Gather the properties from the vertex via a grammes query command.
	vertex, err := quick.VertexByID(localhost, id)
	if err != nil {
		logger.Fatal("Error getting vertex", zap.Error(err))
	}

	logger.Info("Added Properties",
		zap.String("label", vertex.Value.Label),
		zap.String("first", vertex.PropertyValue("first", 0).(string)),
		zap.String("last", vertex.PropertyValue("last", 0).(string)),
	)
}
