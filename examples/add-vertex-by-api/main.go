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
	"encoding/json"
	"flag"
	"io/ioutil"

	"go.uber.org/zap"

	"github.com/northwesternmutual/grammes"
	"github.com/northwesternmutual/grammes/examples/exampleutil"
)

func prepareAPI(path string) (grammes.APIData, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return grammes.APIData{}, err
	}

	var apiData grammes.APIData

	err = json.Unmarshal(raw, &apiData)

	return apiData, err
}

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

	// Create a new Grammes client with a standard websocket.
	client, err := grammes.DialWithWebSocket(addr)
	if err != nil {
		logger.Fatal("Couldn't create client", zap.Error(err))
	}

	defer client.DropAll()

	// DropAll will remove all vertices from the graph currently.
	// Essentially blank slating all of our data.
	client.DropAll()

	// Get the count of vertices on the graph before adding through the API.
	count, err := client.VertexCount()
	if err != nil {
		logger.Fatal("Failed to count vertices", zap.Error(err))
	}

	logger.Info("Number of Vertices Before", zap.Int64("count", count))

	// Read the request example JSON from the assets/ folder.
	api, err := prepareAPI("../../assets/test/request.json")
	if err != nil {
		logger.Fatal("Failed to prepare API Json", zap.Error(err))
	}

	// Add a vertex using the API and print ouf the resulting vertex.
	vertex, err := client.AddAPIVertex(api)
	if err != nil {
		logger.Fatal("Failed to add vertex", zap.Error(err))
	}

	// Print out the vertex struct of the added vertex.
	logger.Sugar().Infow("Resulting Vertex",
		"Label", vertex.Label(),
		"ID", vertex.ID(),
		"NAME_TXT", vertex.PropertyValue("NAME_TXT", 0),
	)

	// Show the final vertex count.
	// Should be 1.
	count, err = client.VertexCount()
	if err != nil {
		logger.Fatal("Failed to count vertices", zap.Error(err))
	}

	logger.Info("Number of Vertices After", zap.Int64("count", count))
}
