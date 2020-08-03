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
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/northwesternmutual/grammes"
	"github.com/northwesternmutual/grammes/gremconnect"
	"github.com/northwesternmutual/grammes/logging"
	"github.com/northwesternmutual/grammes/manager"
)

var (
	connect        error
	isConnected    = true
	isDisposed     = false
	vertexResponse = `
	[
		{
			"@type": "g:Vertex",
			"@value": {
				"id": {
					"@type": "g:Int64",
					"@value": 28720
				},
				"label": "newvertex"
			}
		}
	]
	`
	idResponse = `
	[
		{
			"@type": "g:Id",
			"@value": 255
		}
	]
	`
)

// MOCKDIALER

type mockDialer gremconnect.WebSocket

func (*mockDialer) Connect() error                    { return connect }
func (*mockDialer) Close() error                      { return nil }
func (*mockDialer) Write([]byte) error                { return nil }
func (*mockDialer) Read() ([]byte, error)             { return nil, nil }
func (*mockDialer) Ping(chan error)                   {}
func (*mockDialer) IsConnected() bool                 { return isConnected }
func (*mockDialer) IsDisposed() bool                  { return isDisposed }
func (*mockDialer) Auth() (*gremconnect.Auth, error)  { return &gremconnect.Auth{}, nil }
func (*mockDialer) Address() string                   { return "" }
func (*mockDialer) GetQuit() chan struct{}            { return make(chan struct{}) }
func (*mockDialer) SetAuth(string, string)            {}
func (*mockDialer) SetTimeout(time.Duration)          {}
func (*mockDialer) SetPingInterval(time.Duration)     {}
func (*mockDialer) SetWritingWait(time.Duration)      {}
func (*mockDialer) SetReadingWait(time.Duration)      {}
func (*mockDialer) SetWriteBufferSize(int)            {}
func (*mockDialer) SetReadBufferSize(int)             {}
func (*mockDialer) SetHandshakeTimeout(time.Duration) {}
func (*mockDialer) SetCompression(bool)               {}

// MOCKQUERY

type mockQuery string

func (mockQuery) String() string { return "TEST" }

func TestAddAPIVertex(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, map[string]string, map[string]string) ([][]byte, error) {
		return [][]byte{[]byte(vertexResponse)}, nil
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string and APIData object", t, func() {
		host := "testhost"
		var data grammes.APIData
		Convey("When AddAPIVertex is called", func() {
			_, err := AddAPIVertex(host, data)
			Convey("Then no errors should be thrown", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestAddAPIVertexClientError(t *testing.T) {
	tempcheckForClient := checkForClient
	defer func() {
		checkForClient = tempcheckForClient
	}()
	checkForClient = func(string) error { return errors.New("ERROR") }
	Convey("Given a host string and APIData object", t, func() {
		host := "testhost"
		var data grammes.APIData
		Convey("When AddAPIVertex is called and there is an error checking for the client", func() {
			_, err := AddAPIVertex(host, data)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddAPIVertexQueryError(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, map[string]string, map[string]string) ([][]byte, error) {
		return nil, errors.New("ERROR")
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string and APIData object", t, func() {
		host := "testhost"
		var data grammes.APIData
		Convey("When AddAPIVertex is called and there is an error while querying", func() {
			_, err := AddAPIVertex(host, data)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddVertexByStruct(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, map[string]string, map[string]string) ([][]byte, error) {
		return [][]byte{[]byte(vertexResponse)}, nil
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string and Vertex object", t, func() {
		host := "testhost"
		var vertex grammes.Vertex
		Convey("When AddVertexByStruct is called", func() {
			_, err := AddVertexByStruct(host, vertex)
			Convey("Then no errors should be thrown", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestAddVertexByStructClientError(t *testing.T) {
	tempcheckForClient := checkForClient
	defer func() {
		checkForClient = tempcheckForClient
	}()
	checkForClient = func(string) error { return errors.New("ERROR") }
	Convey("Given a host string and Vertex object", t, func() {
		host := "testhost"
		var vertex grammes.Vertex
		Convey("When AddVertexByStruct is called and there is an error checking for the client", func() {
			_, err := AddVertexByStruct(host, vertex)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddVertexByStructQueryError(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, map[string]string, map[string]string) ([][]byte, error) {
		return nil, errors.New("ERROR")
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string and Vertex object", t, func() {
		host := "testhost"
		var vertex grammes.Vertex
		Convey("When AddVertexByStruct is called and there is an error while querying", func() {
			_, err := AddVertexByStruct(host, vertex)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddVertex(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, map[string]string, map[string]string) ([][]byte, error) {
		return [][]byte{[]byte(vertexResponse)}, nil
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string and label string", t, func() {
		host := "testhost"
		label := "testlabel"
		Convey("When AddVertex is called", func() {
			_, err := AddVertex(host, label)
			Convey("Then no errors should be thrown", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestAddVertexClientError(t *testing.T) {
	tempcheckForClient := checkForClient
	defer func() {
		checkForClient = tempcheckForClient
	}()
	checkForClient = func(string) error { return errors.New("ERROR") }
	Convey("Given a host string and labelv", t, func() {
		host := "testhost"
		label := "testlabel"
		Convey("When AddVertex is called and there is an error checking for the client", func() {
			_, err := AddVertex(host, label)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddVertexQueryError(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, map[string]string, map[string]string) ([][]byte, error) {
		return nil, errors.New("ERROR")
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string and label string", t, func() {
		host := "testhost"
		label := "testlabel"
		Convey("When AddVertex is called and there is an error while querying", func() {
			_, err := AddVertex(host, label)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddVertexLabels(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, map[string]string, map[string]string) ([][]byte, error) {
		return [][]byte{[]byte(vertexResponse)}, nil
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string and label string", t, func() {
		host := "testhost"
		label := "testlabel"
		Convey("When AddVertexLabelsis called", func() {
			_, err := AddVertexLabels(host, label)
			Convey("Then no errors should be thrown", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestAddVertexLabelsClientError(t *testing.T) {
	tempcheckForClient := checkForClient
	defer func() {
		checkForClient = tempcheckForClient
	}()
	checkForClient = func(string) error { return errors.New("ERROR") }
	Convey("Given a host string", t, func() {
		host := "testhost"
		Convey("When AddVertexLabels is called and there is an error checking for the client", func() {
			_, err := AddVertexLabels(host)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddVertexLabelsQueryError(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, map[string]string, map[string]string) ([][]byte, error) {
		return nil, errors.New("ERROR")
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string and label string", t, func() {
		host := "testhost"
		label := "testlabel"
		Convey("When AddVertexLabels is called and there is an error while querying", func() {
			_, err := AddVertexLabels(host, label)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddVertexByQuery(t *testing.T) {
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
		Convey("When AddVertexByQuery is called", func() {
			_, err := AddVertexByQuery(host, q)
			Convey("Then no errors should be thrown", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestAddVertexByQueryClientError(t *testing.T) {
	tempcheckForClient := checkForClient
	defer func() {
		checkForClient = tempcheckForClient
	}()
	checkForClient = func(string) error { return errors.New("ERROR") }
	Convey("Given a host string and query", t, func() {
		host := "testhost"
		var q mockQuery
		Convey("When AddVertexByQuery is called and there is an error checking for the client", func() {
			_, err := AddVertexByQuery(host, q)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddVertexByQueryQueryError(t *testing.T) {
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
		Convey("When AddVertexByQuery is called and there is an error while querying", func() {
			_, err := AddVertexByQuery(host, q)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
