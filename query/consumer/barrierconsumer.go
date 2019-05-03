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
Package consumer contains the object to control how barriers emit their values.

See: http://tinkerpop.apache.org/javadocs/3.2.1/core/org/apache/tinkerpop/gremlin/process/traversal/SackFunctions.Barrier.html

BarrierConsumer is any step that requires all left traversers to be processed
prior to emitting result traversers to the right.
This acts like a barrier (hence the name).

A note about BarrierConsumer:

This object implements the Parameter interfaces used by graph traversals.
*/
package consumer

// http://tinkerpop.apache.org/javadocs/3.2.1/core/org/apache/tinkerpop/gremlin/process/traversal/SackFunctions.Barrier.html

// BarrierConsumer is any step that requires all left traversers to be processed
// prior to emitting result traversers to the right.
type BarrierConsumer string

const (
	// NormSack assures an evenly distributed sack of 1.0
	NormSack BarrierConsumer = "normSack"
)

func (c BarrierConsumer) String() string {
	return string(c)
}
