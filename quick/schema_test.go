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
	"github.com/northwesternmutual/grammes/logging"
	"github.com/northwesternmutual/grammes/manager"
	"github.com/northwesternmutual/grammes/query/cardinality"
	"github.com/northwesternmutual/grammes/query/datatype"
	"github.com/northwesternmutual/grammes/query/multiplicity"
)

func TestAddEdgeLabel(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, *time.Duration, map[string]string, map[string]string) ([][]byte, error) {
		return [][]byte{[]byte(vertexResponse)}, nil
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string, multiplicity and label", t, func() {
		host := "testhost"
		label := "testlabel"
		var m multiplicity.Multiplicity
		Convey("When AddEdgeLabel is called", func() {
			_, err := AddEdgeLabel(m, host, label)
			Convey("Then no errors should be thrown", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestAddEdgeLabelClientError(t *testing.T) {
	tempcheckForClient := checkForClient
	defer func() {
		checkForClient = tempcheckForClient
	}()
	checkForClient = func(string) error { return errors.New("ERROR") }
	Convey("Given a host string, multiplicity and label", t, func() {
		host := "testhost"
		label := "testlabel"
		var m multiplicity.Multiplicity
		Convey("When AddEdgeLabel is called and encounters an error checking for the client", func() {
			_, err := AddEdgeLabel(m, host, label)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddEdgeLabelQueryError(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, *time.Duration, map[string]string, map[string]string) ([][]byte, error) {
		return nil, errors.New("ERROR")
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string, multiplicity and label", t, func() {
		host := "testhost"
		label := "testlabel"
		var m multiplicity.Multiplicity
		Convey("When AddEdgeLabel is called and there is querying error", func() {
			_, err := AddEdgeLabel(m, host, label)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddEdgeLabels(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, *time.Duration, map[string]string, map[string]string) ([][]byte, error) {
		return [][]byte{[]byte(vertexResponse)}, nil
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string, multiplicity and labels", t, func() {
		host := "testhost"
		label := "testlabel"
		var m multiplicity.Multiplicity
		Convey("When AddEdgeLabel is called", func() {
			_, err := AddEdgeLabels(host, m, label)
			Convey("Then no errors should be thrown", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestAddEdgeLabelsClientError(t *testing.T) {
	tempcheckForClient := checkForClient
	defer func() {
		checkForClient = tempcheckForClient
	}()
	checkForClient = func(string) error { return errors.New("ERROR") }
	Convey("Given a host string, multiplicity and labels", t, func() {
		host := "testhost"
		label := "testlabel"
		var m multiplicity.Multiplicity
		Convey("When AddEdgeLabel is called and encounters an error checking for the client", func() {
			_, err := AddEdgeLabels(host, m, label)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddEdgeLabelsQueryError(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, *time.Duration, map[string]string, map[string]string) ([][]byte, error) {
		return nil, errors.New("ERROR")
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string, multiplicity and labels", t, func() {
		host := "testhost"
		label := "testlabel"
		var m multiplicity.Multiplicity
		Convey("When AddEdgeLabels is called and there is querying error", func() {
			_, err := AddEdgeLabels(host, m, label)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddPropertyKey(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, *time.Duration, map[string]string, map[string]string) ([][]byte, error) {
		return [][]byte{[]byte(vertexResponse)}, nil
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string, property name, datatype and cardinality", t, func() {
		host := "testhost"
		propertyName := "property"
		var d datatype.DataType
		var c cardinality.Cardinality
		Convey("When AddPropertyKey is called", func() {
			_, err := AddPropertyKey(host, propertyName, d, c)
			Convey("Then no errors should be thrown", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestAddPropertyKeyClientError(t *testing.T) {
	tempcheckForClient := checkForClient
	defer func() {
		checkForClient = tempcheckForClient
	}()
	checkForClient = func(string) error { return errors.New("ERROR") }
	Convey("Given a host string, property name, datatype and cardinality", t, func() {
		host := "testhost"
		propertyName := "property"
		var d datatype.DataType
		var c cardinality.Cardinality
		Convey("When AddPropertyKey is called and encounters an error checking for the client", func() {
			_, err := AddPropertyKey(host, propertyName, d, c)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddPropertyKeyQueryError(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, *time.Duration, map[string]string, map[string]string) ([][]byte, error) {
		return nil, errors.New("ERROR")
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string, property name, datatype and cardinality", t, func() {
		host := "testhost"
		propertyName := "property"
		var d datatype.DataType
		var c cardinality.Cardinality
		Convey("When AddPropertyKey is called and there is querying error", func() {
			_, err := AddPropertyKey(host, propertyName, d, c)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestCommitSchema(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, *time.Duration, map[string]string, map[string]string) ([][]byte, error) {
		return [][]byte{[]byte(vertexResponse)}, nil
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string", t, func() {
		host := "testhost"
		Convey("When CommitSchema is called", func() {
			_, err := CommitSchema(host)
			Convey("Then no errors should be thrown", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestCommitSchemaClientError(t *testing.T) {
	tempcheckForClient := checkForClient
	defer func() {
		checkForClient = tempcheckForClient
	}()
	checkForClient = func(string) error { return errors.New("ERROR") }
	Convey("Given a host string", t, func() {
		host := "testhost"
		Convey("When CommitSchema is called and encounters an error checking for the client", func() {
			_, err := CommitSchema(host)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestCommitSchemaQueryError(t *testing.T) {
	defer func() {
		client = nil
	}()
	dialer := &mockDialer{}
	client, _ = grammes.Dial(dialer)
	execute := func(string, *time.Duration, map[string]string, map[string]string) ([][]byte, error) {
		return nil, errors.New("ERROR")
	}
	client.GraphManager = manager.NewGraphManager(dialer, logging.NewNilLogger(), execute)
	Convey("Given a host string", t, func() {
		host := "testhost"
		Convey("When CommitSchema is called and there is querying error", func() {
			_, err := CommitSchema(host)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
