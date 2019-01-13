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
	"encoding/json"

	"github.com/northwesternmutual/grammes/gremerror"
)

// UnmarshalVertexList is a utility to unmarshal a list
// or array of vertices properly.
func UnmarshalVertexList(data []byte) ([]Vertex, error) {
	var list VertexList

	if err := json.Unmarshal(data, &list); err != nil {
		return nil, gremerror.NewUnmarshalError("UnmarshalVertexList", data, err)
	}

	return list.Vertices, nil
}

// UnmarshalEdgeList is a utility to unmarshal a list
// or array of edges properly.
func UnmarshalEdgeList(data []byte) ([]Edge, error) {
	var list EdgeList

	if err := json.Unmarshal(data, &list); err != nil {
		return nil, gremerror.NewUnmarshalError("UnmarshalEdgeList", data, err)
	}

	return list.Edges, nil
}

// UnmarshalIDList is a utility to unmarshal a list
// or array of IDs properly.
func UnmarshalIDList(data []byte) ([]ID, error) {
	var list IDList

	if err := json.Unmarshal(data, &list); err != nil {
		return nil, gremerror.NewUnmarshalError("UnmarshalIDList", data, err)
	}

	return list.IDs, nil
}
