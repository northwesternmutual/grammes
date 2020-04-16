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
	"github.com/northwesternmutual/grammes/gremerror"

	"github.com/northwesternmutual/grammes/query/traversal"
)

var newTrav = traversal.NewTraversal

// QueryRefresh gets the vertex from the graph
// and refreshes its values to match.
func (v *Vertex) QueryRefresh(client queryClient) error {
	if client == nil {
		return gremerror.NewGrammesError("QueryRefresh", gremerror.ErrNilClient)
	}

	var query = newTrav().V().HasID(v.ID())

	responses, err := client.ExecuteQuery(query)
	if err != nil {
		return gremerror.NewQueryError("QueryRefresh", query.String(), err)
	}

	var vertices VertexList
	vertList, err := UnmarshalVertexList(responses)
	if err != nil {
		return err
	}

	vertices.Vertices = vertList

	if len(vertices.Vertices) == 0 {
		return gremerror.NewGrammesError("QueryRefresh", gremerror.ErrEmptyResponse)
	}

	*v = vertices.Vertices[0]

	return nil
}

// QueryBothEdges will execute bothE() on this vertex for you
// without having to make another lengthy call to ExecuteQuery.
func (v *Vertex) QueryBothEdges(client queryClient, labels ...string) ([]Edge, error) {
	if client == nil {
		return nil, gremerror.NewGrammesError("QueryBothEdges", gremerror.ErrNilClient)
	}

	var query = newTrav().V().HasID(v.ID()).BothE(labels...)

	responses, err := client.ExecuteQuery(query)
	if err != nil {
		return nil, gremerror.NewQueryError("QueryBothEdges", query.String(), err)
	}

	var edges EdgeList
	edgeList, err := UnmarshalEdgeList(responses)
	if err != nil {
		return nil, err
	}

	edges.Edges = edgeList

	return edges.Edges, nil
}

// QueryOutEdges will execute outE() on this vertex for you
// without having to make another lengthy call to ExecuteQuery.
func (v *Vertex) QueryOutEdges(client queryClient, labels ...string) ([]Edge, error) {
	if client == nil {
		return nil, gremerror.NewGrammesError("QueryOutEdges", gremerror.ErrNilClient)
	}

	responses, err := client.ExecuteQuery(traversal.NewTraversal().V().HasID(v.ID()).OutE(labels...))
	if err != nil {
		return nil, err
	}

	var edges EdgeList
	edgeList, err := UnmarshalEdgeList(responses)
	if err != nil {
		return nil, err
	}

	edges.Edges = edgeList

	return edges.Edges, nil
}

// QueryInEdges will execute inE() on this vertex for you
// without having to make another lengthy call to ExecuteQuery.
func (v *Vertex) QueryInEdges(client queryClient, labels ...string) ([]Edge, error) {
	if client == nil {
		return nil, gremerror.NewGrammesError("QueryInEdges", gremerror.ErrNilClient)
	}

	var query = newTrav().V().HasID(v.ID()).InE(labels...)

	responses, err := client.ExecuteQuery(query)
	if err != nil {
		return nil, gremerror.NewQueryError("QueryInEdges", query.String(), err)
	}

	var edges EdgeList
	edgeList, err := UnmarshalEdgeList(responses)
	if err != nil {
		return nil, err
	}

	edges.Edges = edgeList

	return edges.Edges, nil
}

// AddEdge adds an outgoing edge from this Vertex object to
// another Vertex object via its unique ID.
func (v *Vertex) AddEdge(client queryClient, label string, outVID interface{}, properties ...interface{}) (Edge, error) {
	if client == nil {
		return Edge{}, gremerror.NewGrammesError("AddEdge", gremerror.ErrNilClient)
	}

	var query = newTrav().V().HasID(v.ID()).AddE(label).To(newTrav().V().HasID(outVID).Raw())
	// query := fmt.Sprintf("g.V().hasId(%v).addE(\"%s\").to(V().hasId(%v))", v.ID(), label, outVID)

	if len(properties)%2 != 0 {
		return Edge{}, gremerror.NewGrammesError("AddEdge", gremerror.ErrOddNumberOfParameters)
	}

	if len(properties) > 0 {
		for i := 0; i < len(properties); i += 2 {
			query.AddStep("property", properties[i], properties[i+1])
		}
	}

	// Execute the built command.
	responses, err := client.ExecuteQuery(query)
	if err != nil {
		return Edge{}, gremerror.NewQueryError("AddEdge", query.String(), err)
	}

	var edges EdgeList
	edgeList, err := UnmarshalEdgeList(responses)
	if err != nil {
		return Edge{}, err
	}

	edges.Edges = edgeList

	if len(edges.Edges) == 0 {
		return Edge{}, gremerror.NewGrammesError("AddEdge", gremerror.ErrEmptyResponse)
	}

	return edges.Edges[0], nil
}

// Drop will drop the current vertex that's being called from.
func (v *Vertex) Drop(client queryClient) error {
	if client == nil {
		return gremerror.NewGrammesError("Drop", gremerror.ErrNilClient)
	}

	_, err := client.ExecuteQuery(newTrav().V().HasID(v.ID()).Drop())

	return err
}

// DropProperties drops the properties from the vertex so they don't exist.
func (v *Vertex) DropProperties(client queryClient, properties ...string) error {
	if client == nil {
		return gremerror.NewGrammesError("DropProperties", gremerror.ErrNilClient)
	}

	_, err := client.ExecuteQuery(newTrav().V().HasID(v.ID()).Properties(properties...).Drop())
	if err != nil {
		return err
	}

	return v.QueryRefresh(client)
}

// AddProperty will add a property to the vertex in the graph and return a
// new version of the vertex with the property added to it in the structure.
func (v *Vertex) AddProperty(client queryClient, key string, value interface{}) error {
	if client == nil {
		return gremerror.NewGrammesError("AddProperty", gremerror.ErrNilClient)
	}

	_, err := client.ExecuteQuery(newTrav().V().HasID(v.ID()).Property(key, value))
	if err != nil {
		return err
	}

	return v.QueryRefresh(client)
}
