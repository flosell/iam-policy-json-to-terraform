#!/usr/bin/env bash

set -eu -o pipefail

SCRIPT_DIR=$(cd $(dirname $0) ; pwd -P)
REPO_ORIGIN=$(git remote get-url origin)

mkdir ${SCRIPT_DIR}/../docs
cp ${SCRIPT_DIR}/../web/{index.html,web.js} ${SCRIPT_DIR}/../docs

pushd ${SCRIPT_DIR}/../docs

git init .

cp ${SCRIPT_DIR}/../.git/config .git/config

git add .
git commit -m "Updating website"
git push "${REPO_ORIGIN}" master:gh-pages --force

popd

rm -rf ${SCRIPT_DIR}/../docs