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
	"encoding/base64"
	"encoding/json"

	"github.com/google/uuid"
)

// Request is a container for all evaluation
// request parameters to be sent to the Gremlin Server.
type Request struct {
	RequestID string                 `json:"requestId"`
	Op        string                 `json:"op"`
	Processor string                 `json:"processor"`
	Args      map[string]interface{} `json:"args"`
}

var (
	// GenUUID is a monkey patched function for the Google UUIDv4 generator.
	GenUUID = uuid.NewUUID
	// jsonMarshal is a monkey patched function for the standard json.Marshal.
	jsonMarshal = json.Marshal
)

// PrepareRequest packages a query and binding
// into the format that Gremlin Server accepts
func PrepareRequest(query string, bindings, rebindings map[string]string) (req Request, id string, err error) {
	var guuid uuid.UUID

	if guuid, err = GenUUID(); err != nil {
		return
	}
	id = guuid.String()

	req.RequestID = id
	req.Op = "eval"
	req.Processor = ""

	req.Args = make(map[string]interface{})
	req.Args["language"] = "gremlin-groovy"
	req.Args["gremlin"] = query
	req.Args["bindings"] = bindings
	req.Args["rebindings"] = rebindings

	return
}

// PackageRequest takes a request type and formats
// it into being able to be delivered to the TinkerPop server.
func PackageRequest(req Request, versionNumber string) (msg []byte, err error) {
	j, err := jsonMarshal(req) // Formats request into byte format
	if err != nil {
		return
	}

	// ! is the prefix for the mimeType header.
	msg = []byte("!application/vnd.gremlin-v" + versionNumber + ".0+json")
	msg = append(msg, j...)

	return
}

// PrepareAuthRequest creates a request for
// the dialer to send to the Gremlin-Server.
func PrepareAuthRequest(requestID string, username string, password string) (req Request, err error) {
	req.RequestID = requestID
	req.Op = "authentication"
	req.Processor = "traversal"

	simpleAuth := make([]byte, len(username)+len(password)+2)

	copy(simpleAuth[1:], username)
	copy(simpleAuth[len(username)+2:], password)

	req.Args = make(map[string]interface{})
	req.Args["sasl"] = base64.StdEncoding.EncodeToString(simpleAuth)

	return
}
