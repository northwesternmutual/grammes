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

package predicate

import (
	"bytes"
	"fmt"
)

// Equal checks if this value is
// exactly equal to the querying value.
func Equal(val interface{}) *Predicate {
	s := fmt.Sprintf("eq(%#v)", val)
	a := Predicate(s)
	return &a
}

// NotEqual check if this value is
// NOT equal to the query value.
func NotEqual(val interface{}) *Predicate {
	s := fmt.Sprintf("neq(%#v)", val)
	a := Predicate(s)
	return &a
}

// LessThan checks if this value is
// less than the querying value.
func LessThan(val interface{}) *Predicate {
	s := fmt.Sprintf("lt(%#v)", val)
	a := Predicate(s)
	return &a
}

// LessThanOrEqual checks if this value is
// less than or equal to the querying value.
func LessThanOrEqual(val interface{}) *Predicate {
	s := fmt.Sprintf("lte(%#v)", val)
	a := Predicate(s)
	return &a
}

// GreaterThan checks if this value is
// greater than the querying value.
func GreaterThan(val interface{}) *Predicate {
	s := fmt.Sprintf("gt(%#v)", val)
	a := Predicate(s)
	return &a
}

// GreaterThanOrEqual checks if this value is
// greater than or equal to the querying value.
func GreaterThanOrEqual(val interface{}) *Predicate {
	s := fmt.Sprintf("gte(%#v)", val)
	a := Predicate(s)
	return &a
}

// Inside checks if this value is
// within the minimum and maximum querying values.
func Inside(min, max interface{}) *Predicate {
	s := fmt.Sprintf("inside(%#v, %#v)", min, max)
	a := Predicate(s)
	return &a
}

// Within checks if this value is within the array values.
func Within(params ...interface{}) *Predicate {
	buffer := bytes.NewBufferString("within(")
	sep := ""
	for _, p := range params {
		buffer.WriteString(sep)
		switch t := p.(type) {
		case string:
			buffer.WriteString("\"" + t + "\"")
		default:
			buffer.WriteString(fmt.Sprintf("%v", t))
		}

		sep = ","
	}

	buffer.WriteString(")")
	a := Predicate(buffer.String())
	return &a
}
