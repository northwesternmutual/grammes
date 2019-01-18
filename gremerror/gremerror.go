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
Package gremerror holds custom error types to return in the Grammes package.

All exported functions in this package returns an object that
implements the error interface.
*/
package gremerror

import "errors"

// fmtError will take the key and value then
// return a string that is formatted like a JSON
// that can then be used with the fmtComma function.
func fmtError(k, v string) string { return "{\"" + k + "\":\"" + v + "\"}" }

// fmtComma is the data seperator function.
// used to take in Data from the fmtError function
// and format all of it into a concise list with commas.
func fmtComma(data ...string) string {
	var res string
	for i, d := range data {
		res += d
		if len(data) == i+1 {
			break
		}
		res += ","
	}
	return res
}

var (
	// ErrOddNumberOfParameters is used when there
	// are an odd number of parameters fed into a
	// function that should only have paired parameters.
	ErrOddNumberOfParameters = errors.New("odd number of parameters")
	// ErrDisposedConnection is used when the connection from the
	// dialer is disposed and not reachable.
	ErrDisposedConnection = errors.New("connection is disposed")
	// ErrNilClient is used for functions that have a queryClient as
	// a parameter, and the client is nil.
	ErrNilClient = errors.New("nil client given to function")
	// ErrEmptyResponse is used when the response given back from
	// querying was empty. This is used on rare occasions when
	// the Unmarshal process is successful, but returns something empty.
	ErrEmptyResponse = errors.New("empty response received")
)

// GrammesError is a generic error
// used when something nonspecific is
// wrong inside a Grammes function.
type GrammesError struct {
	function string
	err      error
}

// NewGrammesError returns a new generic error
// for when something goes wrong in a Grammes package function.
func NewGrammesError(function string, err error) error {
	return &GrammesError{
		function: function,
		err:      err,
	}
}

func (g *GrammesError) Error() string {
	return fmtComma(
		fmtError("type", "GRAMMES_ERROR"),
		fmtError("function", g.function),
		fmtError("error", g.err.Error()),
	)
}
