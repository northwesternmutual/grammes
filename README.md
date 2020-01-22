<p align="center">
  <img src="assets/img/grammes-gopher-v4.png" alt="Grammes" title="Grammes" width="50%">
</p>

# Grammes

<p align="center">
<a href="https://godoc.org/github.com/northwesternmutual/grammes"><img src="https://godoc.org/github.com/northwesternmutual/grammes?status.svg" alt="GoDoc" /></a>
<a href="https://github.com/northwesternmutual/grammes/releases"><img src="https://badgen.net/github/release/northwesternmutual/grammes"></a>
<a href="https://goreportcard.com/report/github.com/northwesternmutual/grammes"><img src="https://goreportcard.com/badge/github.com/northwesternmutual/grammes" alt="Go Report Card" /></a>
<a href="https://github.com/northwesternmutual/grammes/blob/master/LICENSE"><img src="https://img.shields.io/github/license/northwesternmutual/grammes.svg" alt="License" /></a>
<a href="https://github.com/northwesternmutual/grammes/actions"><img src="https://github.com/northwesternmutual/grammes/workflows/Build/badge.svg" /></a>
</p>

Grammes is an API/Wrapper for Gremlin and Janusgraph. It's written purely in Golang and allows for easy use of Gremlin without touching the Gremlin terminal.

## Table of Contents

- [Grammes](#grammes)
  - [Table of Contents](#table-of-contents)
  - [Local Setup](#local-setup)
    - [Cloning Grammes](#cloning-grammes)
    - [Setting up JanusGraph](#setting-up-janusgraph)
    - [Using Grammes](#using-grammes)
  - [Testing Grammes](#testing-grammes)
  - [Additional Resources](#additional-resources)
    - [Documentation on Gremlin](#documentation-on-gremlin)
    - [Examples](#examples)
  - [Troubleshooting](#troubleshooting)
    - [Fixing time outs when starting Janusgraph](#fixing-time-outs-when-starting-janusgraph)

## Local Setup

You need to setup all of the following tools to run the service locally

- Go 1.12
- Git
- Elastic Search
- Cassandra
  - Java 8

---

### Cloning Grammes

Begin by opening up a terminal or command prompt and clone the grammes repository.

```sh
go get -u github.com/northwesternmutual/grammes
```

---

### Setting up JanusGraph

*if you have decided to use another graph database then you may move on to [project setup](#using-grammes)*

First off, direct your terminal to the Grammes' `scripts` directory.

```sh
cd $GOPATH/src/github.com/northwesternmutual/grammes/scripts
```

In here you can find the `gremlin.sh` and `janusgraph.sh` scripts. To set up JanusGraph just run the `janusgraph.sh` script.

```sh
./janusgraph.sh
```

This should download and/or begin the graph and TinkerPop server.

To make sure that everything is running try running `gremlin.sh`.

```sh
$ ./gremlin.sh
SLF4J: Class path contains multiple SLF4J bindings.
SLF4J: Found binding in [jar:file:/Users/<username>/projects/nm/gocode/src/github.com/northwesternmutual/grammes/bin/janusgraph-0.3.1-hadoop2/lib/slf4j-log4j12-1.7.12.jar!/org/slf4j/impl/StaticLoggerBinder.class]
SLF4J: Found binding in [jar:file:/Users/<username>/projects/nm/gocode/src/github.com/northwesternmutual/grammes/bin/janusgraph-0.3.1-hadoop2/lib/logback-classic-1.1.2.jar!/org/slf4j/impl/StaticLoggerBinder.class]
SLF4J: See http://www.slf4j.org/codes.html#multiple_bindings for an explanation.
SLF4J: Actual binding is of type [org.slf4j.impl.Log4jLoggerFactory]
15:05:59 WARN  org.apache.hadoop.util.NativeCodeLoader  - Unable to load native-hadoop library for your platform... using builtin-java classes where applicable
gremlin>
```

Once Gremlin starts then you may begin by running this command.

```sh
gremlin> :remote connect tinkerpop.server conf/remote.yaml
===>Configured localhost/127.0.0.1:8182
```

If you see the message that Gremlin was configured to the localhost then quit Gremlin.

```sh
gremlin> :exit
```

Finally, run the `janusgraph.sh` script again, but this time with the `status` flag.

```sh
./janusgraph.sh status
```

---

### Using Grammes

Once you have cloned the repository then you may begin implementing it into a project. Let's begin with creating a place for your code in the `$GOPATH`, i.e.,

```sh
$GOPATH/src/github.com/<username-here>/<project-name-here>
```

Next, you'll want to create a `main.go` file. For this example I will be using [MS Code](https://code.visualstudio.com/download), but you may use any editor you prefer.

```sh
code main.go
```

In this file we can begin by making it a typical empty `main.go` file like this:

``` go
package main

func main() {
}
```

Next, import the grammes package and begin using it by connecting your client to a gremlin server.

``` go
package main

import (
    "log"

    "github.com/northwesternmutual/grammes"
)

func main() {
    // Creates a new client with the localhost IP.
    client, err := grammes.DialWithWebSocket("ws://127.0.0.1:8182")
    if err != nil {
        log.Fatalf("Error while creating client: %s\n", err.Error())
    }

    // Executing a basic query to assure that the client is working.
    res, err := client.ExecuteStringQuery("1+3")
    if err != nil {
        log.Fatalf("Querying error: %s\n", err.Error())
    }

    // Print out the result as a string
    for _, r := range res {
        log.Println(string(r))
    }
}
```

Once the client is created then you can begin querying the gremlin server via the `.ExecuteQuery` method in the client. To use this function you must put in a `Query` which is an interface for any kind of Query-able type in the package. These include: `graph.String` and `traversal.String`. For an example of querying the gremlin server for all of the Vertex labels:

``` go
package main

import (
    "log"

    "github.com/northwesternmutual/grammes"
)

func main() {
    // Creates a new client with the localhost IP.
    client, err := grammes.DialWithWebSocket("ws://127.0.0.1:8182")
    if err != nil {
        log.Fatalf("Error while creating client: %s\n", err.Error())
    }

    // Executing a query to add a vertex to the graph.
    client.AddVertex("testingvertex")

    // Create a new traversal string to build your traverser.
    g := grammes.Traversal()

    // Executing a query to fetch all of the labels from the vertices.
    res, err := client.ExecuteQuery(g.V().Label())
    if err != nil {
        log.Fatalf("Querying error: %s\n", err.Error())
    }

    // Log out the response.
    for _, r := range res {
        log.Println(string(r))
    }
}
```

After this is all written you may run this file by saving it and hopping back on to your terminal. After starting your Gremlin Server and graph database in the terminal let's run this command to run the file:

```sh
go run main.go
```

For more examples look in the `examples/` directory of the project. In there you'll find multiple examples on how to use the Grammes package.

## Testing Grammes

Grammes uses [goconvey](https://www.github.com/smartystreets/goconvey/) by [smartystreets](https://www.github.com/smartystreets/) for its tests. Before trying to run the unit tests in Grammes you should update your version of this repository using this command.

```sh
go get -u github.com/smartystreets/goconvey/convey
```

Once you have this downloaded you may run the tests in Grammes how you normally would in Golang.

```sh
go test ./...
```

## Additional Resources

### Documentation on Gremlin

To learn more about how to use Gremlin I highly recommend looking through their [Tinkerpop3 documentation](http://tinkerpop.apache.org/docs/current/reference/). It's full of examples and documentation on every traversal step available.

### Examples

To find examples look in the `examples/` directory of Grammes. In there you'll find plenty of examples related to how to use this package. Make sure you're running Janusgraph before you begin running the examples.

## Troubleshooting

### Fixing time outs when starting Janusgraph

If Nodetool times out or any other part of the setup times out then the most common issue is that Cassandra is already running on your machine. To fix this run this command.

``` bash
# only for Unix at the moment.
launchctl unload ~/Library/LaunchAgents/homebrew.mxcl.cassandra.plist
```
