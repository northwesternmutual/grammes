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

// NilLogger is the default logger used for the Grammes client.
// This logger will not print anything out.
type NilLogger struct{}

// NewNilLogger returns a nil logging
// object for the Grammes client to use.
func NewNilLogger() *NilLogger {
	return &NilLogger{}
}

// PrintQuery nothing.
func (*NilLogger) PrintQuery(q string) {}

// Error logs nothing.
func (*NilLogger) Error(msg string, err error) {}

// Fatal logs nothing.
func (*NilLogger) Fatal(msg string, err error) {}

// Debug logs nothing.
func (*NilLogger) Debug(msg string, params map[string]interface{}) {}
