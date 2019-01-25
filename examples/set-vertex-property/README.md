# set-vertex-property

The basics of setting vertex properties using the `SetVertexProperty` and `AddProperty` function.

## Description

**set-vertex-property** demonstrates how to set or add vertex properties in two ways. First way is using the Grammes client with `SetVertexProperty` and the second is with a `Vertex` struct and using the `AddProperty` function. They both do the same thing, but you can use either when convenient. They either ADD or OVERWRITE the current value of the properties of the vertex.

## Prerequisites

- Go 1.11.1
- Git
- Elastic Search
- Cassandra
  - Java 8

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

## Steps

### General steps

- Create a [zap](https://github.com/uber-go/zap) logger to help explain what's going on in the test and display the results.
- Creates a Grammes client that connects to a locally hosted [TinkerPop](http://tinkerpop.apache.org/) server with a WebSocket.
  - For testing this was created using JanusGraph. This can be run in the `/scripts` directory.
- Drop all of the possible interfering vertices that were already on the graph.
- Defer a drop of all the testing vertices. This is done as clean up.

---

### Test specific steps

- Adds testing vertices
- Sets/Adds properties to the vertices after they're created
  - Using the GraphManager
  - Using the Vertex struct
