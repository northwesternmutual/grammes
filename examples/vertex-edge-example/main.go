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

	"github.com/northwesternmutual/grammes"
	"github.com/northwesternmutual/grammes/examples/exampleutil"
)

// This example will create a new client
// and show the ability to get edges from
// vertices and vertices from edges.

var (
	// addr is used for holding the connection IP address.
	// for example this could be, "ws://127.0.0.1:8182"
	addr string

	logger *zap.Logger
)

func main() {
	flag.StringVar(&addr, "h", "", "Connection IP")
	flag.Parse()

	logger = exampleutil.SetupLogger()
	defer logger.Sync()

	if addr == "" {
		logger.Fatal("No host address provided. Please run: go run main.go -h <host address>")
		return
	}

	// Create a new Grammes client with a standard websocket.
	client, err := grammes.DialWithWebSocket(addr)
	if err != nil {
		logger.Fatal("Couldn't create client", zap.Error(err))
	}

	// Drop all vertices on the graph currently.
	client.DropAll()

	// Drop the testing vertices when finished.
	defer client.DropAll()

	// Adding two testing vertices.
	vertex1, err := client.AddVertex("person1", "name", "damien")
	if err != nil {
		logger.Fatal("Failed to add vertex", zap.Error(err))
	}

	vertex2, err := client.AddVertex("person2", "name", "george")
	if err != nil {
		logger.Fatal("Failed to add vertex", zap.Error(err))
	}

	// Add an edge between the two vertices and
	// add two properties to the edge.
	vertex1.AddEdge(client, "friendsWith", vertex2.ID(),
		"driveDist", "10min",
		"ageDiff", 44,
	)

	// Get the edges based on vertex1's out edges.
	edges, err := vertex1.QueryOutEdges(client)
	if err != nil {
		logger.Fatal("Error while querying for outer edges", zap.Error(err))
	}

	printEdges(client, edges)
}

func printEdges(client *grammes.Client, edges []grammes.Edge) {
	if len(edges) > 0 {
		// Get the vertices based on the edge's
		// stored ID's about its related vertices.
		v1, err := edges[0].QueryOutVertex(client)
		if err != nil {
			logger.Fatal("Error while querying for outer vertices", zap.Error(err))
		}
		v2, err := edges[0].QueryInVertex(client)
		if err != nil {
			logger.Fatal("Error while querying for outer vertices", zap.Error(err))
		}
		edges[0].QueryOutVertex(client)
		if err != nil {
			logger.Fatal("Error while querying for outer vertices", zap.Error(err))
		}
		edges[0].QueryInVertex(client)
		if err != nil {
			logger.Fatal("Error while querying for outer vertices", zap.Error(err))
		}
		// Print the information about the edge including
		// its ID, label, and its properties.
	
		logger.Info("Edge",
			zap.Any("ID", edges[0].ID()),
			zap.String("Label", edges[0].Label()),
			zap.Any("ageDiff", edges[0].PropertyValue("ageDiff")),
			zap.Any("driveDist", edges[0].PropertyValue("driveDist")),
		)
	
		logger.Info("OutVertex",
			zap.Any("ID", edges[0].OutVertexID()),
			zap.String("Label", edges[0].OutVertexLabel()),
			zap.Any("Name", v1.PropertyValue("name", 0)),
		)
	
		logger.Info("InVertex",
			zap.Any("ID", edges[0].InVertexID()),
			zap.String("Label", edges[0].InVertexLabel()),
			zap.Any("Name", v2.PropertyValue("name", 0)),
		)
	}
}