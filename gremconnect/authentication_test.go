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
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOptAuthEnvInvalidUsername(t *testing.T) {
	tempOsLookupEnv := osLookupEnv
	defer func() {
		osLookupEnv = tempOsLookupEnv
	}()
	osLookupEnv = func(key string) (string, bool) { return "", false }

	Convey("Given an AuthInfo object", t, func() {
		auth := &AuthInfo{}

		Convey("and OptAuthEnv() is called, along with its return function", func() {
			optAuth := OptAuthEnv()
			err := optAuth(auth)
			expectedErr := errors.New("variable GREMLIN_USER is not set")

			Convey("Then err should not be nil and the auth username and password should not be set", func() {
				So(err, ShouldResemble, expectedErr)
				So(auth.User, ShouldEqual, "")
				So(auth.Pass, ShouldEqual, "")
			})
		})
	})
}

func TestOptAuthEnvInvalidPassword(t *testing.T) {
	tempOsLookupEnv := osLookupEnv
	defer func() {
		osLookupEnv = tempOsLookupEnv
	}()
	osLookupEnv = func(key string) (string, bool) {
		if key == "GREMLIN_PASS" {
			return "", false
		}
		return "", true
	}

	Convey("Given an AuthInfo object", t, func() {
		auth := &AuthInfo{}

		Convey("and OptAuthEnv() is called, along with its return function", func() {
			optAuth := OptAuthEnv()
			err := optAuth(auth)
			expectedErr := errors.New("variable GREMLIN_PASS is not set")

			Convey("Then err should not be nil and the auth username and password should not be set", func() {
				So(err, ShouldResemble, expectedErr)
				So(auth.User, ShouldEqual, "")
				So(auth.Pass, ShouldEqual, "")
			})
		})
	})
}

func TestOptAuthEnvValid(t *testing.T) {
	tempOsLookupEnv := osLookupEnv
	defer func() {
		osLookupEnv = tempOsLookupEnv
	}()
	osLookupEnv = func(key string) (string, bool) { return "testing", true }

	Convey("Given an AuthInfo object", t, func() {
		auth := &AuthInfo{}

		Convey("and OptAuthEnv() is called, along with its return function", func() {
			optAuth := OptAuthEnv()
			err := optAuth(auth)

			Convey("Then err should be nil and the auth username and password should be set", func() {
				So(err, ShouldBeNil)
				So(auth.User, ShouldEqual, "testing")
				So(auth.Pass, ShouldEqual, "testing")
			})
		})
	})
}

func TestOptAuthUserPass(t *testing.T) {
	Convey("Given an AuthInfo object, username and password", t, func() {
		auth := &AuthInfo{}
		user := "testuser"
		pass := "testpass"

		Convey("And we try to authenticate", func() {
			optAuth := OptAuthUserPass(user, pass)
			err := optAuth(auth)

			Convey("Then the user and pass should be set in the auth object", func() {
				So(err, ShouldBeNil)
				So(auth.User, ShouldEqual, "testuser")
				So(auth.Pass, ShouldEqual, "testpass")
			})
		})
	})
}
