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

	. "github.com/smartystreets/goconvey/convey"

	"github.com/northwesternmutual/grammes"
	"github.com/northwesternmutual/grammes/logging"
	"github.com/northwesternmutual/grammes/manager"
)

func TestVerticesByQuery(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, map[string]string, map[string]string) ([][]byte, error) {
		return [][]byte{[]byte(vertexResponse)}, nil
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string and query", t, func() {
		host := "testhost"
		var q mockQuery
		Convey("When VerticesByQuery is called", func() {
			_, err := VerticesByQuery(host, q)
			Convey("Then no errors should be thrown", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestVerticesByQueryClientError(t *testing.T) {
	tempcheckForClient := checkForClient
	defer func() {
		checkForClient = tempcheckForClient
	}()
	checkForClient = func(string) error { return errors.New("ERROR") }
	Convey("Given a host string and query", t, func() {
		host := "testhost"
		var q mockQuery
		Convey("When VerticesByQuery is called and encounters an error checking for the client", func() {
			_, err := VerticesByQuery(host, q)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestVerticesByQueryQueryError(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, map[string]string, map[string]string) ([][]byte, error) {
		return nil, errors.New("ERROR")
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string and query", t, func() {
		host := "testhost"
		var q mockQuery
		Convey("When VerticesByQuery is called and there is querying error", func() {
			_, err := VerticesByQuery(host, q)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAllVertices(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, map[string]string, map[string]string) ([][]byte, error) {
		return [][]byte{[]byte(vertexResponse)}, nil
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string", t, func() {
		host := "testhost"
		Convey("When AllVertices is called", func() {
			_, err := AllVertices(host)
			Convey("Then no errors should be thrown", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestAllVerticesClientError(t *testing.T) {
	tempcheckForClient := checkForClient
	defer func() {
		checkForClient = tempcheckForClient
	}()
	checkForClient = func(string) error { return errors.New("ERROR") }
	Convey("Given a host string", t, func() {
		host := "testhost"
		Convey("When AllVertices is called and encounters an error checking for the client", func() {
			_, err := AllVertices(host)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAllVerticesQueryError(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, map[string]string, map[string]string) ([][]byte, error) {
		return nil, errors.New("ERROR")
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string", t, func() {
		host := "testhost"
		Convey("When AllVertices is called and there is querying error", func() {
			_, err := AllVertices(host)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestVertexByID(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, map[string]string, map[string]string) ([][]byte, error) {
		return [][]byte{[]byte(vertexResponse)}, nil
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string and ID int", t, func() {
		host := "testhost"
		id := int64(123)
		Convey("When VertexByID is called", func() {
			_, err := VertexByID(host, id)
			Convey("Then no errors should be thrown", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestVertexByIDClientError(t *testing.T) {
	tempcheckForClient := checkForClient
	defer func() {
		checkForClient = tempcheckForClient
	}()
	checkForClient = func(string) error { return errors.New("ERROR") }
	Convey("Given a host string and ID int", t, func() {
		host := "testhost"
		id := int64(123)
		Convey("When VertexByID is called and encounters an error checking for the client", func() {
			_, err := VertexByID(host, id)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestVertexByIDQueryError(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, map[string]string, map[string]string) ([][]byte, error) {
		return nil, errors.New("ERROR")
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string and ID int", t, func() {
		host := "testhost"
		id := int64(123)
		Convey("When VertexByID is called and there is querying error", func() {
			_, err := VertexByID(host, id)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestVertices(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, map[string]string, map[string]string) ([][]byte, error) {
		return [][]byte{[]byte(vertexResponse)}, nil
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string, label and properties", t, func() {
		host := "testhost"
		label := "testlabel"
		Convey("When Vertices is called", func() {
			_, err := Vertices(host, label, "prop1", "prop2")
			Convey("Then no errors should be thrown", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestVerticesClientError(t *testing.T) {
	tempcheckForClient := checkForClient
	defer func() {
		checkForClient = tempcheckForClient
	}()
	checkForClient = func(string) error { return errors.New("ERROR") }
	Convey("Given a host string, label and properties", t, func() {
		host := "testhost"
		label := "testlabel"
		Convey("When Vertices is called and encounters an error checking for the client", func() {
			_, err := Vertices(host, label, "prop1", "prop2")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestVerticesQueryError(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, map[string]string, map[string]string) ([][]byte, error) {
		return nil, errors.New("ERROR")
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string, label and properties", t, func() {
		host := "testhost"
		label := "testlabel"
		Convey("When Vertices is called and there is querying error", func() {
			_, err := Vertices(host, label, "prop1", "prop2")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
