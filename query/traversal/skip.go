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

// http://tinkerpop.apache.org/docs/current/reference/#skip-step

// Skip (filter) is analogous to Range() save that the higher end range is set to -1.
// Signatures:
// Skip(float32)
// Skip(Scope, float32)
func (g String) Skip(first interface{}, extraFloat ...float32) String {
	// g = g.append(fmtStr(".skip(%v", first))

	// if len(extraFloat) > 0 {
	// 	g = g.append(fmtStr(",%v", extraFloat[0]))
	// }

	// g = g.append(")")

	var params []interface{}

	params = append(params, first)

	if len(extraFloat) == 1 {
		params = append(params, extraFloat[0])
	}

	g.AddStep("skip", params...)

	return g
}
