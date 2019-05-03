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
Package scope contains the Scope object to control relations of a graph traversal.

See: http://tinkerpop.apache.org/javadocs/3.3.3/core/org/apache/tinkerpop/gremlin/process/traversal/Scope.html

Scopes can alter how the steps behave given the graph traversal.

A note about Scope:

This object implements the Parameter interface used by graph traversals.
*/
package scope

// http://tinkerpop.apache.org/javadocs/3.3.3/core/org/apache/tinkerpop/gremlin/process/traversal/Scope.html

// Scope alters the manner in which the step will behave
// in relation to how the traversers are processed.
type Scope string

const (
	// Local informs the step to operate on
	// the current object in the step.
	Local Scope = "local"
	// Global informs the step to operate on
	// the entire traversal.
	Global Scope = "global"
)

func (s Scope) String() string {
	return string(s)
}
