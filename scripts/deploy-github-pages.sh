#!/usr/bin/env bash

set -eu -o pipefail

SCRIPT_DIR=$(cd $(dirname $0) ; pwd -P)
ASSEMBLY_DIR="$1"

if [ -z "${ASSEMBLY_DIR}" ]; then
  echo "Usage: $0 <directory with github pages branch checked out>"
  exit 1
fi

cp ${SCRIPT_DIR}/../web/{index.html,app.css,app.js,wasm.wasm,wasm_exec.js,version.txt} "${ASSEMBLY_DIR}"
cp -r ${SCRIPT_DIR}/../web/img "${ASSEMBLY_DIR}"

pushd "${ASSEMBLY_DIR}"

git add .
git commit -m "Updating website"
git push origin gh-pages

popd
