#!/bin/sh

$(ls -dp ../bin/janusgraph* | grep "/$")bin/gremlin.sh -Q -i $PWD/../assets/gremlin/remote.groovy
