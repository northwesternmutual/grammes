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

package manager

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/northwesternmutual/grammes/gremconnect"
	"github.com/northwesternmutual/grammes/logging"
)

func TestSetLogger(t *testing.T) {
	Convey("Given a dialer, string executor and graph query manager", t, func() {
		dialer := gremconnect.NewWebSocketDialer("testaddress")
		execute := func(string, map[string]string, map[string]string) ([]byte, error) { return nil, nil }
		gm := NewGraphManager(dialer, logging.NewBasicLogger(), execute)
		Convey("When SetLogger is called we should not encounter any errors", func() {
			gm.SetLogger(logging.NewBasicLogger())
		})
	})
}

func TestMiscQuerier(t *testing.T) {
	Convey("Given a dialer, string executor and graph query manager", t, func() {
		dialer := gremconnect.NewWebSocketDialer("testaddress")
		execute := func(string, map[string]string, map[string]string) ([]byte, error) { return nil, nil }
		gm := NewGraphManager(dialer, logging.NewBasicLogger(), execute)
		Convey("When MiscQuerier is called", func() {
			mq := gm.MiscQuerier()
			Convey("Then we should return the miscellaneous querier", func() {
				So(mq, ShouldNotBeNil)
			})
		})
	})
}

func TestAddVertexQuerier(t *testing.T) {
	Convey("Given a dialer, string executor and graph query manager", t, func() {
		dialer := gremconnect.NewWebSocketDialer("testaddress")
		execute := func(string, map[string]string, map[string]string) ([]byte, error) { return nil, nil }
		gm := NewGraphManager(dialer, logging.NewBasicLogger(), execute)
		Convey("When AddVertexQuerier is called", func() {
			avq := gm.AddVertexQuerier()
			Convey("Then we should return the addVertex querier", func() {
				So(avq, ShouldNotBeNil)
			})
		})
	})
}

func TestGetVertexQuerier(t *testing.T) {
	Convey("Given a dialer, string executor and graph query manager", t, func() {
		dialer := gremconnect.NewWebSocketDialer("testaddress")
		execute := func(string, map[string]string, map[string]string) ([]byte, error) { return nil, nil }
		gm := NewGraphManager(dialer, logging.NewBasicLogger(), execute)
		Convey("When GetVertexQuerier is called", func() {
			gvq := gm.GetVertexQuerier()
			Convey("Then we should return the getVertex querier", func() {
				So(gvq, ShouldNotBeNil)
			})
		})
	})
}

func TestGetVertexIDQuerier(t *testing.T) {
	Convey("Given a dialer, string executor and graph query manager", t, func() {
		dialer := gremconnect.NewWebSocketDialer("testaddress")
		execute := func(string, map[string]string, map[string]string) ([]byte, error) { return nil, nil }
		gm := NewGraphManager(dialer, logging.NewBasicLogger(), execute)
		Convey("When GetVertexIDQuerier is called", func() {
			gvq := gm.GetVertexIDQuerier()
			Convey("Then we should return the getVertexID querier", func() {
				So(gvq, ShouldNotBeNil)
			})
		})
	})
}

func TestDropQuerier(t *testing.T) {
	Convey("Given a dialer, string executor and graph query manager", t, func() {
		dialer := gremconnect.NewWebSocketDialer("testaddress")
		execute := func(string, map[string]string, map[string]string) ([]byte, error) { return nil, nil }
		gm := NewGraphManager(dialer, logging.NewBasicLogger(), execute)
		Convey("When DropQuerier is called", func() {
			dq := gm.DropQuerier()
			Convey("Then we should return the drop querier", func() {
				So(dq, ShouldNotBeNil)
			})
		})
	})
}

func TestVertexQuerier(t *testing.T) {
	Convey("Given a dialer, string executor and graph query manager", t, func() {
		dialer := gremconnect.NewWebSocketDialer("testaddress")
		execute := func(string, map[string]string, map[string]string) ([]byte, error) { return nil, nil }
		gm := NewGraphManager(dialer, logging.NewBasicLogger(), execute)
		Convey("When VertexQuerier is called", func() {
			vq := gm.VertexQuerier()
			Convey("Then we should return the vertex querier", func() {
				So(vq, ShouldNotBeNil)
			})
		})
	})
}

func TestExecuteQuerier(t *testing.T) {
	Convey("Given a dialer, string executor and graph query manager", t, func() {
		dialer := gremconnect.NewWebSocketDialer("testaddress")
		execute := func(string, map[string]string, map[string]string) ([]byte, error) { return nil, nil }
		gm := NewGraphManager(dialer, logging.NewBasicLogger(), execute)
		Convey("When ExecuteQuerier is called", func() {
			eq := gm.ExecuteQuerier()
			Convey("Then we should return the execute querier", func() {
				So(eq, ShouldNotBeNil)
			})
		})
	})
}

func TestSchemaQuerier(t *testing.T) {
	Convey("Given a dialer, string executor and graph query manager", t, func() {
		dialer := gremconnect.NewWebSocketDialer("testaddress")
		execute := func(string, map[string]string, map[string]string) ([]byte, error) { return nil, nil }
		gm := NewGraphManager(dialer, logging.NewBasicLogger(), execute)
		Convey("When SchemaQuerier is called", func() {
			sq := gm.SchemaQuerier()
			Convey("Then we should return the schema querier", func() {
				So(sq, ShouldNotBeNil)
			})
		})
	})
}
