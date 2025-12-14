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
	"github.com/google/uuid"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/northwesternmutual/grammes"
	"github.com/northwesternmutual/grammes/logging"
	"github.com/northwesternmutual/grammes/manager"
)

func TestDropAll(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, *time.Duration, map[string]string, map[string]string, *uuid.UUID) ([][]byte, error) {
		return [][]byte{[]byte(vertexResponse)}, nil
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string", t, func() {
		host := "testhost"
		Convey("When DropAll is called", func() {
			err := DropAll(host)
			Convey("Then no errors should be thrown", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestDropAllClientError(t *testing.T) {
	tempcheckForClient := checkForClient
	defer func() {
		checkForClient = tempcheckForClient
	}()
	checkForClient = func(string) error { return errors.New("ERROR") }
	Convey("Given a host string", t, func() {
		host := "testhost"
		Convey("When DropAll is called and encounters an error checking for the client", func() {
			err := DropAll(host)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestDropAllQueryError(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, *time.Duration, map[string]string, map[string]string, *uuid.UUID) ([][]byte, error) {
		return nil, errors.New("ERROR")
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string", t, func() {
		host := "testhost"
		Convey("When DropAll is called and there is querying error", func() {
			err := DropAll(host)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestSetVertexProperty(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, *time.Duration, map[string]string, map[string]string, *uuid.UUID) ([][]byte, error) {
		return [][]byte{[]byte(vertexResponse)}, nil
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string, ID int and properties", t, func() {
		host := "testhost"
		id := int64(123)
		Convey("When SetVertexProperty is called", func() {
			err := SetVertexProperty(host, id, "prop1", "prop2")
			Convey("Then no errors should be thrown", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestSetVertexPropertyClientError(t *testing.T) {
	tempcheckForClient := checkForClient
	defer func() {
		checkForClient = tempcheckForClient
	}()
	checkForClient = func(string) error { return errors.New("ERROR") }
	Convey("Given a host string, ID int and properties", t, func() {
		host := "testhost"
		id := int64(123)
		Convey("When SetVertexProperty is called and encounters an error checking for the client", func() {
			err := SetVertexProperty(host, id, "prop1", "prop2")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestSetVertexPropertyQueryError(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, *time.Duration, map[string]string, map[string]string, *uuid.UUID) ([][]byte, error) {
		return nil, errors.New("ERROR")
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string, ID int and properties", t, func() {
		host := "testhost"
		id := int64(123)
		Convey("When SetVertexProperty is called and there is querying error", func() {
			err := SetVertexProperty(host, id, "prop1", "prop2")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestVertexCount(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, *time.Duration, map[string]string, map[string]string, *uuid.UUID) ([][]byte, error) {
		return [][]byte{[]byte(idResponse)}, nil
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string", t, func() {
		host := "testhost"
		Convey("When VertexCount is called", func() {
			_, err := VertexCount(host)
			Convey("Then no errors should be thrown", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestVertexCountClientError(t *testing.T) {
	tempcheckForClient := checkForClient
	defer func() {
		checkForClient = tempcheckForClient
	}()
	checkForClient = func(string) error { return errors.New("ERROR") }
	Convey("Given a host string", t, func() {
		host := "testhost"
		Convey("When VertexCount is called and encounters an error checking for the client", func() {
			_, err := VertexCount(host)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestVertexCountQueryError(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, *time.Duration, map[string]string, map[string]string, *uuid.UUID) ([][]byte, error) {
		return nil, errors.New("ERROR")
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string", t, func() {
		host := "testhost"
		Convey("When VertexCount is called and there is querying error", func() {
			_, err := VertexCount(host)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
