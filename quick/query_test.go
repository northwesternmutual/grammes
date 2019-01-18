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
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/northwesternmutual/grammes/logging"
	"github.com/northwesternmutual/grammes/manager"
	"github.com/northwesternmutual/grammes"
)

func TestExecuteQuery(t *testing.T) {
	defer func() {
		client = nil
		logger = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, map[string]string, map[string]string) ([]byte, error) {
		return nil, nil
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewBasicLogger(), execute)
	Convey("Given a host string a query", t, func() {
		host := "testhost"
		var q mockQuery
		Convey("When ExecuteQuery is called", func() {
			_, err := ExecuteQuery(host, q)
			Convey("Then no errors should be thrown", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
