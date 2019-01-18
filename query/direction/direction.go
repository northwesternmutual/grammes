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

/*
Package direction contains the object to denote the direction of an edge.

See: http://tinkerpop.apache.org/javadocs/3.3.3/core/org/apache/tinkerpop/gremlin/structure/Direction.html

Direction controls the direction of an edge or location of a vertex on an edge.

A note about Direction:

This object implements the Parameter interfaces used by graph traversals.
*/
package direction

// Tinkerpop:
// http://tinkerpop.apache.org/javadocs/3.3.3/core/org/apache/tinkerpop/gremlin/structure/Direction.html

// Direction is used to denote the direction of an Edge
// or Location of a Vertex on an Edge.
type Direction string

const (
	// In refers to an incoming direction.
	In Direction = "IN"
	// Out refers to an outgoing direction.
	Out Direction = "OUT"
	// Both refers to either direction (IN or OUT).
	Both Direction = "BOTH"
)

func (d Direction) String() string {
	return string(d)
}
