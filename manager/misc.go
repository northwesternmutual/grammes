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
	"github.com/northwesternmutual/grammes/query/traversal"
)

type miscQueryManager struct {
	logger             logging.Logger
	executeStringQuery stringExecutor
}

func newMiscQueryManager(logger logging.Logger, execute stringExecutor) *miscQueryManager {
	return &miscQueryManager{
		executeStringQuery: execute,
		logger:             logger,
	}
}

func (m *miscQueryManager) DropAll() error {
	_, err := m.executeStringQuery("g.V().drop()")
	return err
}

func (m *miscQueryManager) SetVertexProperty(id int64, keyAndVals ...interface{}) error {
	if len(keyAndVals)%2 != 0 {
		m.logger.Error("number of parameters ["+strconv.Itoa(len(keyAndVals))+"]",
			gremerror.NewGrammesError("SetVertexProperty", gremerror.ErrOddNumberOfParameters),
		)
		return gremerror.ErrOddNumberOfParameters
	}

	query := traversal.NewTraversal().V().HasID(id)
	for i := 0; i < len(keyAndVals); i += 2 {
		query.AddStep("property", keyAndVals[i], keyAndVals[i+1])
	}

	if _, err := m.executeStringQuery(query.String()); err != nil {
		m.logger.Error("invalid query",
			gremerror.NewQueryError("SetVertexProperty", query.String(), err),
		)
		return err
	}

	return nil
}

// VertexCount retrieves the number of vertices
// that are currently on the graph as an int64.
func (m *miscQueryManager) VertexCount() (int64, error) {
	// Query the graph for the count using IDs.
	query := traversal.NewTraversal().V().Count()

	responses, err := m.executeStringQuery(query.String())
	if err != nil {
		m.logger.Error("VertexCount",
			gremerror.NewQueryError("VertexCount", query.String(), err),
		)
		return 0, err
	}

	var resultingIDs model.IDList

	for _, res := range responses {
		var rawIDs model.IDList

		err = jsonUnmarshal(res, &rawIDs)
		if err != nil {
			m.logger.Error("id unmarshal",
				gremerror.NewUnmarshalError("VertexCount", res, err),
			)
			return 0, err
		}

		resultingIDs.IDs = append(resultingIDs.IDs, rawIDs.IDs...)
	}

	if len(resultingIDs.IDs) == 0 {
		return 0, gremerror.ErrEmptyResponse
	}

	return resultingIDs.IDs[0].Value, nil
}
