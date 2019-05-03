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
	"github.com/northwesternmutual/grammes/logging"
)

// GraphQueryManager has all the function related to interacting with the graph.
type GraphQueryManager struct {
	*queryManager
	*vertexQueryManager
	*miscQueryManager
	*schemaManager

	dialer gremconnect.Dialer
	logger logging.Logger
}

// NewGraphManager will give a manager to handle all
// graph interactions through the TinkerPop server.
func NewGraphManager(dialer gremconnect.Dialer, logger logging.Logger, executeRequest executor) *GraphQueryManager {
	g := &GraphQueryManager{
		queryManager: newQueryManager(dialer, logger, executeRequest),
	}

	g.vertexQueryManager = newVertexQueryManager(logger, g.ExecuteStringQuery)
	g.miscQueryManager = newMiscQueryManager(logger, g.ExecuteStringQuery)
	g.schemaManager = newSchemaManager(logger, g.ExecuteStringQuery)

	return g
}

// SetLogger will set the logger being used by all of the functions in the graph functions.
func (g *GraphQueryManager) SetLogger(newLogger logging.Logger) {
	g.logger = newLogger
	g.queryManager.logger = newLogger
	g.schemaManager.logger = newLogger
	g.miscQueryManager.logger = newLogger
	g.vertexQueryManager.addVertexQueryManager.logger = newLogger
	g.vertexQueryManager.getVertexQueryManager.logger = newLogger
}

// MiscQuerier returns the manager for miscellaneous queries.
func (g *GraphQueryManager) MiscQuerier() MiscQuerier {
	return g.miscQueryManager
}

// AddVertexQuerier returns the manager for adding vertices to the graph.
func (g *GraphQueryManager) AddVertexQuerier() AddVertexQuerier {
	return g.vertexQueryManager.addVertexQueryManager
}

// GetVertexQuerier returns the manager for getting vertices from the graph.
func (g *GraphQueryManager) GetVertexQuerier() GetVertexQuerier {
	return g.vertexQueryManager.getVertexQueryManager
}

// GetVertexIDQuerier returns the manager for getting vertex IDs.
func (g *GraphQueryManager) GetVertexIDQuerier() GetVertexIDQuerier {
	return g.vertexQueryManager.vertexIDQueryManager
}

// DropQuerier returns the manager for dropping elements from the graph.
func (g *GraphQueryManager) DropQuerier() DropQuerier {
	return g.vertexQueryManager.dropQueryManager
}

// VertexQuerier returns the manager for all vertex related queries.
func (g *GraphQueryManager) VertexQuerier() VertexQuerier {
	return g.vertexQueryManager
}

// ExecuteQuerier returns the manager for executing the raw queries.
func (g *GraphQueryManager) ExecuteQuerier() ExecuteQuerier {
	return g.queryManager
}

// SchemaQuerier returns the manager for executing the raw queries.
func (g *GraphQueryManager) SchemaQuerier() SchemaQuerier {
	return g.schemaManager
}