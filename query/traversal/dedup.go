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

// http://tinkerpop.apache.org/docs/current/reference/#dedup-step

// Dedup (filter) repeatedly seen objects are removed from the traversal stream.
// Signatures:
// Dedup(Scope, ...string)
// Dedup(...string)
func (g String) Dedup(params ...interface{}) String {
	// g = g.append(".dedup(")

	// if len(params) > 0 {
	// 	switch params[0].(type) {
	// 	case string:
	// 		g = g.append("\"" + params[0].(string) + "\"")
	// 	case scope.Scope:
	// 		g = g.append(params[0].(scope.Scope).String())
	// 	default:
	// 		fmt.Println("mismatching types used for Dedup()")
	// 	}
	// }

	// if len(params) > 1 {
	// 	for _, v := range params[1:] {
	// 		g = g.append(",\"" + v.(string) + "\"")
	// 	}
	// }

	// g = g.append(")")

	// if len(params) > 0 {
	// 	g.AddStep("dedup", params...)
	// } else {
	// 	g.AddStep("dedup")
	// }

	g.AddStep("dedup", params...)

	return g
}
