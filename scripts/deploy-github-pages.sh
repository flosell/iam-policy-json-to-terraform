#!/usr/bin/env bash

set -eu -o pipefail

SCRIPT_DIR=$(dirname "$0")
REPO_ORIGIN=$(git remote get-url origin)

mkdir ${SCRIPT_DIR}/../docs
cp ${SCRIPT_DIR}/../web/{index.html,web.js} ${SCRIPT_DIR}/../docs

pushd ${SCRIPT_DIR}/../docs

git init .
git add .
git commit -m "Updating website"
git push "${REPO_ORIGIN}" master:gh-pages --force

popd

rm -rf ${SCRIPT_DIR}/../docs