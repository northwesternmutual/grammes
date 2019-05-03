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
	"fmt"

	"github.com/northwesternmutual/grammes"
	"github.com/northwesternmutual/grammes/examples/exampleutil"
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

	g := grammes.Traversal()

	// Add testing vertex.
	vert, err := client.AddVertex("person",
		"age", 19,
		"firstname", "damien",
		"middlename", "socrates",
		"lastname", "stamates",
	)
	if err != nil {
		logger.Fatal("Couldn't add vertex", zap.Error(err))
	}

	// Print out columns of the received data.
	fmt.Printf("%-20s %-16s %-16s\n", "Type", "Label", "Value")
	for _, v := range vert.PropertyMap() {
		for _, p := range v {
			fmt.Printf("%#-20v %#-16v %#-16v\n", p.Type, p.Value.Label, p.GetValue())
		}
		// fmt.Printf("%#-20v %#-8v\n", p.Value.Label, p.GetValue())
	}

	fmt.Println()

	// Execute a query to get the properties from testing vertex.
	res, err := client.ExecuteQuery(
		g.V().HasLabel("person").Properties(),
	)
	if err != nil {
		logger.Fatal("Querying error", zap.Error(err))
	}

	var props grammes.PropertyList

	// Unmarshal raw data into a PropertyList
	json.Unmarshal(res, &props)

	// Print out columns of the received data.
	fmt.Printf("%-20s %-8s\n", "Label", "Value")
	for _, p := range props.Properties {
		fmt.Printf("%#-20v %#-8v\n", p.Value.Label, p.GetValue())
	}
}
