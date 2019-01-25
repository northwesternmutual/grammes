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
	"github.com/northwesternmutual/grammes/query/traversal"
)

// Vertex maintains pointers to both a set
// of incoming and outgoing Edge objects. The
// outgoing edges are the edges for which the
// Vertex is a tail. The incoming edges are those
// edges for which the Vertex is the head.
//
// TinkerPop: http://tinkerpop.apache.org/javadocs/3.2.1/core/org/apache/tinkerpop/gremlin/structure/Vertex.html
//
//  ---inEdges---> vertex ---outEdges--->.
type Vertex struct {
	Type  string      `json:"@type"`
	Value VertexValue `json:"@value"`
}

// NewVertex is to create a Vertex struct without all the hassle.
func NewVertex(label string, properties ...interface{}) Vertex {
	var v = Vertex{
		Type: "g:Vertex",
		Value: VertexValue{
			ID:         ID{},
			Properties: make(PropertyMap),
			Label:      label,
		},
	}

	if len(properties)%2 != 0 {
		return v
	}

	for i := 0; i < len(properties); i += 2 {
		v.Value.Properties[properties[i].(string)] = []Property{
			NewProperty(properties[i].(string), properties[i+1]),
		}
	}

	return v
}

// PropertyValue returns the value of a property
// without having to traverse all the way through the structures.
func (v *Vertex) PropertyValue(key string, index int) interface{} {
	return v.Value.Properties[key][index].GetValue()
}

// PropertyMap returns a copy of the properties in a map
// within the Vertex itself. Altering this copy will not
// affect the vertex on the graph.
func (v *Vertex) PropertyMap() PropertyMap {
	return v.Value.Properties
}

// ID will retrieve the Vertex ID for you
// without having to traverse all the way through the structures.
func (v *Vertex) ID() int64 {
	return v.Value.ID.Value
}

// Label retrieves the label of the vertex
// without having to traverse all the wya through the structures.
func (v *Vertex) Label() string {
	return v.Value.Label
}

// Traversal returns a new query string that is
// directed to the current vertex being referenced.
func (v *Vertex) Traversal() traversal.String {
	return traversal.NewTraversal().V().HasID(v.ID())
}
