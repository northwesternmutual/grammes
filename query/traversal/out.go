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

// Out moves to the outgoing adjacent vertices given the edge labels.
func (g String) Out(labels ...string) String {
	var p []interface{}

	for _, l := range labels {
		p = append(p, l)
	}

	g.AddStep("out", p...)

	return g
}

// OutE moves to the outgoing incident edges given the edge labels.
func (g String) OutE(labels ...string) String {
	var p []interface{}

	for _, l := range labels {
		p = append(p, l)
	}

	g.AddStep("outE", p...)

	return g
}

// OutV moves to the outgoing vertex.
func (g String) OutV() String {
	g.AddStep("outV")

	return g
}
