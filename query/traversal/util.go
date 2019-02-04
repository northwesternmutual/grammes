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
	"strconv"
	"strings"
)

// fmtStr is used because it prevents from
// importing fmt in multiple files.
var fmtStr = fmt.Sprintf

// NewTraversal will return a new Query with
// a default value of 'g' to start a command.
func NewTraversal() (g String) {
	g = "g"
	return
}

func (g String) String() (query string) {
	query = string(g)
	return
}

// Raw will return the raw traversal commands
// to be used as a parameter for other steps.
func (g String) Raw() String {
	str := g.String()
	res := strings.TrimPrefix(str, "g.")
	cmd := String(res)
	return cmd
}

// append will simply take this Query
// and append a string to it
func (g String) append(str string) (trav String) {
	trav = g + String(str)
	return
}

// AddStep will add a new step to the traversal string
// using a list of parameters.
func (g *String) AddStep(step string, params ...interface{}) {
	buffer.Reset()

	buffer.WriteString("." + step + "(")

	for i, p := range params {
		switch t := p.(type) {
		case String:
			buffer.WriteString(t.Raw().String())
		case Parameter:
			buffer.WriteString(t.String())
		case byte:
			buffer.WriteByte(t)
		case []byte:
			buffer.Write(t)
		case string:
			buffer.WriteString("\"" + t + "\"")
		default:
			buffer.WriteString(fmt.Sprintf("%v", t))
		}

		g.commaSeperator(i, params...)
	}

	buffer.WriteString(")")

	*g = String(g.String() + buffer.String())
}

func (g *String) commaSeperator(i int, params ...interface{}) {
	if len(params) > i+1 {
		if params[i+1] != nil && params[i+1] != "" {
			buffer.WriteString(",")
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
