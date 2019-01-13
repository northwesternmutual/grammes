#!/bin/sh

JANUS_VERSION=0.3.1

STOP='stop'
STATUS='status'
LIST='list'
CLEAN='clean'
DELETE='delete'

function realpathMac() {
    [[ $1 = /* ]] && echo "$1" || echo "$PWD/${1#./}"
}

case "$OSTYPE" in
	darwin*) WORKDIR="$(realpathMac `dirname $0`)/../bin" ;;
    *) WORKDIR="$(realpath `dirname $0`)/../bin" ;;
esac

JDIR=janusgraph-${JANUS_VERSION}-hadoop2
PKGNAME=${JDIR}.zip

JPATH="${WORKDIR}/${JDIR}"

case $1 in
  $STOP)
    cd "$JPATH" && ./bin/janusgraph.sh stop
    exit 1
    ;;
  $STATUS)
    cd "$JPATH" && ./bin/janusgraph.sh status
    exit 1
    ;;
  $CLEAN)
    cd "$JPATH" && ./bin/janusgraph.sh clean
    exit 1
    ;;
  $LIST)
    echo $STOP
    echo $STATUS
    echo $CLEAN
    echo $LIST
    exit 1
    ;;
  $DELETE)
    rm -rf "$JPATH"
    exit 1
    ;;
esac

function janus_exists(){
  test -d "${JPATH}"
}

function download_janus(){
  echo Downloading janus
  mkdir -p $WORKDIR
  cd "$WORKDIR"
  case "$OSTYPE" in
  darwin*)  curl -LO https://github.com/JanusGraph/janusgraph/releases/download/v${JANUS_VERSION}/${PKGNAME} ;; 
  *)        wget -c https://github.com/JanusGraph/janusgraph/releases/download/v${JANUS_VERSION}/${PKGNAME} ;;
  esac
  echo Extracting
  unzip -o -q $PKGNAME

  # replace_janusfiles
}

function replace_janusfiles(){
  echo Replacing Janusgraph files

  cd "$JPATH" && ./bin/janusgraph.sh stop
  
  cp $WORKDIR/../assets/gremlin/gremlin-server.sh $JPATH/bin/gremlin-server.sh
  echo replaced gremlin-server.sh
  cp $WORKDIR/../assets/gremlin/janusgraph.sh $JPATH/bin/janusgraph.sh
  echo replaced janusgraph.sh
  cp $WORKDIR/../assets/gremlin/janusgraph-cql-configurationgraph.properties $JPATH/conf/janusgraph-cql-configurationgraph.properties
  echo replaced janusgraph-cql-configurationgraph.properties
  cp $WORKDIR/../assets/gremlin/gremlin-server-configuration.yaml $JPATH/conf/gremlin-server/gremlin-server-configuration.yaml
  echo replaced gremlin-server-configuration.yaml

  chmod +x $JPATH/bin/gremlin-server.sh
  chmod +x $JPATH/bin/janusgraph.sh
  chmod +x $JPATH/conf/janusgraph-cql-configurationgraph.properties
  chmod +x $JPATH/conf/gremlin-server/gremlin-server-configuration.yaml
  
  cd "$JPATH" && ./bin/janusgraph.sh clean
}


janus_exists || download_janus

echo starting janus

cd "$JPATH" && ./bin/janusgraph.sh stop
cd "$JPATH" && ./bin/janusgraph.sh start
# make sure it is not running first

echo "JanusGraph Begun!"