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

// http://tinkerpop.apache.org/docs/current/reference/#valuemap-step

// ValueMap (map) yields a map representation of the properties of an element.
// Signatures:
// ValueMap()
// ValueMap(bool, string...)
// ValueMap(string...)
func (g String) ValueMap(boolOrStrings ...interface{}) String {
	if len(boolOrStrings) < 1 {
		g.AddStep("valueMap") // empty command if there are no parameters given.
		return g
	}

	// append the command beginning along with the first parameter value.
	switch boolOrStrings[0].(type) {
	case string:
		g = g.append(".valueMap(\"" + boolOrStrings[0].(string) + "\"")
	default:
		g = g.append(fmtStr(".valueMap(%v", boolOrStrings[0]))
	}

	// append the rest of the parameters
	if len(boolOrStrings) > 1 {
		for _, v := range boolOrStrings[1:] {
			g = g.append(fmtStr(",\"%v\"", v))
		}
	}

	g = g.append(")")

	return g
}
