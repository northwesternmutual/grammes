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

package manager

import (
	"encoding/json"

	"github.com/northwesternmutual/grammes/logging"
	"github.com/northwesternmutual/grammes/model"
	"github.com/northwesternmutual/grammes/query"
	"github.com/northwesternmutual/grammes/query/cardinality"
	"github.com/northwesternmutual/grammes/query/datatype"
	"github.com/northwesternmutual/grammes/query/multiplicity"
)

var (
	// jsonUnmarshal is for monkey patching the
	// unmarshal process when testing these files.
	jsonUnmarshal = json.Unmarshal
	// nilVertex is used for returning nothing in
	// a vertex related function.
	nilVertex = model.Vertex{}
)

// unmarshalID will simply take a raw response and
// unmarshal it into an ID.
func unmarshalID(data [][]byte) (id int64, err error) {
	var resp model.VertexList

	for _, res := range data {
		var resPart model.VertexList
		err = jsonUnmarshal(res, &resPart)
		if err == nil {
			if len(resp.Vertices) > 0 {
				id = resp.Vertices[0].ID()
			}
		}

		resp.Vertices = append(resp.Vertices, resPart.Vertices...)
	}

	return id, err
}

// executor is the function type that is used when passing in executeRequest.
type executor func(string, map[string]string, map[string]string) ([][]byte, error)

// executor is the function type that is used when passing in ExecuteStringQuery.
type stringExecutor func(string) ([][]byte, error)

// MiscQuerier are miscellaneous queries for the server to perform.
type MiscQuerier interface {
	// DropAll will drop all vertices on the graph.
	DropAll() error
	// VertexCount will return the number of vertices on the graph.
	VertexCount() (count int64, err error)
	// SetVertexProperty will either add or set the property of a vertex.
	SetVertexProperty(id int64, keyAndVals ...interface{}) error
}

// SchemaQuerier handles all schema related queries to the graph.
type SchemaQuerier interface {
	// AddEdgeLabel adds a new edge label to the schema.
	AddEdgeLabel(multi multiplicity.Multiplicity, label string) (id int64, err error)
	// AddEdgeLabels adds new edge labels to the schema.
	AddEdgeLabels(multiplicityAndLabels ...interface{}) (ids []int64, err error)
	// AddPropertyKey adds a new property key to the schema.
	AddPropertyKey(label string, dt datatype.DataType, card cardinality.Cardinality) (id int64, err error)
	// CommitSchema will finalize your changes and apply them to the schema.
	CommitSchema() (res [][]byte, err error)
}

// GetVertexQuerier are functions specifically related to getting vertices.
type GetVertexQuerier interface {
	// AllVertices will return a slice of all vertices on the graph.
	AllVertices() (vertices []model.Vertex, err error)
	// VertexByID will return a single vertex based on the ID provided.
	VertexByID(id int64) (vertex model.Vertex, err error)
	// VerticesByString will return already unmarshalled vertex structs from a string query.
	VerticesByString(stringQuery string) (vertices []model.Vertex, err error)
	// VerticesByQuery will return already unmarshalled vertex structs from a query object.
	VerticesByQuery(queryObj query.Query) (vertices []model.Vertex, err error)
	// Vertices will return vertices based on the label and properties.
	Vertices(label string, properties ...interface{}) (vertices []model.Vertex, err error)
}

// GetVertexIDQuerier holds functions to gather IDs from the graph.
type GetVertexIDQuerier interface {
	// VertexIDsByString returns a slice of IDs from a string query.
	VertexIDsByString(stringQuery string) (ids []int64, err error)
	// VertexIDsByQuery returns a slice of IDs from a query object.
	VertexIDsByQuery(queryObj query.Query) (ids []int64, err error)
	// VertexIDs returns a slice of IDs based on the label and properties.
	VertexIDs(label string, properties ...interface{}) (ids []int64, err error)
}

// AddVertexQuerier are queries specific to adding vertices.
type AddVertexQuerier interface {
	// AddAPIVertex adds a vertex to the graph based on the API struct.
	AddAPIVertex(api model.APIData) (vertex model.Vertex, err error)
	// AddVertexByString adds a vertex to the graph using a string query.
	AddVertexByString(stringQuery string) (vertex model.Vertex, err error)
	// AddVertexLabels adds multiple labels to the graph.
	AddVertexLabels(labels ...string) (vertices []model.Vertex, err error)
	// AddVertexByQuery adds a vertex to the graph using a query object.
	AddVertexByQuery(queryObj query.Query) (vertex model.Vertex, err error)
	// AddVertexByStruct adds a vertex to the graph with a vertex struct.
	AddVertexByStruct(vertexStruct model.Vertex) (vertex model.Vertex, err error)
	// AddVertex adds a vertex to the graph with label and properties provided.
	AddVertex(label string, properties ...interface{}) (vertex model.Vertex, err error)
}

// DropQuerier has functions related to dropping vertices from the graph.
type DropQuerier interface {
	// DropVertexLabel drops all vertices with given label.
	DropVertexLabel(label string) error
	// DropVertexByID drops vertices based on their IDs.
	DropVertexByID(ids ...int64) error
	// DropVerticesByString drops vertices using a string query.
	DropVerticesByString(stringQuery string) error
	// DropVerticesByQuery drops vertices using a query object.
	DropVerticesByQuery(queryObj query.Query) error
}

// ExecuteQuerier handles the raw queries to the server.
type ExecuteQuerier interface {
	// ExecuteQuery will execute a query object and return its raw result.
	ExecuteQuery(queryObj query.Query) (res [][]byte, err error)
	// ExecuteStringQuery will execute a string query and return its raw result.
	ExecuteStringQuery(stringQuery string) (res [][]byte, err error)
	// ExecuteBoundQuery will execute a query object with bindings and return its raw result.
	ExecuteBoundQuery(queryObj query.Query, bindings map[string]string, rebindings map[string]string) (res [][]byte, err error)
	// ExecuteBoundStringQuery will execute a string query with bindings and return its raw result.
	ExecuteBoundStringQuery(stringQuery string, bindings map[string]string, rebindings map[string]string) (res [][]byte, err error)
}

// VertexQuerier handles the vertices on the graph.
type VertexQuerier interface {
	DropQuerier
	AddVertexQuerier
	GetVertexQuerier
	GetVertexIDQuerier
}

// GraphManager will handle all interactions with the graph itself.
type GraphManager interface {
	MiscQuerier
	VertexQuerier
	ExecuteQuerier
	SchemaQuerier

	// Returns the interface and functions associated with the MiscQuerier.
	MiscQuerier() MiscQuerier
	// Returns the interface and functions associated with the AddVertexQuerier.
	AddVertexQuerier() AddVertexQuerier
	// Returns the interface and functions associated with the GetVertexQuerier.
	GetVertexQuerier() GetVertexQuerier
	// Returns the interface and functions associated with the GetVertexIDQuerier.
	GetVertexIDQuerier() GetVertexIDQuerier
	// Returns the interface and functions associated with the DropQuerier.
	DropQuerier() DropQuerier
	// Returns the interface and functions associated with the VertexQuerier.
	VertexQuerier() VertexQuerier
	// Returns the interface and functions associated with the ExecuteQuerier.
	ExecuteQuerier() ExecuteQuerier
	// Returns the interface and functions associated with the SchemaQuerier.
	SchemaQuerier() SchemaQuerier

	// Sets the logging object used by the GraphManager.
	SetLogger(logging.Logger)
}
