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
	"github.com/northwesternmutual/grammes/query"
)

// AddAPIVertex will add a vertex to the graph belonging
// to the given host with the API object. This API object can
// be used like a normal vertex in case you'd rather make a vertex
// through a structured format rather than a gremlin command.
func AddAPIVertex(host string, data grammes.APIData) (grammes.Vertex, error) {
	err := CheckForClient(host)
	if err != nil {
		return nilVertex, err
	}

	vq := client.GraphManager.AddVertexQuerier()
	res, err := vq.AddAPIVertex(data)
	if err != nil {
		return nilVertex, err
	}

	return res, nil
}

// AddVertexByStruct will take a Vertex struct and create
// a new vertex out of it in the Gremlin server. The only
// exception is that you cannot manually set the ID.
func AddVertexByStruct(host string, vertex grammes.Vertex) (grammes.Vertex, error) {
	err := CheckForClient(host)
	if err != nil {
		return nilVertex, err
	}

	vq := client.GraphManager.AddVertexQuerier()
	res, err := vq.AddVertexByStruct(vertex)
	if err != nil {
		return nilVertex, err
	}

	return res, nil
}

// AddVertex will add a vertex label to the
// graph that is associated with the given host.
func AddVertex(host, label string, properties ...interface{}) (grammes.Vertex, error) {
	err := CheckForClient(host)
	if err != nil {
		return nilVertex, err
	}

	vq := client.GraphManager.AddVertexQuerier()
	res, err := vq.AddVertex(label, properties...)
	if err != nil {
		return nilVertex, err
	}

	return res, nil
}

// AddVertexLabels will do the same as AddVertexLabel, but with
// the ability to add multiple labels at a time.
func AddVertexLabels(host string, labels ...string) ([]grammes.Vertex, error) {
	err := CheckForClient(host)
	if err != nil {
		return nil, err
	}

	vq := client.GraphManager.AddVertexQuerier()
	res, err := vq.AddVertexLabels(labels...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// AddVertexByQuery takes a query and returns an added Vertex
// by turning it into a string.
func AddVertexByQuery(host string, q query.Query) (grammes.Vertex, error) {
	err := CheckForClient(host)
	if err != nil {
		return nilVertex, err
	}

	vq := client.GraphManager.AddVertexQuerier()
	res, err := vq.AddVertexByQuery(q)
	if err != nil {
		return nilVertex, err
	}

	return res, nil
}
