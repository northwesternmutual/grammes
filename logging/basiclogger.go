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

// BasicLogger is the default logger
// used by the Grammes client. This
// particular logger uses zap by default.
// The reasoning behind using a wrapper for
// logging is for flexibility when using
// the Grammes package. You are given
// the freedom to choose any kind of
// logger you wish as long as you create
// a wrapper that meets the criteria set
// by the Logger interface.
type BasicLogger struct {
	zapper *zap.Logger
}

// NewBasicLogger returns a logger that is used for
// development, but only logs at the Error level.
func NewBasicLogger() *BasicLogger {
	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	config.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	zap, _ := config.Build()
	return &BasicLogger{zapper: zap}
}

// PrintQuery will print the query at DebugLevel by default
func (logger *BasicLogger) PrintQuery(q string) {
	logger.zapper.Debug("[Grammes]", zap.String("query", q))
}

// Error logs at ErrorLevel
func (logger *BasicLogger) Error(msg string, err error) {
	logger.zapper.Error(msg, zap.Error(err))
}

// Fatal logs at FatalLevel
func (logger *BasicLogger) Fatal(msg string, err error) {
	logger.zapper.Fatal(msg, zap.Error(err))
}

// Debug logs at DebugLevel
func (logger *BasicLogger) Debug(msg string, params map[string]interface{}) {
	var fields []zap.Field
	for k, v := range params {
		fields = append(fields, zap.Any(k, v))
	}
	logger.zapper.Debug(msg, fields...)
}
