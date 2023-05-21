#!/bin/bash
SCRIPT_DIR=$(cd $(dirname $0) ; pwd -P)
set -e

testcases=$(cat README.md | sed -n -e 's/.*```bash testcase=\(.*\)/\1/p' | sort | uniq)

for x in ${testcases}; do
  echo "Testcase \"${x}\""
  echo

  commands=$(cat ${SCRIPT_DIR}/../README.md | sed -e 's_git@github.com:_https://github.com/_g' | sed -n "/\`\`\`bash testcase=${x}/,/\`\`\`/p" | sed -e 's/^[[:space:]]*//' | grep '^\$' | sed -e 's/^\$ //g')

  docker run -e NO_TOOLS_WEB=true --rm golang:1.17-buster bash -e -x -c "${commands}"

  echo
  echo
done
