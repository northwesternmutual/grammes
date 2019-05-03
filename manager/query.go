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
	"github.com/northwesternmutual/grammes/gremconnect"
	"github.com/northwesternmutual/grammes/gremerror"
	"github.com/northwesternmutual/grammes/logging"
	"github.com/northwesternmutual/grammes/query"
)

// Query handles the querying actions to the server.
type queryManager struct {
	dialer         gremconnect.Dialer
	logger         logging.Logger
	executeRequest executor
}

// NewQueryManager returns a new Query Manager that
// implements the QueryManager interface.
func newQueryManager(dialer gremconnect.Dialer, logger logging.Logger, executor executor) *queryManager {
	return &queryManager{
		dialer:         dialer,
		logger:         logger,
		executeRequest: executor,
	}
}

func (m *queryManager) setLogger(newLogger logging.Logger) {
	m.logger = newLogger
}

// ExecuteQuery takes a query object to form a
// request to the gremlin server after turning it
// into a string.
func (m *queryManager) ExecuteQuery(query query.Query) ([][]byte, error) {
	return m.ExecuteBoundStringQuery(query.String(), map[string]string{}, map[string]string{})
}

// ExecuteStringQuery takes a string query and
// uses it to make a request to the gremlin server.
func (m *queryManager) ExecuteStringQuery(query string) ([][]byte, error) {
	return m.ExecuteBoundStringQuery(query, map[string]string{}, map[string]string{})
}

// Query Bindings:
// https://www.codeigniter.com/userguide3/database/queries.html#query-bindings

// ExecuteBoundQuery takes a query object and bindings to allow
// for simplified queries to the gremlin server.
func (m *queryManager) ExecuteBoundQuery(query query.Query, bindings, rebindings map[string]string) ([][]byte, error) {
	return m.ExecuteBoundStringQuery(query.String(), bindings, rebindings)
}

// ExecuteBoundStringQuery uses bindings and rebindings to allow
// for simplified queries to the gremlin server.
func (m *queryManager) ExecuteBoundStringQuery(query string, bindings, rebindings map[string]string) ([][]byte, error) {
	if m.dialer.IsDisposed() {
		return nil, gremerror.ErrDisposedConnection
	}

	// log the command that will be executed.
	m.logger.PrintQuery(query)

	return m.executeRequest(query, bindings, rebindings)
}
