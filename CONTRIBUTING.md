# How to Contribute to Grammes

## Getting Started

This project uses Go modules to manage dependencies, and is built on Go version `1.11`.

```sh
==> go get github.com/northwesternmutual/grammes
```

## Testing

Grammes unit tests can be run by directing yourself to the root directory for the project and running this command in a terminal:

```sh
==> go test -parallel 3 --race ./...
```

## Project Structure

Below is an overview of this project's structure (files are not listed)

```flat
github.com/northwesternmutual/grammes
  assets/             - Contains images and shell scripts for customizing JanusGraph
  docs/               - Miscellaneous documents and charts to help explain Grammes
  examples/           - Examples of all resources on how to use Grammes
  gremconnect/        - The package for handling connections and dialers
  gremerror/          - Contains custom error types for Grammes
  logging/            - Has the custom loggers for the Grammes client
  manager/            - Contains the GraphManager and all graph related functions
  model/              - Holds all of the structs and methods for Go counterpart objects
  query/              - Contains query related packages and the Query interface
    cardinality/      - Describes the number of relationship occurrences
    column/           - Used to reference particular parts of a column in Gremlin
    consumer/         - Controls how barriers emit their values
    datatype/         - String counterparts to reference Gremlin types
    direction/        - Denotes the direction of edges
    graph/            - The verbose graph traversal object (Query)
    multiplicity/     - Controls property sets for edges
    operator/         - Used to apply mathematical operations in a graph traversal
    pop/              - Determines the gathered values
    predicate/        - Controls the search conditions of a query
    scope/            - Controls the relations of a graph traversal
    token/            - Defines the parts of a vertex
    traversal/        - The preferred graph traversal (Query)
  quick/              - Used when executing queries without a Grammes client
  testing/            - Has the mock for the GraphManager when testing Grammes in your project
```

## Imports Grouping

This projects adheres to the following pattern when grouping imports in Go files:

* imports from standard library
* imports from other projects
* imports from internal project

In addition, imports in each group must be sorted by length. For example:

```go
import (
  "time"
  "errors"

  "go.uber.org/zap"

  "github.com/northwesternmutual/grammes/query/token"
  "github.com/northwesternmutual/grammes"
)
```

## Making a Change

Before making any significant changes, please [open an issue](https://github.com/northwesternmutual/grammes/issues). Discussing your proposed changes ahead of time will make the contribution process smooth for everyone.

Once we've discussed your changes and you've got your code ready, make sure that tests are passing (go test) and open your PR. Your pull request is most likely to be accepted if it:

* Includes tests for new functionality.
* Follows the guidelines in [Effective Go](https://golang.org/doc/effective_go.html) and the [Go team's common code review comments](https://github.com/golang/go/wiki/CodeReviewComments).
* Has a good commit message.

## License

By contributing your code, you agree to license your contribution under the terms of the [Apache License](https://github.com/northwesternmutual/grammes/blob/master/LICENSE).

If you are adding a new file it should have a header like below.

```flat
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
```