#!/bin/bash
set -e

SCRIPT_DIR=$(cd $(dirname $0) ; pwd -P)

background_pids=()
cleanup() {
  for pid in "${background_pids[@]}"; do
    if ps "${pid}" &>/dev/null; then
      kill -- "${pid}"
    fi
  done
}
trap cleanup EXIT

die() {
  red=$(tput setaf 1)
  normal=$(tput sgr0)

  echo "${red}${1}${normal}"
  exit 1
}

goal_cli_build() { ## build all CLI variants
  goal_iam_policy_json_to_terraform_amd64
  goal_iam_policy_json_to_terraform_alpine
  goal_iam_policy_json_to_terraform_darwin
  goal_iam_policy_json_to_terraform_darwin_arm
  goal_iam_policy_json_to_terraform_exe
}

goal_clean() { ## Remove build and test artifacts as well as dependencies
  rm -f -- *_amd64 *_darwin *_alpine *.exe
  rm -rf vendor
  rm -f web/web.js*
  rm -rf web/node_modules
  rm -rf web/screenshots
  rm -rf .bin
}

goal_test() { ## Run all tests
  go test -v ./converter/
}

goal_check() { ## Check code style and common bug patterns
  goal_check_format
  goal_check_style
  goal_check_security
}

goal_fmt() { ## Format code
  go fmt ./...
}

goal_tools_tinygo() {
  TINYGO_VERSION="0.36.0" # cross-reference this with the referenced version in the duckdb-wasm release used by evidence (see its package-lock.json)

  if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    OS="linux"
  else
    OS="darwin"
  fi

  if [[ "$(uname -m)" == "x86_64" ]]; then
    ARCH="amd64"
  else
    ARCH="arm64"
  fi

  mkdir -p "${SCRIPT_DIR}/.bin"

  curl --output "${SCRIPT_DIR}/.bin/tinygo.tar.gz" -L "https://github.com/tinygo-org/tinygo/releases/download/v${TINYGO_VERSION}/tinygo${TINYGO_VERSION}.${OS}-${ARCH}.tar.gz"

  cd ${SCRIPT_DIR}/.bin
  tar xzf tinygo.tar.gz
  cd ..
}

goal_tools_web() { ## Install additional required tooling for the web version
  if [ -z "${NO_TOOLS_WEB}" ]; then
    cd web && npm install; cd ..
    goal_tools_tinygo
    cat $(.bin/tinygo/bin/tinygo env TINYGOROOT)/targets/wasm_exec.js > web/wasm_exec.js
  else
    echo "skipping tools web because of environment variable (only for testing readme)"
  fi
}


goal_iam_policy_json_to_terraform_amd64() {
  GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o iam-policy-json-to-terraform_amd64
}

goal_iam_policy_json_to_terraform_alpine() {
  GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o iam-policy-json-to-terraform_alpine
}

goal_iam_policy_json_to_terraform_darwin() {
  GOOS=darwin GOARCH=amd64 go build -o iam-policy-json-to-terraform_darwin
}

goal_iam_policy_json_to_terraform_darwin_arm() {
  GOOS=darwin GOARCH=arm64 go build -o iam-policy-json-to-terraform_darwin_arm
}

goal_iam_policy_json_to_terraform_exe() {
  GOOS=windows GOARCH=amd64 go build -o iam-policy-json-to-terraform.exe
}

goal_check_format() { ## Run linter
  gofmt_files=$(gofmt -l $(find . -name '*.go' | grep -v vendor))
  if [ -n "${gofmt_files}" ]; then
    echo 'gofmt needs running on the following files:'
    echo "${gofmt_files}"
    echo "You can use the command: \`$0 fmt\` to reformat code."
    exit 1
  fi
}

goal_check_style() { ## Check code style
  go tool revive -config revive.toml -set_exit_status ./converter
  go tool revive -config revive.toml -set_exit_status .
  go vet ./converter
  go vet .
}

goal_check_security() { ## Run security checks
  go tool gosec -exclude G104 ./converter
  go tool gosec -exclude G104 .
}

goal_test_readme() { ## Run the commands mentioned in the README for sanity-checking
  scripts/test-readme.sh
}

web_serve_background() {
  cd web
  python -m http.server  --bind 0.0.0.0 8080 &
  background_pids+=("$!")
  cd ..
}

goal_web_serve() { ## Serve the web version on a local development server
  web_serve_background
  wait
}

goal_web_build() { ## Build the web version
  ${SCRIPT_DIR}/.bin/tinygo/bin/tinygo build -o web/wasm.wasm -target=wasm web/web.go
}

goal_web_build_watch() {
  fswatch -o web/web.go | xargs -n1 -I{} bash -c 'echo "rebuilding" && ./go web_build'
}

goal_web_e2e() { ## Run end to end tests for web version (requires web-build)
  web_serve_background
  cd web
  TARGET_URL="${TARGET_URL:-localhost:8080/}" npm test
  cd ..
}

goal_wait_for_deployed() {
  TARGET_URL="https://flosell.github.io/iam-policy-json-to-terraform/version.txt"

  expected_version="$(cat docs/version.txt)" # this should come from artifact
  while true; do
    current_version=$(curl -s "${TARGET_URL}")
    echo "Current version: '${current_version}', expected version: '${expected_version}'"

    if [ "${current_version}" != "${expected_version}" ]; then
      sleep 1
    else
      return 0
    fi
  done
}

goal_web_e2e_deployed() { ## Run end to end tests against the deployed web version
  cd web
  TARGET_URL="${TARGET_URL:-https://flosell.github.io/iam-policy-json-to-terraform/}" npm test
  cd ..
}

goal_web_e2e_live() { ## Run end to end tests for web version in live mode for development (requires web-build)
  web_serve_background
  cd web
  TARGET_URL="${TARGET_URL:-localhost:8080/}" npm run test-live
  cd ..
}

goal_web_deploy() { ## Deploy the web version to GitHub pages
  scripts/deploy-github-pages.sh
}

goal_web_visual_regression_test() { ## Test for changes in Web UI visuals
  web_serve_background
  # docker host networking is different locally and on GitHub actions
  if [ -z "$CI" ]; then
    export TARGET_URL="http://host.docker.internal:8080/"
  else
    export TARGET_URL="http://172.17.0.1:8080/"
  fi
  cd web && npx backstop test --config=backstop.js --docker
}

goal_web_visual_regression_approve() { ## Accept changes in Web UI visuals
  web_serve_background
  export TARGET_URL="http://host.docker.internal:8080/" # not planning to run this on CI so only local option
  cd web && npx backstop --config=backstop.js approve
}

goal_help() { ## this help message
  bold=$(tput bold)
  normal=$(tput sgr0)
  echo "usage: ${bold}$0 <goal>${normal}
goals:"
  cat "$0" | sed -nr -e "s/goal_([a-zA-Z0-9_-]+).*#(.*)# *(.*)/    ${bold}\1${normal} \2 |--- \3/p" | column -ts '|' | sort
}

if type -t "goal_$1" &>/dev/null; then
  goal_$1 "${@:2}"
else
  goal_help
  exit 1
fi
