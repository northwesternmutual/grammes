# quick-newlogger-example

The basics of setting up a custom logger and implementing it using the `quick` package.

## Description

**quick-newlogger-example** demonstrates how to create a custom logger by overriding the `PrintQuery`, `Debug`, `Error`, and `Fatal` methods. This logger is then implemented by executing the `quick.SetLogger` function.

---

## Prerequisites

- Go 1.11.1
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

- Declares a CustomLogger struct that implements all logger methods.
- Sets the new logger via the `SetLogger` method.
- Calls `DropAll` to display the print out of the new logger.

---

### Test specific steps

