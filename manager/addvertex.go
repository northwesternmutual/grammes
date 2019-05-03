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
	"strconv"

	"github.com/northwesternmutual/grammes/gremerror"
	"github.com/northwesternmutual/grammes/logging"
	"github.com/northwesternmutual/grammes/model"
	"github.com/northwesternmutual/grammes/query"
	"github.com/northwesternmutual/grammes/query/traversal"
)

type addVertexQueryManager struct {
	logger             logging.Logger
	executeStringQuery stringExecutor
}

func newAddVertexQueryManager(logger logging.Logger, executeString stringExecutor) *addVertexQueryManager {
	return &addVertexQueryManager{
		logger:             logger,
		executeStringQuery: executeString,
	}
}

// AddAPIVertex is used for adding a vertex to the graph with an available
// API object if you want to create it in a struct format rather than a command.
func (v *addVertexQueryManager) AddAPIVertex(data model.APIData) (model.Vertex, error) {
	query := traversal.NewTraversal().AddV(data.Label)
	// Add properties to the vertex based on the API.
	for k, v := range data.Properties {
		query.AddStep("property", k, v)
	}

	addedVertex, err := v.AddVertexByString(query.String())
	if err != nil {
		v.logger.Error("AddAPIVertex: invalid query adding vertex", err)
		return addedVertex, err
	}

	return addedVertex, nil
}

// AddVertexByStruct will take a Vertex struct and create
// a new vertex out of it in the Gremlin server. The only
// exception is that you cannot manually set the ID.
func (v *addVertexQueryManager) AddVertexByStruct(vertex model.Vertex) (model.Vertex, error) {
	var properties []interface{}

	for key, vals := range vertex.Value.Properties {
		properties = append(properties, key)
		for _, val := range vals {
			properties = append(properties, val.Value.Value.Value)
		}
	}

	addedVertex, err := v.AddVertex(vertex.Label(), properties...)
	if err != nil {
		v.logger.Error("AddVertexByStruct: invalid query adding vertex", err)
		return addedVertex, err
	}

	return addedVertex, nil
}

// AddVertex will take in a label and optional properties
// to construct a command to create a vertex in the gremlin
// server/graph database and return the structured data of your
// added vertex.
// Note:
// The properties interface can be any Token, string, or
// default Golang types such as int, bool, or byte. Other
// custom types like cardinality should not be used for this.
func (v *addVertexQueryManager) AddVertex(label string, properties ...interface{}) (model.Vertex, error) {
	if len(properties) > 0 && len(properties)%2 != 0 {
		v.logger.Error("number of parameters ["+strconv.Itoa(len(properties))+"]",
			gremerror.NewGrammesError("AddVertex", gremerror.ErrOddNumberOfParameters),
		)
		return nilVertex, gremerror.ErrOddNumberOfParameters
	}

	// Begin the command with signature and label.
	query := traversal.NewTraversal().AddV(label)

	for i := 0; i < len(properties); i += 2 {
		query.AddStep("property", properties[i], properties[i+1])
	}

	return v.AddVertexByString(query.String())
}

// AddVertexLabels will do the same as AddVertexLabel, but with
// the ability to add multiple labels at a time.
func (v *addVertexQueryManager) AddVertexLabels(labels ...string) ([]model.Vertex, error) {
	var vertices []model.Vertex

	for _, l := range labels {
		vertex, err := v.AddVertex(l)
		if err != nil {
			return nil, err
		}

		vertices = append(vertices, vertex)
	}

	return vertices, nil
}

// AddVertexByQuery takes a query and returns an added Vertex
// by turning it into a string.
func (v *addVertexQueryManager) AddVertexByQuery(q query.Query) (model.Vertex, error) {
	return v.AddVertexByString(q.String())
}

// AddVertexByString will take a query that's intended to add a vertex
// and return it as a Vertex struct.
func (v *addVertexQueryManager) AddVertexByString(query string) (model.Vertex, error) {
	responses, err := v.executeStringQuery(query)
	if err != nil {
		v.logger.Error("invalid query",
			gremerror.NewQueryError("AddVertexByString", query, err),
		)
		return nilVertex, err
	}

	var list model.VertexList

	for _, res := range responses {
		var vertPart model.VertexList
		// Create the resulting vertices from the query.
		err = jsonUnmarshal(res, &vertPart)
		if err != nil {
			v.logger.Error("vertices unmarshal",
				gremerror.NewUnmarshalError("AddVertexByString", res, err),
			)
			return nilVertex, err
		}

		list.Vertices = append(list.Vertices, vertPart.Vertices...)
	}

	if len(list.Vertices) > 0 {
		return list.Vertices[0], nil
	}
	return nilVertex, nil
}
