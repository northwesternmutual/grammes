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

	"github.com/northwesternmutual/grammes/gremerror"
)

// Response is the structure representation of
// the response json received from the Gremlin-Server.
type Response struct {
	Data      interface{}
	RequestID string
	Code      int
}

// patch for mocking test values
var jsonUnmarshal = json.Unmarshal

// MarshalResponse creates a Response struct for
// every incoming Response for further manipulation
func MarshalResponse(msg []byte) (Response, error) {
	var j map[string]interface{}

	err := jsonUnmarshal(msg, &j)
	if err != nil {
		return Response{}, gremerror.NewUnmarshalError("MarshalResponse", msg, err)
	}

	var (
		status = j["status"].(map[string]interface{})
		result = j["result"].(map[string]interface{})
		code   = status["code"].(float64)
		resp   = Response{Code: int(code)}
	)
	message, _ := status["message"].(string)

	err = responseDetectError(resp.Code, message)
	if err != nil {
		resp.Data = err // Use the Data field as a vehicle for the error.
	} else {
		resp.Data = result["data"]
	}
	resp.RequestID = j["requestId"].(string)

	return resp, nil
}

// responseDetectError detects any possible errors in responses
// from Gremlin Server and generates an error for each code
func responseDetectError(code int, message string) error {
	switch code {
	case 200:
		break
	case 204:
		break
	case 206:
		break
	case 401:
		return gremerror.NewNetworkError(401, "UNAUTHORIZED", message)
	case 407:
		return gremerror.NewNetworkError(407, "AUTHENTICATION REQUIRED", message)
	case 498:
		return gremerror.NewNetworkError(498, "MALFORMED REQUEST", message)
	case 499:
		return gremerror.NewNetworkError(499, "INVALID REQUEST ARGUMENTS", message)
	case 500:
		return gremerror.NewNetworkError(500, "INTERNAL SERVER ERROR", message)
	case 503:
		return gremerror.NewNetworkError(503, "SERVER UNAVAILABLE", message)
	case 597:
		return gremerror.NewNetworkError(597, "SCRIPT EVALUATION ERROR", message)
	case 598:
		return gremerror.NewNetworkError(598, "SERVER TIMEOUT", message)
	case 599:
		return gremerror.NewNetworkError(599, "SERIALIZATION ERROR", message)
	default:
		return gremerror.NewNetworkError(code, "UNKNOWN ERROR", message)
	}
	return nil
}
