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

package logging

// Logger is a wrapper for any kind of logger
// you wish to use. This can be customized and
// changed within the Grammes client itself.
type Logger interface {
	// This function specifically is meant
	// to log queries. This is in case you are
	// debugging an application and you wish to
	// print out the queries to Stdout or log
	// them somewhere else you may without
	// getting extra logs that you don't want.
	PrintQuery(msg string)
	// Debug is used when confirming when things
	// are doing their jobs such as when adding
	// vertex labels to the schema.
	Debug(msg string, fieldAndVals map[string]interface{})
	// Error is used when there is a problem but
	// not a big enough problem to stop an app.
	// These problems are minor, but not major.
	Error(msg string, err error)
	// Fatal's purpose is to stop the application
	// because something really wrong happened.
	// A case of this being used is when trying to
	// put an odd number of properties in an AddVertex
	// function. Which would not create a proper query
	// for the gremlin server and should stop.
	Fatal(msg string, err error)
}
