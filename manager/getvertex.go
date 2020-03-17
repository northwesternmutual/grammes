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
	"fmt"
	"strconv"

	"github.com/northwesternmutual/grammes/gremerror"
	"github.com/northwesternmutual/grammes/logging"
	"github.com/northwesternmutual/grammes/model"
	"github.com/northwesternmutual/grammes/query"
	"github.com/northwesternmutual/grammes/query/traversal"
)

type getVertexQueryManager struct {
	logger             logging.Logger
	executeStringQuery stringExecutor
}

func newGetVertexQueryManager(logger logging.Logger, executor stringExecutor) *getVertexQueryManager {
	return &getVertexQueryManager{
		logger:             logger,
		executeStringQuery: executor,
	}
}

func (c *getVertexQueryManager) VerticesByString(query string) ([]model.Vertex, error) {
	// Query the gremlin server with the given traversal.
	responses, err := c.executeStringQuery(query)
	if err != nil {
		c.logger.Error("invalid query",
			gremerror.NewQueryError("Vertices", query, err),
		)
		return nil, err
	}

	var vertices model.VertexList

	for _, res := range responses {
		var vertPart model.VertexList
		// Unmarshal the response into the structs.
		err = jsonUnmarshal(res, &vertPart)
		if err != nil {
			c.logger.Error("vertices unmarshal",
				gremerror.NewUnmarshalError("Vertices", res, err),
			)
			return nil, err
		}

		vertices.Vertices = append(vertices.Vertices, vertPart.Vertices...)
	}

	c.logger.Debug("RESPONSE INFO", map[string]interface{}{"FULL LENGTH": len(vertices.Vertices)})

	return vertices.Vertices, nil
}

// Vertices will gather any vertices and return them
// based on the fed in traversal query.
func (c *getVertexQueryManager) VerticesByQuery(query query.Query) ([]model.Vertex, error) {
	vertices, err := c.VerticesByString(query.String())
	if err != nil {
		c.logger.Error("error gathering vertices",
			gremerror.NewGrammesError("VerticesByQuery", err),
		)
	}
	return vertices, err
}

// AllVertices will return every vertex on the graph
// and return them in a structured format.
func (c *getVertexQueryManager) AllVertices() ([]model.Vertex, error) {
	// Query the graph database for all vertices.
	vertices, err := c.VerticesByString("g.V()")
	if err != nil {
		c.logger.Error("error gathering vertices",
			gremerror.NewGrammesError("AllVertices", err),
		)
		return nil, err
	}

	return vertices, nil
}

// Vertex will get a Vertex with the following
// ID assigned to it. This ID is unique to every
// vertex on the graph. This is the best way of finding
// vertices without any conflicting labels or properties.
func (c *getVertexQueryManager) VertexByID(id interface{}) (model.Vertex, error) {
	// Query the graph for a vertex with this ID.
	vertices, err := c.VerticesByString("g.V().hasId(" + fmt.Sprint(id) + ")")
	if err != nil {
		c.logger.Error("error gathering vertices",
			gremerror.NewGrammesError("VerticesByID", err),
		)
		return nilVertex, err
	}

	if len(vertices) == 0 {
		return nilVertex, gremerror.NewGrammesError("VertexByID", gremerror.ErrEmptyResponse)
	}
	// There should only be one vertex with this unique ID.
	return vertices[0], nil
}

func (c *getVertexQueryManager) Vertices(label string, properties ...interface{}) ([]model.Vertex, error) {
	if len(properties)%2 != 0 {
		c.logger.Error("number of parameters ["+strconv.Itoa(len(properties))+"]",
			gremerror.NewGrammesError("AddVertex", gremerror.ErrOddNumberOfParameters),
		)
		return nil, gremerror.ErrOddNumberOfParameters
	}

	query := traversal.NewTraversal().V().HasLabel(label)
	for i := 0; i < len(properties); i += 2 {
		query = query.Has(properties[i], properties[i+1])
	}

	vertices, err := c.VerticesByString(query.String())
	if err != nil {
		c.logger.Error("error gathering vertices",
			gremerror.NewGrammesError("Vertices", err),
		)
	}

	return vertices, err
}
