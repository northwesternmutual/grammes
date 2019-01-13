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

// http://tinkerpop.apache.org/docs/current/reference/#sample-step

// Sample (step-modulator) is useful for sampling some number
// of traversers previous in the traversal.
// Signatures:
// Sample(int)
// Sample(Scope, int)
func (g String) Sample(params ...interface{}) String {
	if len(params) < 1 {
		fmt.Println("Not enough arguments to call Sample()")
		return g
	} else if len(params) > 2 {
		fmt.Println("Too many arguments to call Sample()")
	}

	// append the scope or int
	g = g.append(fmtStr(".sample(%v", params[0]))

	// if there is another parameter then append it.
	if len(params) > 1 {
		g = g.append(fmtStr(",%v", params[1]))
	}

	// append the ending parenthesis.
	g = g.append(")")

	return g
}
