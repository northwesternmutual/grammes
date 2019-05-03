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
Package operator contains the object to apply mathematical operations to a graph traversal.

Contains the name for mathematical operators in the Gremlin traversal language.

A note about Operator:

This object implements the Parameter interfaces used by graph traversals.
*/
package operator

// Operator unlike Numberhelper doesn't have functions
// associated to applying the mathematical operator.
// Instead this enumeration only uses the keywords and strings.
type Operator string

const (
	// AddAll will add all of the number provided in the traversal.
	AddAll Operator = "addAll"
	// And will tell the traversal to wait for another parameter.
	And Operator = "and"
	// Assign will assign a number to an alias.
	Assign Operator = "assign"
	// Div takes the amount of times one number fits into another.
	Div Operator = "div"
	// Max finds the largest number found in the traversal.
	Max Operator = "max"
	// Min finds the smallest number in the traversal.
	Min Operator = "min"
	// Minus subtracts from two or more numbers.
	Minus Operator = "minus"
	// Mult will multiply the numbers found in the traversal.
	Mult Operator = "mult"
	// Or will make the traversal choose best fit number.
	Or Operator = "or"
	// Sum will take the sum of all numbers in the traversal.
	Sum Operator = "sum"
	// SumLong will take the sum of all numbers in the traversal as a long.
	SumLong Operator = "sumLong"
)

func (o Operator) String() string {
	return string(o)
}
