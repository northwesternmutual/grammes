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

// String search predicates which match
// against the entire string value.

// TextPrefix finds if the string value starts
// with the given string.
func TextPrefix(str string) *Predicate {
	s := "textPrefix(" + str + ")"
	a := Predicate(s)
	return &a
}

// TextRegex finds if the string value matches
// the given regular expression in its entirety.
func TextRegex(str string) *Predicate {
	s := "textRegex(" + str + ")"
	a := Predicate(s)
	return &a
}

// TextFuzzy finds if the string value is
// similar to the given query string.
func TextFuzzy(str string) *Predicate {
	s := "textFuzzy(" + str + ")"
	a := Predicate(s)
	return &a
}
