#!/usr/bin/env bash

HOST=localhost
PORT=8080




function index() {
  curl -X GET  \
    "http://$HOST:$PORT/" \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json'

}

function create() {
  TITLE=$1
  DATA=$2
  curl -X POST  \
    "http://$HOST:$PORT/objects" \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json' \
    -d '{
      "title": "'$TITLE'",
      "data": ["1", "2", "3"]
    }'
}

function get() {
  ID=$1
  curl -X GET  \
    "http://$HOST:$PORT/objects/$ID" \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json'
}


function delete() {
  ID=$1
  curl -X DELETE  \
    "http://$HOST:$PORT/objects/$ID" \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json'
}

function update() {
  ID=$1
  TITLE=$2
  #DATA=[1, 2, 3]
  curl -X PUT  \
    "http://$HOST:$PORT/objects" \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json' \
    -d '{
       "_id": "'$ID'",
      "title": "'$TITLE'",
      "data": ["1111", "2222", "3333"]
    }'
}