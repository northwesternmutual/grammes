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
Package grammes is an API/Wrapper for the Gremlin traversal language to interact with graph databases.
It includes various levels of functionality to add, remove, and change vertices and edges in the database.
Usage of higher level API is shown in various examples in the `examples/` directory with full documentation.

To get started with this package you may begin by making a Grammes client using
either the Dial function or the DialWithWebSocket function. With this client you may begin interacting
with your graph database with the client's multitude of function options.
To narrow down what you want to do it may be easier to choose one of the `client.Querier` options.

What this example does is create a new Grammes Client using the DialWithWebSocket function.
With this client it executes a simple string query that just does some simple addition. Then it will return the raw result out.

For further customizability you may check out packages within the `query/` directory.

To see examples on how to use this package further then check out the `examples/` directory.
*/
package grammes
