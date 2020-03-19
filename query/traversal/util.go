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
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// fmtStr is used because it prevents from
// importing fmt in multiple files.
var fmtStr = fmt.Sprintf

// NewTraversal will return a new Query with
// a default value of 'g' to start a command.
func NewTraversal() (g String) {
	g.string = "g"
	g.buffer = bytes.NewBufferString("")
	return
}

// NewCustomTraversal could be used when you need to specifically
// need to change some property of the traversal.
// This can be something such as:
//  // ==> graph.traversal().withoutStrategies(LazyBarrierStrategy)
func NewCustomTraversal(str string) (g String) {
	g.string = str
	g.buffer = bytes.NewBufferString("")
	return g
}

func (g String) String() string {
	return g.string
}

// Raw will return the raw traversal commands
// to be used as a parameter for other steps.
func (g String) Raw() String {
	str := g.String()
	res := strings.TrimPrefix(str, "g.")
	cmd := NewCustomTraversal(res)
	return cmd
}

// append will simply take this Query
// and append a string to it
func (g String) append(str string) String {
	g.string += str
	return g
}

// AddStep will add a new step to the traversal string
// using a list of parameters.
func (g *String) AddStep(step string, params ...interface{}) {
	g.buffer.Reset()

	g.buffer.WriteString("." + step + "(")

	for i, p := range params {
		switch t := p.(type) {
		case String:
			g.buffer.WriteString(t.Raw().String())
		case Parameter:
			g.buffer.WriteString(t.String())
		case byte:
			g.buffer.WriteByte(t)
		case []byte:
			g.buffer.Write(t)
		case string:
			g.buffer.WriteString("\"" + strings.ReplaceAll(t, "\"", "\\\"") + "\"")
		default:
			g.buffer.WriteString(fmt.Sprintf("%v", t))
		}

		g.commaSeperator(i, params...)
	}

	g.buffer.WriteString(")")

	g.string += g.buffer.String()
}

func (g *String) commaSeperator(i int, params ...interface{}) {
	if len(params) > i+1 {
		if params[i+1] != nil {
			g.buffer.WriteString(",")
		}
	}
}

// gatherInts will act as a filter for
// pseudo optional parameters.
func gatherInts(params ...int) string {
	switch len(params) {
	case 1:
		return strconv.Itoa(params[0])
	default:
		return ""
	}
}
