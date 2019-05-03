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

// http://tinkerpop.apache.org/docs/current/reference/#choose-step

// Choose (branch) routes the current traverser to a
// particular traversal branch option.
// Note:
// Signatures:
// Choose(string (Function))
// Choose(string (Predicate), *String (Traversal))
// Choose(string (Predicate), *String (Traversal), *String (Traversal))
// Choose(*String (Traversal), *String (Traversal))
// Choose(*String (Traversal), *String (Traversal), *String (Traversal))
// Choose(*String (Traversal))
func (g String) Choose(first interface{}, optTraversals ...String) String {
	g = g.append(".choose(")

	switch first.(type) {
	case string:
		// Choose(string (Function)...
		// Choose(string (Predicate)...
		g = g.append(first.(string))
	case String:
		// Choose(*String (Traversal)...
		g = g.append(first.(String).String())
	default:
		fmt.Println("mismatching type used in Choose()")
	}

	if len(optTraversals) > 0 {
		for _, v := range optTraversals {
			g = g.append("," + v.String())
		}
	}

	g = g.append(")")

	return g
}
