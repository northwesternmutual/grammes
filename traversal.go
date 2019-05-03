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

package grammes

import (
	"github.com/northwesternmutual/grammes/query/graph"
	"github.com/northwesternmutual/grammes/query/traversal"
)

// The graph's object to traverse
// the nodes/vertices on the graph.

// Traversal will return a new traversal string
// ready to use when executing a query.
func Traversal() traversal.String {
	return traversal.NewTraversal()
}

// CustomTraversal could be used when you need to specifically
// need to change some property of the traversal.
// This can be something such as:
//  // ==> graph.traversal().withoutStrategies(LazyBarrierStrategy)
func CustomTraversal(q string) traversal.String {
	return traversal.NewCustomTraversal(q)
}

// The graph's verbose traversal
// object to find nodes/vertices.

// VerboseTraversal will return a new graph traversal
// string ready to use when executing a query.
func VerboseTraversal() graph.String {
	return graph.String("graph")
}
