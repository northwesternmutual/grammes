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

// http://tinkerpop.apache.org/docs/current/reference/#where-step

// Where (filter) filters the current objects based on either the object
// itself or the path history of the object.
// Notes:
// - Where does not handle string formatting. If you wish to input
//   a string has a parameter you must surround it with single-quotes beforehand.
// Signatures:
// Where(string (P))
// Where(string, string (P))
// Where(*String (Traversal))
func (g String) Where(first interface{}, extra ...string) String {
	g = g.append(".where(")

	switch first.(type) {
	case string:
		g = g.append(first.(string))
	case String: // Where(*String (Traversal))
		g = g.append(fmtStr("%v)", first.(String).String()))
		return g
	default:
		fmt.Println("Mismatching types used for Where()")
	}

	if len(extra) > 0 {
		for _, v := range extra {
			g = g.append("," + v)
		}
	}

	g = g.append(")")

	return g
}
