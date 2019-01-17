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
	"errors"
	"testing"

	"github.com/northwesternmutual/grammes/logging"
	"github.com/northwesternmutual/grammes/manager"

	grammes "github.com/northwesternmutual/grammes"

	. "github.com/smartystreets/goconvey/convey"
)

func TestVertexIDsByQuery(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, map[string]string, map[string]string) ([]byte, error) {
		return []byte(idResponse), nil
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewBasicLogger(), execute)
	Convey("Given a host string and query", t, func() {
		host := "testhost"
		var q mockQuery
		Convey("When VertexIDsByQuery is called", func() {
			_, err := VertexIDsByQuery(host, q)
			Convey("Then no errors should be thrown", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestVertexIDsByQueryClientError(t *testing.T) {
	tempCheckForClient := CheckForClient
	defer func() {
		CheckForClient = tempCheckForClient
	}()
	CheckForClient = func(string) error { return errors.New("ERROR") }
	Convey("Given a host string and query", t, func() {
		host := "testhost"
		var q mockQuery
		Convey("When VertexIDsByQuery is called and encounters an error checking for the client", func() {
			_, err := VertexIDsByQuery(host, q)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestVertexIDsByQueryQueryError(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, map[string]string, map[string]string) ([]byte, error) {
		return nil, errors.New("ERROR")
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewBasicLogger(), execute)
	Convey("Given a host string and query", t, func() {
		host := "testhost"
		var q mockQuery
		Convey("When VertexIDsByQuery is called and there is querying error", func() {
			_, err := VertexIDsByQuery(host, q)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestVertexIDs(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, map[string]string, map[string]string) ([]byte, error) {
		return []byte(idResponse), nil
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewBasicLogger(), execute)
	Convey("Given a host string, label and properties", t, func() {
		host := "testhost"
		label := "testlabel"
		Convey("When VertexIDs is called", func() {
			_, err := VertexIDs(host, label, "prop1", "prop2")
			Convey("Then no errors should be thrown", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestVertexIDsClientError(t *testing.T) {
	tempCheckForClient := CheckForClient
	defer func() {
		CheckForClient = tempCheckForClient
	}()
	CheckForClient = func(string) error { return errors.New("ERROR") }
	Convey("Given a host string, label and properties", t, func() {
		host := "testhost"
		label := "testlabel"
		Convey("When VertexIDs is called and encounters an error checking for the client", func() {
			_, err := VertexIDs(host, label, "prop1", "prop2")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

type testLogger struct{}

func (testLogger) PrintQuery(string)                    {}
func (testLogger) Debug(string, map[string]interface{}) {}
func (testLogger) Error(string, error)                  {}
func (testLogger) Fatal(string, error)                  {}

func TestVertexIDsQueryError(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer, grammes.WithLogger(&testLogger{}))
	execute := func(string, map[string]string, map[string]string) ([]byte, error) {
		return nil, errors.New("ERROR")
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewBasicLogger(), execute)
	Convey("Given a host string, label and properties", t, func() {
		host := "testhost"
		label := "testlabel"
		Convey("When VertexIDs is called and there is querying error", func() {
			_, err := VertexIDs(host, label, "prop1", "prop2")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
