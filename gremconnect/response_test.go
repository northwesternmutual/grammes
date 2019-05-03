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

package gremconnect

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/northwesternmutual/grammes/gremerror"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	response200 = `
{
    "requestId": "d2476e5b-b2bc-6a70-2647-3991f68ab415",
    "status": {
        "code": 200,
        "attributes": {}
    },
    "result": {
        "data": [{
        }],
        "meta": {}
    }
}
`
	response204 = `
{
    "requestId": "d2476e5b-b2bc-6a70-2647-3991f68ab415",
    "status": {
        "code": 204,
        "attributes": {}
    },
    "result": {
        "data": [{
        }],
        "meta": {}
    }
}
`
	response206 = `
{
    "requestId": "d2476e5b-b2bc-6a70-2647-3991f68ab415",
    "status": {
        "code": 206,
        "attributes": {}
    },
    "result": {
        "data": [{
        }],
        "meta": {}
    }
}
`
	response401 = `
{
    "requestId": "d2476e5b-b2bc-6a70-2647-3991f68ab415",
    "status": {
        "code": 401,
        "attributes": {}
    },
    "result": {
    }
}
`
	response407 = `
{
    "requestId": "d2476e5b-b2bc-6a70-2647-3991f68ab415",
    "status": {
        "code": 407,
        "attributes": {}
    },
    "result": {
    }
}
`
	response498 = `
{
    "requestId": "d2476e5b-b2bc-6a70-2647-3991f68ab415",
    "status": {
        "code": 498,
        "attributes": {}
    },
    "result": {
    }
}
`
	response499 = `
{
    "requestId": "d2476e5b-b2bc-6a70-2647-3991f68ab415",
    "status": {
        "code": 499,
        "attributes": {}
    },
    "result": {
    }
}
`
	response500 = `
{
    "requestId": "d2476e5b-b2bc-6a70-2647-3991f68ab415",
    "status": {
        "code": 500,
        "attributes": {}
    },
    "result": {
    }
}
`
	response597 = `
{
    "requestId": "d2476e5b-b2bc-6a70-2647-3991f68ab415",
    "status": {
        "code": 597,
        "attributes": {}
    },
    "result": {
    }
}
`
	response598 = `
{
    "requestId": "d2476e5b-b2bc-6a70-2647-3991f68ab415",
    "status": {
        "code": 598,
        "attributes": {}
    },
    "result": {
    }
}
`
	response599 = `
{
    "requestId": "d2476e5b-b2bc-6a70-2647-3991f68ab415",
    "status": {
        "code": 599,
        "attributes": {}
    },
    "result": {
    }
}
`

	responseDefault = `
{
    "requestId": "d2476e5b-b2bc-6a70-2647-3991f68ab415",
    "status": {
        "code": 403,
        "attributes": {}
    },
    "result": {
    }
}
`
)

func TestMarshalRespone(t *testing.T) {
	Convey("Given a valid []byte response", t, func() {
		byteResponse200 := []byte(response200)
		byteResponse204 := []byte(response204)
		byteResponse206 := []byte(response206)
		byteResponse401 := []byte(response401)
		byteResponse407 := []byte(response407)
		byteResponse498 := []byte(response498)
		byteResponse499 := []byte(response499)
		byteResponse500 := []byte(response500)
		byteResponse597 := []byte(response597)
		byteResponse598 := []byte(response598)
		byteResponse599 := []byte(response599)
		byteResponseDefault := []byte(responseDefault)

		Convey("And the raw response is marshalled", func() {
			resp200, err200 := MarshalResponse(byteResponse200)
			resp204, err204 := MarshalResponse(byteResponse204)
			resp206, err206 := MarshalResponse(byteResponse206)
			resp401, _ := MarshalResponse(byteResponse401)
			resp407, _ := MarshalResponse(byteResponse407)
			resp498, _ := MarshalResponse(byteResponse498)
			resp499, _ := MarshalResponse(byteResponse499)
			resp500, _ := MarshalResponse(byteResponse500)
			resp597, _ := MarshalResponse(byteResponse597)
			resp598, _ := MarshalResponse(byteResponse598)
			resp599, _ := MarshalResponse(byteResponse599)
			respDefault, _ := MarshalResponse(byteResponseDefault)

			Convey("Then the status code should be 200", func() {
				So(resp200.Code, ShouldEqual, 200)
			})

			Convey("And err200 should be nil", func() {
				So(err200, ShouldBeNil)
			})

			Convey("Then the status code should be 204", func() {
				So(resp204.Code, ShouldEqual, 204)
			})

			Convey("And err204 should be nil", func() {
				So(err204, ShouldBeNil)
			})

			Convey("Then the status code should be 206", func() {
				So(resp206.Code, ShouldEqual, 206)
			})

			Convey("And err206 should be nil", func() {
				So(err206, ShouldBeNil)
			})

			Convey("Then the status code should be 401", func() {
				So(resp401.Code, ShouldEqual, 401)
			})

			Convey("And resp401.Data should be 'UNAUTHORIZED'", func() {
				So(resp401.Data, ShouldResemble, gremerror.NewNetworkError(401, "UNAUTHORIZED"))
			})

			Convey("Then the status code should be 407", func() {
				So(resp407.Code, ShouldEqual, 407)
			})

			Convey("And resp407.Data should be 'AUTHENTICATE'", func() {
				So(resp407.Data, ShouldResemble, gremerror.NewNetworkError(407, "AUTHENTICATION REQUIRED"))
			})

			Convey("Then the status code should be 498", func() {
				So(resp498.Code, ShouldEqual, 498)
			})

			Convey("And resp498.Data should be 'MALFORMED REQUEST", func() {
				So(resp498.Data, ShouldResemble, gremerror.NewNetworkError(498, "MALFORMED REQUEST"))
			})

			Convey("Then the status code should be 499", func() {
				So(resp499.Code, ShouldEqual, 499)
			})

			Convey("And resp499.Data should be 'INVALID REQUEST ARGUMENTS'", func() {
				So(resp499.Data, ShouldResemble, gremerror.NewNetworkError(499, "INVALID REQUEST ARGUMENTS"))
			})

			Convey("Then the status code should be 500", func() {
				So(resp500.Code, ShouldEqual, 500)
			})

			Convey("And resp500.Data should be 'SERVER ERROR'", func() {
				So(resp500.Data, ShouldResemble, gremerror.NewNetworkError(500, "INTERNAL SERVER ERROR"))
			})

			Convey("Then the status code should be 597", func() {
				So(resp597.Code, ShouldEqual, 597)
			})

			Convey("And resp597.Data should be 'SCRIPT EVALUATION ERROR'", func() {
				So(resp597.Data, ShouldResemble, gremerror.NewNetworkError(597, "SCRIPT EVALUATION ERROR"))
			})

			Convey("Then the status code should be 598", func() {
				So(resp598.Code, ShouldEqual, 598)
			})

			Convey("And resp598.Data should be 'SERVER TIMEOUT'", func() {
				So(resp598.Data, ShouldResemble, gremerror.NewNetworkError(598, "SERVER TIMEOUT"))
			})

			Convey("Then the status code should be 599", func() {
				So(resp599.Code, ShouldEqual, 599)
			})

			Convey("And resp599.Data should be 'SERVER SERIALIZATION ERROR'", func() {
				So(resp599.Data, ShouldResemble, gremerror.NewNetworkError(599, "SERIALIZATION ERROR"))
			})

			Convey("Then respDefault.Data should be 'UNKNOWN ERROR'", func() {
				So(respDefault.Data, ShouldResemble, gremerror.NewNetworkError(403, "UNKNOWN ERROR"))
			})
		})
	})
}

func TestMarshalResponseUnmarshalError(t *testing.T) {
	defer func() {
		jsonUnmarshal = json.Unmarshal
	}()
	testErr := errors.New("ERROR")
	jsonUnmarshal = func([]byte, interface{}) error { return testErr }
	Convey("Given a byte message", t, func() {
		var msg []byte
		Convey("And MarshalResponse throws an error while unmarshalling", func() {
			_, err := MarshalResponse(msg)
			Convey("Then err should equal the test error", func() {
				So(err, ShouldResemble, gremerror.NewUnmarshalError("MarshalResponse", msg, testErr))
			})
		})
	})
}
