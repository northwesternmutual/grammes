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

// VerticesByQuery will gather any vertices and return them
// based on the fed in traversal query.
func VerticesByQuery(host string, q query.Query) ([]grammes.Vertex, error) {
	err := checkForClient(host)
	if err != nil {
		return nil, err
	}

	vq := client.GraphManager.GetVertexQuerier()
	res, err := vq.VerticesByQuery(q)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// AllVertices will query the gremlin server for all of
// the vertices on the graph and store them into a structured
// format for an easier way to interact with the data.
func AllVertices(host string) ([]grammes.Vertex, error) {
	err := checkForClient(host)
	if err != nil {
		return nil, err
	}

	vq := client.GraphManager.GetVertexQuerier()
	res, err := vq.AllVertices()
	if err != nil {
		return nil, err
	}

	return res, nil
}

// VertexByID will get a Vertex with the following
// ID assigned to it. This ID is unique to every
// vertex on the graph. This is the best way of finding
// vertices without any conflicting labels or properties.
func VertexByID(host string, id interface{}) (grammes.Vertex, error) {
	err := checkForClient(host)
	if err != nil {
		return nilVertex, err
	}

	vq := client.GraphManager.GetVertexQuerier()
	res, err := vq.VertexByID(id)
	if err != nil {
		return nilVertex, err
	}

	return res, nil
}

// Vertices will gather any vertices and return them
// based on the fed in traversal.
func Vertices(host, label string, properties ...interface{}) ([]grammes.Vertex, error) {
	err := checkForClient(host)
	if err != nil {
		return nil, err
	}

	vq := client.GraphManager.GetVertexQuerier()
	res, err := vq.Vertices(label, properties...)
	if err != nil {
		return nil, err
	}

	return res, nil
}
