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

// Text search predicates which match against the individual words inside a text
// string after it has been tokenized. These predicates are not case sensitive.

// TextContains finds if at least one word inside
// the text string matches the query string.
func TextContains(str string) *Predicate {
	s := "textContains(" + str + ")"
	a := Predicate(s)
	return &a
}

// TextContainsPrefix finds if one word inside
// the text string begins with the query string.
func TextContainsPrefix(str string) *Predicate {
	s := "textContains(" + str + ")"
	a := Predicate(s)
	return &a
}

// TextContainsRegex finds if one word inside
// the text string matches the given regular expression.
func TextContainsRegex(str string) *Predicate {
	s := "textContains(" + str + ")"
	a := Predicate(s)
	return &a
}

// TextContainsFuzzy finds if one word inside
// the text string is similar to the query string.
func TextContainsFuzzy(str string) *Predicate {
	s := "textContains(" + str + ")"
	a := Predicate(s)
	return &a
}
