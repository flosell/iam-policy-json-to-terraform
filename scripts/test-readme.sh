#!/bin/bash
SCRIPT_DIR=$(cd $(dirname $0) ; pwd -P)

commands=$(cat ${SCRIPT_DIR}/../README.md | sed -e 's_git@github.com:_https://github.com/_g' | sed -n '/```bash testcase=.*/,/```/p' | sed -e 's/^[[:space:]]*//' | grep '^\$' | sed -e 's/^\$ //g')

docker run --rm golang:1-buster bash -e -x -c "${commands}"
