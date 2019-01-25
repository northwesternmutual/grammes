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

package grammes_test

import (
	"log"

	"github.com/northwesternmutual/grammes"
)

func Example_newClient() {
	// Creates a new client with the localhost IP.
	client, err := grammes.DialWithWebSocket("ws://127.0.0.1:8182")
	if err != nil {
		log.Fatalf("Error while creating client: %s\n", err.Error())
	}

	_ = client
}

func Example_executeQuery() {
	// Creates a new client with the localhost IP.
	client, err := grammes.DialWithWebSocket("ws://127.0.0.1:8182")
	if err != nil {
		log.Fatalf("Error while creating client: %s\n", err.Error())
	}

	// Create a graph traversal to use when querying.
	g := grammes.Traversal()

	// Executing a basic query to add a vertex to the graph
	// with label "testingVertex" and property "name" that equals "damien"
	res, err := client.ExecuteQuery(g.AddV("testingVertex").Property("name", "damien"))
	if err != nil {
		log.Fatalf("Querying error: %s\n", err.Error())
	}

	// Print out the result as a string
	log.Println(string(res))
}

func Example_executeStringQuery() {
	// Creates a new client with the localhost IP.
	client, err := grammes.DialWithWebSocket("ws://127.0.0.1:8182")
	if err != nil {
		log.Fatalf("Error while creating client: %s\n", err.Error())
	}

	// Executing a basic query to assure that the client is working.
	res, err := client.ExecuteStringQuery("1+3")
	if err != nil {
		log.Fatalf("Querying error: %s\n", err.Error())
	}

	// Print out the result as a string
	log.Println(string(res))
}
