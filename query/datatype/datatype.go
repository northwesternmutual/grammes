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
Package datatype contains the object to represent Datatypes in Gremlin.

DataType and its constant values are the string equivalent of
the Gremlin language's datatypes when referencing them in a traversal/query.

A note about DataType:

This object implements the Parameter interfaces used by graph traversals.
*/
package datatype

// DataType represents all the data types that
// are inside of Gremlin.
type DataType string

const (
	// String represents the class for String
	String DataType = "String.class"
	// Character represents the class for Character
	Character DataType = "Character.class"
	// Boolean represents the class for Boolean
	Boolean DataType = "Boolean.class"
	// Byte represents the class for Byte
	Byte DataType = "Byte.class"
	// Short represents the class for Short
	Short DataType = "Short.class"
	// Integer represents the class for Integer
	Integer DataType = "Integer.class"
	// Long represents the class for Long
	Long DataType = "Long.class"
	// Float represents the class for Float
	Float DataType = "Float.class"
	// Double represents the class for Double
	Double DataType = "Double.class"
	// Decimal represents the class for Decimal
	Decimal DataType = "Decimal.class"
	// Precision represents the class for Precision
	Precision DataType = "Precision.class"
	// Geoshape represents the class for Geoshape
	Geoshape DataType = "Geoshape.class"
)

func (d DataType) String() string {
	return string(d)
}
