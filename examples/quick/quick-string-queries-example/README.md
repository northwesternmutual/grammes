# quick-string-queries-example

The basics of executing custom queries using the `ExecuteStringQuery` method in the `quick` package.

## Description

**quick-string-queries-example** demonstrates how to execute custom queries using the `quick` package. This example first drops all existing vertices, adds a new vertex, and queries the graph for the label of the newly added vertex.

---

## Prerequisites

- go 1.12
- Git
- Elastic Search
- Cassandra
  - Java 8

---

## Running

To run this test you will need a TinkerPop server running and a graph database to connect to locally. This example was tested while using JanusGraph which can be used by locating yourself to the root directory of the Grammes project.

``` sh
cd $GOPATH/src/github.com/northwesternmutual/grammes
```

After locating yourself here then you may change directory to the `/scripts` folder.

``` sh
cd scripts
```

Finally you may run the `janusgraph.sh` script to begin a local instance of JanusGraph. This will begin the TinkerPop server for you as well.

``` sh
./janusgraph.sh
```

For further instructions please find yourself to the root [README.md](../../README.md) file.

---

## Steps

### General steps

- Create a [zap](https://github.com/uber-go/zap) logger to help explain what's going on in the test and display the results.
- Adds a test vertex to the graph using `ExecuteStringQuery` found in the `quick` package.
  - For testing this was created using JanusGraph. This can be run in the `/scripts` directory.
- Adds a label to the vertex that we have created.
- Recovers the label by running the query `g.V().label()` via the `ExecuteStringQuery` method.
- Drop all of the possible interfering vertices that were already on the graph.
- Defer a drop of all the testing vertices. This is done as clean up.

---

### Test specific steps

- Shows how to execute string queries without a client
