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
	"errors"
	"fmt"
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

func (m *miscQueryManager) SetVertexProperty(id interface{}, keyAndVals ...interface{}) error {
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

	if len(responses) == 0 {
		return 0, gremerror.ErrEmptyResponse
	}

	var rawResp model.IDList

	err = jsonUnmarshal(responses[0], &rawResp)
	if err != nil {
		m.logger.Error("unmarshal",
			gremerror.NewUnmarshalError("VertexCount", responses[0], err),
		)
		return 0, err
	}

	if len(rawResp.IDs) == 0 {
		return 0, errors.New(fmt.Sprintf("invalid response %s", string(responses[0])))
	}
	v, ok := rawResp.IDs[0].(map[string]interface{})
	if !ok {
		return 0, errors.New(fmt.Sprintf("invalid response %s", string(responses[0])))
	}
	count, ok := v["@value"].(float64)
	if !ok {
		return 0, errors.New(fmt.Sprintf("invalid response %s", string(responses[0])))
	}

	return int64(count), nil
}
