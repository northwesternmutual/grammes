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

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// DebugLogger is the basic logger used for debugging an application.
// This will log everything in the debug level.
type DebugLogger struct {
	zapper *zap.Logger
}

// NewDebugLogger returns a new debug logging
// object for the Grammes client to use.
func NewDebugLogger() *DebugLogger {
	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	config.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	zap, _ := config.Build()
	return &DebugLogger{zapper: zap}
}

// PrintQuery will print the query at DebugLevel by default
func (logger *DebugLogger) PrintQuery(q string) {
	logger.zapper.Debug("[Grammes]", zap.String("query", q))
}

// Error logs at ErrorLevel
func (logger *DebugLogger) Error(msg string, err error) {
	logger.zapper.Error(msg, zap.Error(err))
}

// Fatal logs at FatalLevel
func (logger *DebugLogger) Fatal(msg string, err error) {
	logger.zapper.Fatal(msg, zap.Error(err))
}

// Debug logs at DebugLevel
func (logger *DebugLogger) Debug(msg string, params map[string]interface{}) {
	var fields []zap.Field
	for k, v := range params {
		fields = append(fields, zap.Any(k, v))
	}
	logger.zapper.Debug(msg, fields...)
}
