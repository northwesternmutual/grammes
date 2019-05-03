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

// http://tinkerpop.apache.org/docs/current/reference/#as-step

// As (step modulator) similar to By() & Option(). With As(), it is
// possible to provide a label to the step that can later be accessed by
// steps and data structures that make use of such labels.
// Signatures:
// As(string)
// As(string, string...)
func (g String) As(labels ...string) String {
	g = g.append(".as(")

	if len(labels) < 1 {
		fmt.Println("Not enough parameters to use As()")
	}

	if len(labels) > 0 {
		g = g.append("\"" + labels[0] + "\"")
	}

	if len(labels) > 1 {
		for _, v := range labels[1:] {
			g = g.append(",\"" + v + "\"")
		}
	}

	g = g.append(")")

	return g
}
