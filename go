#!/bin/bash
set -e

SCRIPT_DIR=$(cd $(dirname $0) ; pwd -P)

die() {
  red=$(tput setaf 1)
  normal=$(tput sgr0)

  echo "${red}${1}${normal}"
  exit 1
}

goal_build() { ## build all variants
  goal_iam_policy_json_to_terraform_amd64
  goal_iam_policy_json_to_terraform_alpine
  goal_iam_policy_json_to_terraform_darwin
  goal_iam_policy_json_to_terraform_darwin_arm
  goal_iam_policy_json_to_terraform_exe
  # TODO: should this also build wasm?
}

goal_clean() { ## Remove build and test artifacts as well as dependencies
  rm -f -- *_amd64 *_darwin *_alpine *.exe
  rm -rf vendor
  rm -f web/web.js*
  rm -rf web/node_modules
  rm -rf web/screenshots
}

goal_test() { ## Run all tests
  go test -v ./...
}

goal_check() { ## Check code style and common bug patterns
  goal_check_format
  goal_check_style
  goal_check_security
}

goal_fmt() { ## Format code
  go fmt ./...
}

goal_tools() { ## Install additional required tooling
  goal_tools_main
  goal_tools_web
}

goal_tools_web() { ## Install additional required tooling for the web version
  test -z "${NO_TOOLS_WEB}" && (cd web && npm install) || echo "skipping tools web because of environment variable (only for testing readme)"
}

goal_tools_main() { ## Install additional required tooling for the main version
  cat tools.go | grep _ | cut -f2 -d '_' | xargs -n1 go install
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
  golint -set_exit_status ./converter
  golint -set_exit_status .
  go vet ./...
}

goal_check_security() { ## Run security checks
  gosec -exclude G104 ./...
}

goal_test_readme() { ## Run the commands mentioned in the README for sanity-checking
  scripts/test-readme.sh
}

goal_web_serve() { ## Serve the web version on a local development server
  cd web && gopherjs serve github.com/flosell/iam-policy-json-to-terraform/web/
}

goal_web_build() { ## Build the web version
  cd web && gopherjs build --minify
}

goal_web_e2e() { ## Run end to end tests for web version (requires web-build)
  cd web && npm test
}

goal_web_e2e_live() { ## Run end to end tests for web version in live mode for development (requires web-build)
  cd web && npm run test-live
}

goal_web_deploy() { ## Deploy the web version to GitHub pages
  scripts/deploy-github-pages.sh
}

goal_web_visual_regression_test() { ## Test for changes in Web UI visuals
  cd web && npx backstop test --docker
}

goal_web_visual_regression_approve() { ## Accept changes in Web UI visuals
  cd web && npx backstop approve
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
