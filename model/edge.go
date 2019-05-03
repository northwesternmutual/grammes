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

package model

import (
	"errors"

	"github.com/northwesternmutual/grammes/query/traversal"
)

// Edge is the object that builds a
// connection between two or more vertices.
//
// Tinkerpop: http://tinkerpop.apache.org/javadocs/3.2.1/core/org/apache/tinkerpop/gremlin/structure/Edge.html
//
//  outVertex ---label---> inVertex.
type Edge struct {
	Type  string    `json:"@type"`
	Value EdgeValue `json:"@value"`
}

// PropertyValue will retrieve the Property for you.
func (e *Edge) PropertyValue(key string) interface{} {
	return e.Value.Properties[key].Value.Value.PropertyDetailedValue.Value
}

// ID will retrieve the Edge ID for you.
func (e *Edge) ID() string {
	return e.Value.ID.Value.RelationID
}

// Label will retrieve the Edge Label for you.
func (e *Edge) Label() string {
	return e.Value.Label
}

// OutVertexID will retrieve the id for the
// vertex that the edge goes out of.
func (e *Edge) OutVertexID() (id int64) {
	return e.Value.OutV.Value
}

// InVertexID will retrieve the id for the
// vertex that the edge goes into.
func (e *Edge) InVertexID() (id int64) {
	return e.Value.InV.Value
}

// OutVertexLabel will retrieve the label
// for the vertex the edge goes out of.
func (e *Edge) OutVertexLabel() string {
	return e.Value.OutVLabel
}

// InVertexLabel will retrieve the label
// for the vertex the edge goes into.
func (e *Edge) InVertexLabel() string {
	return e.Value.InVLabel
}

// QueryOutVertex will retrieve the vertex that
// the edge comes out of.
func (e *Edge) QueryOutVertex(client queryClient) (Vertex, error) {
	if client == nil {
		return Vertex{}, errors.New("QueryOutVertex: nil client given to Edge")
	}

	responses, err := client.ExecuteQuery(traversal.NewTraversal().
		V().HasID(e.OutVertexID()))
	if err != nil {
		return Vertex{}, err
	}

	var vertices VertexList
	vertList, err := UnmarshalVertexList(responses)
	if err != nil {
		return Vertex{}, err
	}

	vertices.Vertices = vertList

	return vertices.Vertices[0], nil
}

// QueryInVertex will retrieve the vertex that
// the edge comes out of.
func (e *Edge) QueryInVertex(client queryClient) (Vertex, error) {
	if client == nil {
		return Vertex{}, errors.New("QueryInVertex: nil client given to Edge")
	}

	responses, err := client.ExecuteQuery(traversal.NewTraversal().
		V().HasID(e.InVertexID()))
	if err != nil {
		return Vertex{}, nil
	}

	var vertices VertexList
	vertList, err := UnmarshalVertexList(responses)
	if err != nil {
		return Vertex{}, err
	}
	vertices.Vertices = vertList

	return vertices.Vertices[0], nil
}
