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

	"github.com/northwesternmutual/grammes/query/direction"
)

// http://tinkerpop.apache.org/docs/current/reference/#to-step

// To (step-modulator) similar to As() and By(). If a step is able to accept
// traversals or strings then To() is the means by which they are added.
// Note:
// - To does not handle any string formatting for you. If you wish to pass
//   a regular string as a parameter you must add single-quotes around it beforehand.
// Signatures:
// To(Direction, ...string)
// To(string)
// To(*String (Traversal))
// To(string Vertex)
func (g String) To(first interface{}, extraStrings ...string) String {
	g = g.append(".to(")

	switch first.(type) {
	case string:
		g = g.append(first.(string))
	case String:
		g = g.append(first.(String).Raw().String())
	case direction.Direction:
		g = g.append(fmtStr("%v", first.(direction.Direction)))
	default:
		fmt.Println("Type mismatch used in To()")
	}

	if len(extraStrings) > 0 {
		for _, v := range extraStrings {
			g = g.append(fmtStr(",%v", v))
		}
	}

	g = g.append(")")

	return g
}

// ToE (step-modulator) is a part of To()
// Signatures:
// ToE(Direction, string)
func (g String) ToE(dir direction.Direction, str string) String {
	g = g.append(fmtStr(".toE(%v, \"%v\")", dir, str))

	return g
}

// ToV (step-modulator) is a part of To()
// Signatures:
// ToV(Direction)
func (g String) ToV(dir direction.Direction) String {
	g.AddStep("toV", dir)

	return g
}

// ToVId can be used to make a string query that will take a vertex id as a parameter,
// and can be used to point an edge towards this vertex ID.
func (g String) ToVId(vertexID int64) String {
	g = g.append(fmtStr(".to(V().hasId(%v))", vertexID))

	return g
}
