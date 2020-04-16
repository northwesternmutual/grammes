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

	"github.com/northwesternmutual/grammes/gremerror"
	"github.com/northwesternmutual/grammes/logging"
	"github.com/northwesternmutual/grammes/query/cardinality"
	"github.com/northwesternmutual/grammes/query/datatype"
	"github.com/northwesternmutual/grammes/query/graph"
	"github.com/northwesternmutual/grammes/query/multiplicity"
)

type schemaManager struct {
	logger             logging.Logger
	executeStringQuery stringExecutor
}

func newSchemaManager(logger logging.Logger, executor stringExecutor) *schemaManager {
	return &schemaManager{
		logger:             logger,
		executeStringQuery: executor,
	}
}

// AddEdgeLabel adds the edge label to the
// graph directly. This method returns the schema id
// of the edge label added.
func (s *schemaManager) AddEdgeLabel(multi multiplicity.Multiplicity, label string) (id interface{}, err error) {
	var (
		data  [][]byte
		query = graph.NewGraph().OpenManagement().MakeEdgeLabel(label).Multiplicity(multi).Make()
	)

	if data, err = s.executeStringQuery(query.String()); err != nil {
		s.logger.Error("invalid query",
			gremerror.NewQueryError("AddEdgeLabel", query.String(), err),
		)
		return
	}

	if id, err = unmarshalID(data); err != nil {
		s.logger.Error("id unmarshal",
			gremerror.NewUnmarshalError("AddEdgeLabel", data[0], err),
		)
	}

	s.logger.Debug("Appended Edge Label", map[string]interface{}{
		"NewLabel":     label,
		"Multiplicity": multi.String(),
		"ID":           id,
	})

	return id, err
}

// AddEdgeLabels does the same thing as AddEdgeLabel
// but with the ability to do multiple labels at a
// time. This function is called similarly to your
// favorite logger.
func (s *schemaManager) AddEdgeLabels(multiplicityAndLabels ...interface{}) (ids []interface{}, err error) {
	if len(multiplicityAndLabels)%2 != 0 {
		s.logger.Error(fmt.Sprintf("number of parameters [%d]", len(multiplicityAndLabels)),
			gremerror.NewGrammesError("AddEdgeLabels", gremerror.ErrOddNumberOfParameters),
		)
		return nil, gremerror.ErrOddNumberOfParameters
	}

	var (
		multi multiplicity.Multiplicity
		label string
		id    interface{}
		ok    bool
	)

	for i := 0; i < len(multiplicityAndLabels); i += 2 {
		if multi, ok = multiplicityAndLabels[i].(multiplicity.Multiplicity); !ok {
			return nil, fmt.Errorf("invalid multiplicity [%v]", multiplicityAndLabels[i])
		}
		if label, ok = multiplicityAndLabels[i+1].(string); !ok {
			return nil, fmt.Errorf("invalid label [%v]", multiplicityAndLabels[i+1])
		}
		if id, err = s.AddEdgeLabel(multi, label); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}

// AddPropertyKey adds the edge label to the
// graph directly. This method returns the schema id
// of the edge label added.
func (s *schemaManager) AddPropertyKey(propertyName string, datatype datatype.DataType, cardinality cardinality.Cardinality) (id interface{}, err error) {
	var (
		data  [][]byte
		query = graph.NewGraph().OpenManagement().MakePropertyKey(propertyName, datatype, cardinality).Make()
	)

	if data, err = s.executeStringQuery(query.String()); err != nil {
		s.logger.Error("invalid query",
			gremerror.NewQueryError("AddPropertyKey", query.String(), err),
		)
		return
	}
	if id, err = unmarshalID(data); err != nil {
		s.logger.Error("id unmarshal",
			gremerror.NewUnmarshalError("AddPropertyKey", data[0], err),
		)
	}

	return
}

// Commit will take all of your schema changes
// and apply them to the schema once they are ready.
func (s *schemaManager) CommitSchema() ([][]byte, error) {
	data, err := s.executeStringQuery("graph.openManagement().commit()")
	if err != nil {
		s.logger.Error("invalid query",
			gremerror.NewQueryError("Commit", "graph.openManagement().commit()", err),
		)
		return nil, err
	}

	return data, nil
}
