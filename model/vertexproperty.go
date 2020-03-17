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

// Tinkerpop:
// http://tinkerpop.apache.org/javadocs/3.2.1/core/org/apache/tinkerpop/gremlin/structure/Property.html

// PropertyMap is the map used to hold
// the properties itself. the string key is equivalent
// to the Gremlin key and the []Property is the value.
// Properties can have multiple values; this is why we must
// have it as a slice of Property.
type PropertyMap map[string][]Property

// Property holds the type and
// value of the property. It's extra
// information used by PropertyDetail.
type Property struct {
	// Type is the Gremlin-type. For this particular
	// structure it typically will be "g:VertexProperty"
	Type string `json:"@type"`
	// Value stores the actual data of the Property.
	// This would be its Key, Value, and ID.
	Value PropertyValue `json:"@value"`
}

// NewProperty will just shorten the struggle of filling
// a property struct. This is meant to be used when creating a Vertex struct.
func NewProperty(label string, value interface{}) Property {
	return Property{
		Value: PropertyValue{
			Label: label,
			Value: ValueWrapper{
				PropertyDetailedValue: PropertyDetailedValue{
					Value: value,
				},
			},
		},
	}
}

// GetValue is a shortcut for taking the raw interface{}
// from the property itself without redundancy.
func (p *Property) GetValue() interface{} {
	return p.Value.Value.Value
}

// GetLabel will return the key that is used to find
// this particular property and its value.
func (p *Property) GetLabel() string {
	return p.Value.Label
}

// PropertyValue contains the ID,
// value, and label of this property's value.
type PropertyValue struct {
	ID    PropertyID   `json:"id"`
	Value ValueWrapper `json:"value"`
	Label string       `json:"label"`
}

// PropertyID holds the ID that is used
// for the property itself.
type PropertyID struct {
	Type  string      `json:"@type"`
	Value interface{} `json:"@value"`
}
