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

package quick

import (
	"github.com/northwesternmutual/grammes/query"
)

// ExecuteQuery is used to execute a query
// to a gremlin server without having a client already
// created. For example you can use this if you are making
// quick small changes across multiple packages.
func ExecuteQuery(host string, query query.Query) ([]byte, error) {
	// Store the query into a string.
	strQuery := query.String()

	// Execute the query.
	res, err := ExecuteStringQuery(host, strQuery)

	return res, err
}

// ExecuteStringQuery is used to execute a query
// via a string to a gremlin server without have a client
// already created. For example this can be used if you are
// altering a graph through various packages.
func ExecuteStringQuery(host, query string) ([]byte, error) {
	res, err := executeQuery(host, query)

	return res, err
}
