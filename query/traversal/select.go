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

package traversal

import "fmt"

// http://tinkerpop.apache.org/docs/current/reference/#select-step

// Select (map) Can go back in a traversal in the previously seen area of computation.
// Note:
// - Select does not handle any string formatting. If you wish to input a normal string
//   as a parameter then you must add single-quotes around the string beforehand.
// Signatures:
// Select(string (Column))
// Select(Pop, string)
// Select(string)
// Select(string, string, ...string)
// Select(*String (Traversal))
// Select(Pop, *String (Traversal))
func (g String) Select(first interface{}, extras ...interface{}) String {
	var params []interface{}

	params = append(params, first)

	for _, e := range extras {
		params = append(params, e)
	}

	for _, p := range params {
		switch t := p.(type) {
		case string:
		case Parameter:
		default:
			fmt.Printf("invalid parameter: [%v]\n", t)
			params = nil
		}
	}

	g.AddStep("select", params...)

	return g
}