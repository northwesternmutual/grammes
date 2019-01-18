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

/*
Package multiplicity contains the object to control property sets for an edge.

See:
	Titan: http://titan.thinkaurelius.com/javadoc/1.0.0/com/thinkaurelius/titan/core/Multiplicity.html
	Wikipedia: http://en.wikipedia.org/wiki/Class_diagram#Multiplicity
	Object-oriented Systems Modeling Laboratory (OSM): http://osm7.cs.byu.edu/OSA/participConst.html

Multiplicity acts as a property set for an edge.
This includes the number of arrows and vertices of association there are.

A note about Multiplicity:

This object implements the Parameter interfaces used by graph traversals.
*/
package multiplicity

import (
	"github.com/northwesternmutual/grammes/query/cardinality"
	"github.com/northwesternmutual/grammes/query/direction"
)

// Titan:
// http://titan.thinkaurelius.com/javadoc/1.0.0/com/thinkaurelius/titan/core/Multiplicity.html

// Wikipedia:
// http://en.wikipedia.org/wiki/Class_diagram#Multiplicity

// Object-oriented Systems Modeling Laboratory (OSM):
// http://osm7.cs.byu.edu/OSA/participConst.html

// Multiplicity acts as a property
// set for an edge. This includes the
// arrow being associated with the
// direction of the edge. This is essentially
// the cardinality and participation.
type Multiplicity string

const (
	// Many2One shows that many vertices
	// connect to a single vertex. Similar
	// to One2Many but in the opposite direction.
	// Example:
	// person <-- --> birth place
	// left      | right
	// 1..* or + | 1
	Many2One Multiplicity = "MANY2ONE"
	// Many2Many shows that many vertices
	// connect to many other vertices. Similar
	// to Many2One, but instead of just one vertex
	// it's all just a cluster of vertices.
	// Example:
	// person <-- --> book
	// left      | right
	// 0..* or * | 0..* or *
	Many2Many Multiplicity = "MANY2MANY"
	// One2Many describes an edge that
	// comes from multiple vertices and
	// point towards one vertex.
	// order  <-- --> line item
	// left      | right
	// 1         | 1..* or +
	One2Many Multiplicity = "ONE2MANY"
	// One2One explains that the vertex
	// being associated connects to one
	// going in a single direction.
	// person <-- --> birth certificate
	// left      | right
	// 1         | 1
	One2One Multiplicity = "ONE2ONE"
	// Simple does not have an associated
	// direction belonging to it. This
	// results in an edge with no arrows.
	Simple Multiplicity = "SIMPLE"
	// Multi is an arbitrary multiplicity
	// explaining that this edge will have
	// multiple arrows associated with it.
	// Wikipedia:
	// http://en.wikipedia.org/wiki/Multigraph
	Multi Multiplicity = "MULTI"
)

// IsConstrained returns whether this multiplicity imposes
// any constraint on the number of edges there may exist between a
// pair of vertices.
func (m Multiplicity) IsConstrained() bool {
	return m != Multi
}

// IsConstrainedDirection returns whether this multiplicity imposes
// any constraint on the number of edges there may exist between a
// pair of vertices.
func (m Multiplicity) IsConstrainedDirection(d direction.Direction) bool {
	if d == direction.Both {
		return m != Multi
	}
	return m != Multi && (m == Simple || m.IsUnique(d))
}

// IsUnique returns true if this multiplicity implies edge
// uniqueness in the given direction for any given vertex.
func (m Multiplicity) IsUnique(d direction.Direction) bool {
	switch d {
	case direction.In:
		return m == One2Many || m == One2One
	case direction.Out:
		return m == Many2One || m == One2One
	case direction.Both:
		return m == One2One
	default:
		return false
	}
}

// Convert will take a cardinality and return a matching Multiplicity.
func Convert(c cardinality.Cardinality) Multiplicity {
	switch c {
	case cardinality.List:
		return Multi
	case cardinality.Set:
		return Simple
	case cardinality.Single:
		return Many2One
	default:
		return Multiplicity("Invalid Cardinality: " + c.String())
	}
}

// String will convert the Multiplicity into a string.
func (m Multiplicity) String() string {
	return string(m)
}

// Cardinality will convert the Multiplicity to a cardinality.
func (m Multiplicity) Cardinality() cardinality.Cardinality {
	switch m {
	case Multi:
		return cardinality.List
	case Simple:
		return cardinality.Set
	case Many2One:
		return cardinality.Single
	default:
		return cardinality.Cardinality("Invalid Multiplicity: " + m.String())
	}
}
