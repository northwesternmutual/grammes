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
	"github.com/northwesternmutual/grammes/query/cardinality"
	"github.com/northwesternmutual/grammes/query/datatype"
	"github.com/northwesternmutual/grammes/query/multiplicity"
)

// AddEdgeLabel adds the edge label to the
// graph directly. This method returns the schema id
// of the edge label added.
func AddEdgeLabel(multi multiplicity.Multiplicity, host, label string) (int64, error) {
	err := checkForClient(host)
	if err != nil {
		return 0, err
	}

	sq := client.GraphManager.SchemaQuerier()
	res, err := sq.AddEdgeLabel(multi, label)
	if err != nil {
		return 0, err
	}

	return res, nil
}

// AddEdgeLabels does the same thing as AddEdgeLabel
// but with the ability to do multiple labels at a
// time. This function is called similarly to your
// favorite logger.
func AddEdgeLabels(host string, multiplicityAndLabels ...interface{}) ([]int64, error) {
	err := checkForClient(host)
	if err != nil {
		return nil, err
	}

	sq := client.GraphManager.SchemaQuerier()
	res, err := sq.AddEdgeLabels(multiplicityAndLabels...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// AddPropertyKey adds the edge label to the
// graph directly. This method returns the schema id
// of the edge label added.
func AddPropertyKey(host, propertyName string, datatype datatype.DataType, cardinality cardinality.Cardinality) (int64, error) {
	err := checkForClient(host)
	if err != nil {
		return 0, err
	}

	sq := client.GraphManager.SchemaQuerier()
	res, err := sq.AddPropertyKey(propertyName, datatype, cardinality)
	if err != nil {
		return 0, err
	}

	return res, nil
}

// CommitSchema will take all of your schema changes
// and apply them to the schema once they are ready.
func CommitSchema(host string) ([][]byte, error) {
	err := checkForClient(host)
	if err != nil {
		return nil, err
	}

	sq := client.GraphManager.SchemaQuerier()
	res, err := sq.CommitSchema()
	if err != nil {
		return nil, err
	}

	return res, nil
}
