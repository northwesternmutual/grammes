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
	"github.com/northwesternmutual/grammes/query/scope"
)

// http://tinkerpop.apache.org/docs/current/reference/#sum-step

// Sum (map) operates on a stream of numbers and sums the numbers
// together to yield a double.
// Signatures:
// Sum()
// Sum(Scope)
func (g String) Sum(scopes ...scope.Scope) String {
	var i scope.Scope

	if len(scopes) == 1 {
		i = scopes[0]
	}

	g.AddStep("sum", i)

	// if len(scope) < 1 {
	// 	g.AddStep("sum")
	// } else {
	// 	g = g.append(fmtStr(".sum(%v)", scope[0]))
	// }

	return g
}
