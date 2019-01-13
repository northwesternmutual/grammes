<!-- STAR (situation, task, action, result) -->

# Grammes

[Grammes](http://www.github.com/northwesternmutual/grammes) is a Golang package built to communicate with a [Gremlin](http://tinkerpop.apache.org/) server via a new Client. A Gremlin server is a computing framework for graph databases such as [Janusgraph](http://janusgraph.org/).

---

<!-- The Situation -->

## Problems to Overcome

### Implementing Graph Databases in Go

Implementing a graph database into a golang workflow can be difficult when there aren't any packages or libraries to make it easy if you don't know the Gremlin language.

### The Current Golang Gremlin packages available

The current packages available to query a Gremlin server are very primitive and only allow for pulling raw data back. No extra functions to interact with the vertices on the graph itself.

---

<!-- The Task -->

## Creating This Package

### A Simple and Easy Client

To interact with the Gremlin server we've opted to create a new client that handles all the connection, graph interactions, and logging.

### Managing Requests and Responses

The client handles requests and responses using channels. This design was used to prevent halting and unnecessary waiting. This was inspired off of another Golang gremlin package go-gremlin.

---

<!-- The Action -->

## What We Did

### Building Blocks for Queries

Theoretically you can construct any kind of query using Grammes that you can using Gremlin. This is achieved by appending strings together using our custom Query type.

### Structures for Vertices, Edges, and Properties

For these types we have structures and methods related to them to interact with them directly. This effects their counterparts on the graph.

---

<!-- The Result -->

## The Finished Product

### Features of Grammes

* Handle secure connections
* Query Gremlin server
* Store vertices and edges in structures
* Query building blocks

### What Makes Grammes Unique

Grammes is a fully featured wrapper and client to interact with a graph database using a Gremlin server.

---

## Getting Started

### Github

---