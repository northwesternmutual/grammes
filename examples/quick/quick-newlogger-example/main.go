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

package main

import (
	"github.com/northwesternmutual/grammes/examples/exampleutil"
	"github.com/northwesternmutual/grammes/quick"

	"go.uber.org/zap"
)

var (
	// localhost contains the local address
	// needed to connect to a gremlin server.
	localhost = "ws://127.0.0.1:8182"
	// this is the designated logging system
	// being used to keep track of this program.
	logger *zap.Logger
)

// CustomLogger is our new logger
// to print using zap.
type CustomLogger struct {}

// PrintQuery will print the query out
// using the zap library rather than log.
func (*CustomLogger) PrintQuery(q string) {
	logger.Info("QUERY", zap.String("cmd", q))
}

// Debug
func (*CustomLogger) Debug(msg string, fields map[string]interface{}) {
	var arguments []zap.Field
	for k, v := range fields {
		arguments = append(arguments, zap.Any(k,v))
	}
	if len(arguments) > 0 {
		logger.Debug(msg, arguments...)
	}
}

// Error
func (*CustomLogger) Error(string, error) {}

// Fatal
func (*CustomLogger) Fatal(string, error) {}

func main() {
	// Setup the logger using zap.
	logger = exampleutil.SetupLogger()
	
	defer logger.Sync()

	quick.SetLogger(&CustomLogger{})

	quick.DropAll(localhost)
}
