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

	"github.com/northwesternmutual/grammes/manager"
	"github.com/northwesternmutual/grammes/logging"
	"github.com/northwesternmutual/grammes"
)

func TestUnexportedExecuteQuery(t *testing.T) {
	defer func() {
		client = nil
		logger = nil
	}()
	Convey("Given a host and query", t, func() {
		host := "testhost"
		query := "testquery"
		Convey("When executeQuery is called with no client or logger established", func() {
			_, err := executeQuery(host, query)
			Convey("Then an error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestUnexportedExecuteQueryWithLogger(t *testing.T) {
	defer func() {
		client = nil
		logger = nil
	}()
	logger = logging.NewBasicLogger()
	Convey("Given a host and query", t, func() {
		host := "testhost"
		query := "testquery"
		Convey("When executeQuery is called with no client", func() {
			_, err := executeQuery(host, query)
			Convey("Then an error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestUnexportedExecuteQueryWithClientAndLogger(t *testing.T) {
	defer func() {
		client = nil
		logger = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	logger = logging.NewBasicLogger()
	execute := func(string, map[string]string, map[string]string) ([]byte, error) {
		return nil, nil
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewBasicLogger(), execute)
	Convey("Given a host and query", t, func() {
		host := "testhost"
		query := "testquery"
		Convey("When executeQuery is called with a client and logger established", func() {
			_, err := executeQuery(host, query)
			Convey("Then no error should be returned", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestCheckForClientWithLogger(t *testing.T) {
	defer func() {
		client = nil
		logger = nil
	}()
	logger = logging.NewBasicLogger()
	Convey("Given a host", t, func() {
		host := "testhost"
		Convey("When checkForClient is called with no client", func() {
			err := checkForClient(host)
			Convey("Then an error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestCheckForClient(t *testing.T) {
	defer func() {
		client = nil
		logger = nil
	}()
	Convey("Given a host", t, func() {
		host := "testhost"
		Convey("When checkForClient is called with no logger or client established", func() {
			err := checkForClient(host)
			Convey("Then an error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestSetLoggerNoClient(t *testing.T) {
	defer func() {
		client = nil
		logger = nil
	}()
	Convey("Given a logger", t, func() {
		l := logging.NewBasicLogger()
		Convey("When SetLogger is called with no client established", func() {
			SetLogger(l)
			Convey("Then the logger var should be set", func() {
				So(logger, ShouldNotBeNil)
			})
		})
	})
}

func TestSetLoggerWithClient(t *testing.T) {
	defer func() {
		client = nil
		logger = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	Convey("Given a logger", t, func() {
		l := logging.NewBasicLogger()
		Convey("When SetLogger is called with a client established", func() {
			SetLogger(l)
			Convey("Then no errors should be encountered", func() {
			})
		})
	})
}
