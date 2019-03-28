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
	"strings"

	"github.com/northwesternmutual/grammes/query/traversal"

	"github.com/northwesternmutual/grammes/gremerror"
	"github.com/northwesternmutual/grammes/logging"
	"github.com/northwesternmutual/grammes/model"
	"github.com/northwesternmutual/grammes/query"
)

type vertexIDQueryManager struct {
	logger             logging.Logger
	executeStringQuery stringExecutor
}

func newVertexIDQueryManager(logger logging.Logger, executor stringExecutor) *vertexIDQueryManager {
	return &vertexIDQueryManager{
		logger:             logger,
		executeStringQuery: executor,
	}
}

// VertexIDsByString executes a string query and unmarshals the
// IDs for the user.
func (v *vertexIDQueryManager) VertexIDsByString(q string) ([]int64, error) {
	if !strings.HasSuffix(q, ".id()") {
		q += ".id()"
	}

	// retrieve all the vertices from the graph.
	responses, err := v.executeStringQuery(q)
	if err != nil {
		v.logger.Error("invalid query",
			gremerror.NewQueryError("VertexIDs", q, err),
		)
		return nil, err
	}

	var rawIDs model.IDList

	for _, res := range responses {
		var idPart model.IDList
		err = jsonUnmarshal(res, &idPart)
		if err != nil {
			v.logger.Error("id unmarshal",
				gremerror.NewUnmarshalError("VertexIDs", res, err),
			)
			return nil, err
		}

		rawIDs.IDs = append(rawIDs.IDs, idPart.IDs...)
	}

	var ids []int64

	for _, id := range rawIDs.IDs {
		ids = append(ids, id.Value)
	}

	return ids, nil
}

// VertexIDsByQuery will take a query and execute it. Then it will
// run through and extract all the vertex IDs matching the
// traversal and return them in an array of int64.
func (v *vertexIDQueryManager) VertexIDsByQuery(query query.Query) ([]int64, error) {
	ids, err := v.VertexIDsByString(query.String())
	if err != nil {
		v.logger.Error("error gathering IDs",
			gremerror.NewGrammesError("VertexIDsByQuery", err),
		)
	}

	return ids, err
}

// VertexIDs takes the label and optional properties to retrieve
// the IDs desired from the graph.
func (v *vertexIDQueryManager) VertexIDs(label string, properties ...interface{}) ([]int64, error) {
	if len(properties)%2 != 0 {
		v.logger.Error("number of parameters ["+strconv.Itoa(len(properties))+"]",
			gremerror.NewGrammesError("VertexIDs", gremerror.ErrOddNumberOfParameters),
		)
		return nil, gremerror.ErrOddNumberOfParameters
	}

	query := traversal.NewTraversal().V().HasLabel(label)

	for i := 0; i < len(properties); i += 2 {
		query.AddStep("has", properties[i], properties[i+1])
	}

	ids, err := v.VertexIDsByString(query.String())
	if err != nil {
		v.logger.Error("error gathering IDs",
			gremerror.NewGrammesError("VertexIDs", err),
		)
	}

	return ids, err
}
