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
	"strings"

	"github.com/northwesternmutual/grammes/query/traversal"
	"github.com/northwesternmutual/grammes/gremerror"
	"github.com/northwesternmutual/grammes/logging"
	"github.com/northwesternmutual/grammes/query"
)

type dropQueryManager struct {
	logger             logging.Logger
	executeStringQuery stringExecutor
}

func newDropQueryManager(logger logging.Logger, executor stringExecutor) *dropQueryManager {
	return &dropQueryManager{
		logger:             logger,
		executeStringQuery: executor,
	}
}

func (v *dropQueryManager) DropVertexLabel(label string) error {
	query := traversal.NewTraversal().V().HasLabel(label).Drop()
	if _, err := v.executeStringQuery(query.String()); err != nil {
		v.logger.Error("invalid query",
			gremerror.NewQueryError("DropVertexLabel", query.String(), err),
		)
		return err
	}

	return nil
}

func (v *dropQueryManager) DropVertexByID(ids ...interface{}) error {
	var err error
	for _, id := range ids {
		query := traversal.NewTraversal().V().HasID(id).Drop()
		if _, err = v.executeStringQuery(query.String()); err != nil {
			v.logger.Error("invalid query",
				gremerror.NewQueryError("DropVerticesByID", query.String(), err),
			)
			return err
		}
	}

	return nil
}

func (v *dropQueryManager) DropVerticesByString(q string) error {
	if !strings.HasSuffix(q, "drop()") {
		q += ".drop()"
	}
	
	_, err := v.executeStringQuery(q)
	if err != nil {
		v.logger.Error("invalid query",
			gremerror.NewQueryError("DropVerticesByString", q, err),
		)
	}
	return err
}

func (v *dropQueryManager) DropVerticesByQuery(q query.Query) error {
	err := v.DropVerticesByString(q.String())
	if err != nil {
		v.logger.Error("invalid query",
			gremerror.NewQueryError("DropVerticesByQuery", q.String(), err),
		)
	}
	return err
}