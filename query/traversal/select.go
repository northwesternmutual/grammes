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

import (
	"fmt"

	"github.com/northwesternmutual/grammes/query/pop"
)

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
	g = g.append(".select(")

	switch first.(type) {
	case string:
		g = g.append(first.(string))
	case String:
		g = g.append(first.(String).String())
	case pop.Pop:
		g = g.append(fmtStr("%v", first.(pop.Pop)))
	default:
		fmt.Println("Mismatch types used for Select()")
	}

	if len(extras) > 0 {
		for _, v := range extras {
			switch v.(type) {
			case String:
				g = g.append(fmtStr(",%v", v.(String).String()))
			case string:
				g = g.append(fmtStr(",%v", v.(string)))
			default:
				fmt.Println("Mismatch types used for Select()")
			}
		}
	}

	g = g.append(")")

	return g
}
