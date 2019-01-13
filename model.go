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

package grammes

import (
	"github.com/northwesternmutual/grammes/gremconnect"
	"github.com/northwesternmutual/grammes/model"
)

var (
	// NewWebSocketDialer returns websocket with established connection.
	NewWebSocketDialer = gremconnect.NewWebSocketDialer
	// NewVertexStruct returns a vertex struct meant for adding it.
	NewVertexStruct = model.NewVertexStruct
	// NewPropertyStruct returns a property struct meant for adding it to a vertex.
	NewPropertyStruct = model.NewPropertyStruct
)

// Localhost is used when connecting to a local Gremlin server.
const Localhost = "ws://127.0.0.1:8182/gremlin"

// Vertex is used to get quick access
// to the model.Vertex without having to
// import it everywhere in the grammes package.
type Vertex = model.Vertex

// VertexValue is used to get quick access
// to the model.VertexValue without having to
// import it everywhere in the grammes package.
type VertexValue = model.VertexValue

// Edge is used to get quick access
// to the model.Edge without having to
// import it everywhere in the grammes package.
type Edge = model.Edge

// Property is used to get quick access
// to the model.Property without having to
// import it everywhere in the grammes package.
type Property = model.Property

// PropertyMap is used to get quick access
// to the model.PropertyMap without having to
// import it everywhere in the grammes package.
type PropertyMap = model.PropertyMap

// PropertyValue is used to get quick access
// to the model.PropertyValue without having to
// import it everywhere in the grammes package.
type PropertyValue = model.PropertyValue

// ID is used to get quick access
// to the model.ID without having to
// import it everywhere in the grammes package.
type ID = model.ID

// APIData is used to get quick access
// to the model.APIData without having to
// import it everywhere in the grammes package.
type APIData = model.APIData

// Data is used to get quick access
// to the model.Data without having to
// import it everywhere in the grammes package.
type Data = model.Data

// VertexList is used to get quick access
// to the model.VertexList without having to
// import it everywhere in the grammes package.
type VertexList = model.VertexList

// EdgeList is used to get quick access
// to the model.EdgeList without having to
// import it everywhere in the grammes package.
type EdgeList = model.EdgeList

// IDList is used to get quick access
// to the model.IDList without having to
// import it everywhere in the grammes package.
type IDList = model.IDList

// PropertyList is used to get quick access
// to the model.PropertyList without having to
// import it everywhere in the grammes package.
type PropertyList = model.PropertyList
