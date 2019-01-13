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
	"go/token"

	"github.com/northwesternmutual/grammes/query/predicate"
)

// http://tinkerpop.apache.org/docs/current/reference/#has-step

// Has (filter) filters vertices, edges, and vertex properties
// based on their properties.
// Signatures:
// Has(string)
// Has(string, string (Object))
// Has(string, string (P))
// Has(string, string, string (Object))
// Has(string, string, string (P))
// Has(string, *String (Traversal))
// Has(Token, string (Object))
// Has(Token, string (P))
// Has(Token, *String (Traversal))
func (g String) Has(first interface{}, params ...interface{}) String {
	g = g.append(".has(")

	switch t := first.(type) {
	case token.Token:
		g = g.append(t.String())
	case string:
		g = g.append("\"" + t + "\"")
	default:
		g = g.append(fmtStr("%v", first))
	}

	if len(params) > 2 {
		fmt.Println("ERROR: Too many parameters to call Has()")
	}

	if len(params) > 0 {
		for _, v := range params {
			switch t := v.(type) {
			case string:
				g = g.append(",\"" + t + "\"")
			case String:
				g = g.append("," + t.Raw().String())
			case *predicate.Predicate:
				g = g.append("," + t.String() + "")
			default:
				g = g.append(fmtStr(",%v", v))
			}
		}
	}

	g = g.append(")")

	return g
}

// HasID (filter) filters vertices, edges, and vertex properties
// based on their properties. In this case the ID.
// Signatures:
// HasID(string (Object), ...string (Object))
// HasID(string (P))
func (g String) HasID(objOrP interface{}, objs ...string) String {
	// g = g.append(fmtStr(".hasId(%v", objOrP))

	// if len(objs) > 0 {
	// 	for _, v := range objs {
	// 		g = g.append(",\"" + v + "\"")
	// 	}
	// }

	// g = g.append(")")

	var p []interface{}

	p = append(p, objOrP)

	for _, s := range objs {
		p = append(p, s)
	}

	g.AddStep("hasId", p...)

	return g
}

// HasKey (filter) filters vertices, edges, and vertex properties
// based on their properties. In this case the Key.
// Signatures:
// HasKey(string (Predicate))
// HasKey(string, ...string)
func (g String) HasKey(pOrStr interface{}, handledStrings ...string) String {
	switch pOrStr.(type) {
	case string:
		g = g.append(".hasKey(\"" + pOrStr.(string) + "\"")
	default:
		g = g.append(fmtStr(".hasKey(%v", pOrStr))
	}

	if len(handledStrings) > 0 {
		for _, v := range handledStrings {
			g = g.append(",\"" + v + "\"")
		}
	}

	g = g.append(")")

	return g
}

// HasLabel (filter) filters vertices, edges, and vertex properties
// based on their properties. In this case the Label.
// Signatures:
// HasLabel(string (Predicate))
// HasLabel(string, ...string)
func (g String) HasLabel(pOrStr interface{}, handledStrings ...string) String {
	switch pOrStr.(type) {
	case string:
		g = g.append(".hasLabel(\"" + pOrStr.(string) + "\"")
	default:
		g = g.append(fmtStr(".hasLabel(%v", pOrStr))
	}

	if len(handledStrings) > 0 {
		for _, v := range handledStrings {
			g = g.append(",\"" + v + "\"")
		}
	}

	g = g.append(")")

	return g
}

// HasNot (filter) filters vertices, edges, and vertex properties
// based on their properties. In this case if it doesn't have this property.
// Signatures:
// HasNot(string)
func (g String) HasNot(str string) String {
	g = g.append(".hasNot(\"" + str + "\")")

	return g
}

// HasValue (filter) filters vertices, edges, and vertex properties
// based on their properties. In this case the Value.
// Signatures:
// HasValue(string (Object), ...string (Object))
// HasValue(string (P))
func (g String) HasValue(objOrP interface{}, objs ...string) String {
	switch objOrP.(type) {
	case string:
		g = g.append(".hasValue(\"" + objOrP.(string) + "\"")
	default:
		g = g.append(fmtStr(".hasValue(%v)", objOrP))
	}

	if len(objs) > 0 {
		for _, v := range objs {
			g = g.append(",\"" + v + "\"")
		}
	}

	g = g.append(")")

	return g
}
