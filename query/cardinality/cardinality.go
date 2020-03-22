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
Package cardinality contains the object that describes number of relationship occurrences for objects.

See:
	TinkerPop: http://tinkerpop.apache.org/javadocs/3.3.3/core/org/apache/tinkerpop/gremlin/structure/VertexProperty.Cardinality.html
	Wikipedia: https://en.wikipedia.org/wiki/Cardinality

Cardinality describes the maximum number of possible relationships.

A note about Cardinality:

This object implements the Parameter interfaces used by graph traversals.
*/
package cardinality

// Tinkerpop:
// http://tinkerpop.apache.org/javadocs/3.3.3/core/org/apache/tinkerpop/gremlin/structure/VertexProperty.Cardinality.html

// Wikipedia:
// https://en.wikipedia.org/wiki/Cardinality

// Cardinality describes maximum number of possible relationship
// occurrences for an entity participating in a given relationship type.
type Cardinality string

const (
	// List allows an arbitrary number of
	// values per element for such key.
	List Cardinality = "list"
	// Set allows multiple values but no
	// duplicate values per element for such key.
	Set Cardinality = "set"
	// Single allows at most one
	// value per element for such key
	Single Cardinality = "single"
)

// String will convert Cardinality to a string
func (c Cardinality) String() string {
	return string(c)
}
