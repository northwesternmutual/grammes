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
)

// http://tinkerpop.apache.org/docs/current/reference/#addproperty-step

// Property (sideEffect) unlike AddV() and AddE(), Property() is
// a full sideEffect step in that it does not return the property
// that it created, but the element that streamed into it.
// Note:
// - This function does not handle your string formatting
//   because of the Cardinality parameter.
// Signatures:
// Property(interface{} (Object), interface{} (Object), ...interface{} (Object))
// Property(Cardinality, string (Object), interface{} (Object), ...interface{} (Object))
func (g String) Property(objOrCard interface{}, obj interface{}, params ...interface{}) String {
	fullParams := make([]interface{}, 0, len(params)+2)
	fullParams = append(fullParams, objOrCard, obj)
	fullParams = append(fullParams, params...)

	for i, p := range fullParams {
		switch t := p.(type) {
		// Write every numeric property as long - to prevent indexing issues
		case int, int16, int32, int64, uint, uint16, uint32, uint64:
			// Convert to byte, so that "" won't be added
			fullParams[i] = []byte(fmt.Sprintf("%vl", t))
		}
	}

	g.AddStep("property", fullParams...)

	return g
}
