#!/usr/bin/env bash

set -eu -o pipefail

SCRIPT_DIR=$(cd $(dirname $0) ; pwd -P)
REPO_ORIGIN=$(git remote get-url origin)
ASSEMBLY_DIR="${SCRIPT_DIR}/../docs"

mkdir "${ASSEMBLY_DIR}"
cp ${SCRIPT_DIR}/../web/{index.html,app.css,app.js,wasm.wasm,wasm_exec.js,version.txt} "${ASSEMBLY_DIR}"

cp -r ${SCRIPT_DIR}/../web/img "${ASSEMBLY_DIR}"

pushd "${ASSEMBLY_DIR}"

git init .

cp ${SCRIPT_DIR}/../.git/config .git/config

git add .
git commit -m "Updating website"
git push "${REPO_ORIGIN}" master:gh-pages --force

popd
