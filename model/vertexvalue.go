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

package model

import "encoding/json"

// VertexValue contains the 'value' data
// from the Vertex object.
type VertexValue struct {
	ID         interface{} `json:"id"`
	Label      string      `json:"label"`
	Properties PropertyMap `json:"properties,omitempty"`
}

// PropertyDetailedValue holds the value
// and optional type depending on how
// the whole struct can unmarshal into a string
// or not. If not then the type is listed.
type PropertyDetailedValue struct {
	Value interface{} `json:"@value"`
	Type  string      `json:"@type,omitempty"`
}

// ValueWrapper will handle storing
// the correct value into the correct
// spot in this struct for not confusion.
type ValueWrapper struct {
	PropertyDetailedValue
	Partial bool `json:"-"`
}

// UnmarshalJSON will override the unmarshal
// process of ValueWrapper and store the correct
// Value into the variables within the struct.
func (w *ValueWrapper) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &w.PropertyDetailedValue); err == nil {
		return nil
	}

	w.Partial = true
	return json.Unmarshal(data, &w.Value)
}
