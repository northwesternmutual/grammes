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

	"github.com/northwesternmutual/grammes/query/cardinality"
	"github.com/northwesternmutual/grammes/query/datatype"
	"github.com/northwesternmutual/grammes/query/multiplicity"

	"github.com/northwesternmutual/grammes/logging"
	"github.com/northwesternmutual/grammes/model"
	"github.com/northwesternmutual/grammes/query"
)

type executor func(string, map[string]string, map[string]string) ([]byte, error)

type stringExecutor func(string) ([]byte, error)

func unmarshalID(data []byte) (id int64, err error) {
	var resp model.VertexList
	err = jsonUnmarshal(data, &resp)
	if err == nil {
		if len(resp.Vertices) > 0 {
			id = resp.Vertices[0].ID()
		}
	}
	return id, err
}

// jsonUnmarshal is for monkey patching the
// unmarshal process when testing these files.
var jsonUnmarshal = json.Unmarshal

// nilVertex is used for returning nothing in
// a vertex related function.
var nilVertex = model.Vertex{}

// MiscQuerier are miscellaneous queries for the server to perform.
type MiscQuerier interface {
	DropAll() error
	VertexCount() (int64, error)
	SetVertexProperty(int64, ...interface{}) error
}

// SchemaQuerier handles all schema related queries to the graph.
type SchemaQuerier interface {
	AddEdgeLabel(multiplicity.Multiplicity, string) (int64, error)
	AddEdgeLabels(...interface{}) ([]int64, error)
	AddPropertyKey(string, datatype.DataType, cardinality.Cardinality) (int64, error)
	CommitSchema() ([]byte, error)
}

// GetVertexQuerier are functions specifically related to getting vertices.
type GetVertexQuerier interface {
	AllVertices() ([]model.Vertex, error)
	VertexByID(int64) (model.Vertex, error)
	VerticesByString(string) ([]model.Vertex, error)
	VerticesByQuery(query.Query) ([]model.Vertex, error)
	Vertices(string, ...interface{}) ([]model.Vertex, error)
}

// GetVertexIDQuerier holds functions to gather IDs from the graph.
type GetVertexIDQuerier interface {
	VertexIDsByString(string) ([]int64, error)
	VertexIDsByQuery(query.Query) ([]int64, error)
	VertexIDs(string, ...interface{}) ([]int64, error)
}

// AddVertexQuerier are queries specific to adding vertices.
type AddVertexQuerier interface {
	AddVertexByString(string) (model.Vertex, error)
	AddAPIVertex(model.APIData) (model.Vertex, error)
	AddVertexLabels(...string) ([]model.Vertex, error)
	AddVertexByQuery(query.Query) (model.Vertex, error)
	AddVertexByStruct(model.Vertex) (model.Vertex, error)
	AddVertex(string, ...interface{}) (model.Vertex, error)
}

// DropQuerier has functions related to dropping vertices from the graph.
type DropQuerier interface {
	DropVertexLabel(string) error
	DropVertexByID(...int64) error
	DropVerticesByString(string) error
	DropVerticesByQuery(query.Query) error
}

// ExecuteQuerier handles the raw queries to the server.
type ExecuteQuerier interface {
	ExecuteQuery(query.Query) ([]byte, error)
	ExecuteStringQuery(string) ([]byte, error)
	ExecuteBoundQuery(query.Query, map[string]string, map[string]string) ([]byte, error)
	ExecuteBoundStringQuery(string, map[string]string, map[string]string) ([]byte, error)
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

	MiscQuerier() MiscQuerier
	AddVertexQuerier() AddVertexQuerier
	GetVertexQuerier() GetVertexQuerier
	GetVertexIDQuerier() GetVertexIDQuerier
	DropQuerier() DropQuerier
	VertexQuerier() VertexQuerier
	ExecuteQuerier() ExecuteQuerier
	SchemaQuerier() SchemaQuerier

	SetLogger(logging.Logger)
}
