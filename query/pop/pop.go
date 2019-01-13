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

package pop

// http://tinkerpop.apache.org/javadocs/3.3.3/core/org/apache/tinkerpop/gremlin/process/traversal/Pop.html

// Pop is used to determine whether the first value, last value
// or all values are gathered. Not that mixed will return results
// as a list if there are multiple or as a singleton if only one
// object in that path is labeled.
type Pop string

const (
	// First is the first item in an ordered collection
	// Such as collection[0]
	First Pop = "first"
	// Last is the last item in an ordered collection
	// Such as collection[collection.size()-1]
	Last Pop = "last"
	// All gets all the items and returns them as a list.
	All Pop = "all"
	// Mixed gets the items as either a list (for multiple)
	// or an object (for singles).
	Mixed Pop = "mixed"
)

// Turn the Pop into a string.
func (p Pop) String() string {
	return string(p)
}
