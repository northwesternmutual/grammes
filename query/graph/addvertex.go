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

package graph

import (
	"fmt"
)

// AddVertex will add vertex to graph
// Note:
// - AddVertex will not handle string formatting because
//   this function can accept any types. Any strings inputted
//   need to be formatted with single quotes beforehand.
// Signatures:
// AddVertex(interface{}...)
func (graph String) AddVertex(params ...interface{}) String {
	graph = graph.append(".addVertex(")

	switch params[0].(type) {
	case string:
		// custom property or other strings.
		graph = graph.append("\"" + params[0].(string) + "\"")
	default:
		// token or other types.
		graph = graph.append(fmt.Sprintf("%v", params[0]))
	}

	if len(params) > 1 {
		for _, p := range params[1:] {
			switch p.(type) {
			case string:
				// custom property or other strings.
				graph = graph.append(",\"" + p.(string) + "\"")
			default:
				// Token or other types.
				graph = graph.append(fmt.Sprintf(",%v", p))
			}
		}
	}

	graph = graph.append(")")

	return graph
}
