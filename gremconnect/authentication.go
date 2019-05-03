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
	"errors"
	"os"
)

// Auth contains the authentication data for the dialer.
type Auth struct {
	Username string
	Password string
}

// AuthInfo holds all the information needed to use SASL
// authentication with the Gremlin server.
// ChallengeID is the requestID in the 407 status (AUTHENTICATE) response
type AuthInfo struct {
	ChallengeID string
	User        string
	Pass        string
}

// OptAuth holds data about the authentication
type OptAuth func(*AuthInfo) error

// Moving this function to a local var so we can mock it in tests
var osLookupEnv = os.LookupEnv

// OptAuthEnv sets authentication info from environment variables GREMLIN_USER and GREMLIN_PASS
func OptAuthEnv() OptAuth {
	return func(auth *AuthInfo) error {
		user, ok := osLookupEnv("GREMLIN_USER")
		if !ok {
			return errors.New("variable GREMLIN_USER is not set")
		}
		pass, ok := osLookupEnv("GREMLIN_PASS")
		if !ok {
			return errors.New("variable GREMLIN_PASS is not set")
		}
		auth.User = user
		auth.Pass = pass
		return nil
	}
}

// OptAuthUserPass sets authentication information from username and password
func OptAuthUserPass(user, pass string) OptAuth {
	return func(auth *AuthInfo) error {
		auth.User = user
		auth.Pass = pass
		return nil
	}
}
