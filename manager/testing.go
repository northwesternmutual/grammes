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

package manager

var (
	readCount   int
	connect     error
	isConnected = true
	isDisposed  = false
	response    = vertexResponse

	vertexResponse = `
	[
		{
			"@type": "g:Vertex",
			"@value": {
				"id": 28720,
				"label": "newvertex"
			}
		}
	]
	`
	edgeResponse = `
	{
		"requestId": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
		"status": {
			"message": "",
			"code": 200,
			"attributes": {}
		},
		"result": {
			"data": [{
				"@type": "g:Edge",
				"@value": {
					"id": {
						"@type": "janusgraph:RelationIdentifier",
						"@value": {
							"relationId": "zz0-yxs-25rf9-1548"
						}
					},
					"inV": {
						"@type": "g:Int64",
						"@value": 53288
					},
					"inVLabel": "person2",
					"label": "friendsWith",
					"outV": {
						"@type": "g:Int64",
						"@value": 45280
					},
					"outVLabel": "person1",
					"properties": {
						"ageDiff": {
							"@type": "g:Property",
							"@value": {
								"key": "ageDiff",
								"value": {
									"@type": "g:Int32",
									"@value": 12
								}
							}
						}
					}
				}
			}],
			"meta": {}
		}
	}
	`
	idResponse = `
	[
		{
			"@type": "g:Id",
			"@value": 255
		}
	]
	`
	propertyResponse = `
	{
		"requestId": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
		"status": {
			"message": "",
			"code": 200,
			"attributes": {}
		},
		"result": {
			"data": [{
				"@type": "g:VertexProperty",
				"@value": {
					"id": {
						"@type": "janusgraph:RelationIdentifier",
						"@value": {
							"relationId": "oe6tp-oefzc-3c3p"
						}
					},
					"label": "name",
					"value": "damien"
				}
			}],
			"meta": {}
		}
	}
	`
	propertyMapResponse = `
	{
		"requestId": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
		"status": {
			"message": "",
			"code": 200,
			"attributes": {}
		},
		"result": {
			"data": [{
				"name": [{
					"@type": "g:VertexProperty",
					"@value": {
						"id": {
							"@type": "janusgraph:RelationIdentifier",
							"@value": {
								"relationId": "oe6tp-oefzc-3c3p"
							}
						},
						"label": "name",
						"value": "damien"
					}
				}]
			}],
			"meta": {}
		}
	}
	`
	valuesResponse = `
	{
		"requestId": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
		"status": {
			"message": "",
			"code": 200,
			"attributes": {}
		},
		"result": {
			"data": ["damien"],
			"meta": {}
		}
	}
	`
	valueMapResponse = `
	{
		"requestId": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
		"status": {
			"message": "",
			"code": 200,
			"attributes": {}
		},
		"result": {
			"data": [
				{
					"name":["damien"]
				}],
			"meta": {}
		}
	}
	`
	emptyResponse = `
	{
		"requestId": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
		"status": {
			"message": "",
			"code": 200,
			"attributes": {}
		},
		"result": {
			"data": [],
			"meta": {}
		}
	}
	`
)
