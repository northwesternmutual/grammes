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

func TestDropVertexLabel(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, *time.Duration, map[string]string, map[string]string, *uuid.UUID) ([][]byte, error) {
		return [][]byte{[]byte(vertexResponse)}, nil
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string and label string", t, func() {
		host := "testhost"
		label := "testlabel"
		Convey("When DropVertexLabel is called", func() {
			err := DropVertexLabel(host, label)
			Convey("Then the error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestDropVertexLabelClientError(t *testing.T) {
	tempcheckForClient := checkForClient
	defer func() {
		checkForClient = tempcheckForClient
	}()
	checkForClient = func(string) error { return errors.New("ERROR") }
	Convey("Given a host string and label string", t, func() {
		host := "testhost"
		label := "testlabel"
		Convey("When DropVertexLabel is called and encounters an error checking for the client", func() {
			err := DropVertexLabel(host, label)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestDropVertexLabelQueryError(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, *time.Duration, map[string]string, map[string]string, *uuid.UUID) ([][]byte, error) {
		return nil, errors.New("ERROR")
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string and label string", t, func() {
		host := "testhost"
		label := "testlabel"
		Convey("When DropVertexLabel is called and there is querying error", func() {
			err := DropVertexLabel(host, label)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestDropVertexByID(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, *time.Duration, map[string]string, map[string]string, *uuid.UUID) ([][]byte, error) {
		return [][]byte{[]byte(vertexResponse)}, nil
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string and ID int", t, func() {
		host := "testhost"
		id := int64(123)
		Convey("When DropVertexByID is called", func() {
			err := DropVertexByID(host, id)
			Convey("Then the error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestDropVertexByIDClientError(t *testing.T) {
	tempcheckForClient := checkForClient
	defer func() {
		checkForClient = tempcheckForClient
	}()
	checkForClient = func(string) error { return errors.New("ERROR") }
	Convey("Given a host string and ID int", t, func() {
		host := "testhost"
		id := int64(123)
		Convey("When DropVertexByID is called and encounters an error checking for the client", func() {
			err := DropVertexByID(host, id)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestDropVertexByIDQueryError(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, *time.Duration, map[string]string, map[string]string, *uuid.UUID) ([][]byte, error) {
		return nil, errors.New("ERROR")
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string and ID int", t, func() {
		host := "testhost"
		id := int64(123)
		Convey("When DropVertexByID is called and there is querying error", func() {
			err := DropVertexByID(host, id)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestDropVerticesByQuery(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, *time.Duration, map[string]string, map[string]string, *uuid.UUID) ([][]byte, error) {
		return [][]byte{[]byte(vertexResponse)}, nil
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string and query", t, func() {
		host := "testhost"
		var q mockQuery
		Convey("When DropVerticesByQuery is called", func() {
			err := DropVerticesByQuery(host, q)
			Convey("Then the error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestDropVerticesByQueryClientError(t *testing.T) {
	tempcheckForClient := checkForClient
	defer func() {
		checkForClient = tempcheckForClient
	}()
	checkForClient = func(string) error { return errors.New("ERROR") }
	Convey("Given a host string and query", t, func() {
		host := "testhost"
		var q mockQuery
		Convey("When DropVerticesByQuery is called and encounters an error checking for the client", func() {
			err := DropVerticesByQuery(host, q)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestDropVerticesByQueryQueryError(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, *time.Duration, map[string]string, map[string]string, *uuid.UUID) ([][]byte, error) {
		return nil, errors.New("ERROR")
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string and query", t, func() {
		host := "testhost"
		var q mockQuery
		Convey("When DropVerticesByQuery is called and there is querying error", func() {
			err := DropVerticesByQuery(host, q)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
