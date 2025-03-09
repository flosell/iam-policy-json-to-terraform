.PHONY: build test clean fmt fmtcheck tools test-readme
.DEFAULT_GOAL := help
export GOFLAGS = -mod=readonly

build:
	./go build

clean: ## Remove build artifacts
	./go clean

test:  ## Run all tests
	./go test

fmt: **/*.go ## Format code
	./go fmt

tools: tools-main tools-web ## Install additional required tooling

tools-web: ## Install additional required tooling for the web version
	./go tools-web
tools-main:  ## Install additional required tooling for the main version
	./go tools-main

iam-policy-json-to-terraform_amd64: **/*.go
	 ./go iam_policy_json_to_terraform_amd64

iam-policy-json-to-terraform_alpine: **/*.go
	 ./go iam_policy_json_to_terraform_alpine

iam-policy-json-to-terraform_darwin: **/*.go
	./go iam_policy_json_to_terraform_darwin

iam-policy-json-to-terraform_darwin_arm: **/*.go
	./go iam_policy_json_to_terraform_darwin_arm

iam-policy-json-to-terraform.exe: **/*.go
	../go iam_policy_json_to_terraform_exe

fmtcheck: **/*.go ## Run linter
	./go fmtcheck

seccheck: **/*.go ## Run security checks
	./go seccheck

test-readme: README.md scripts/test-readme.sh ## Run the commands mentioned in the README for sanity-checking
	./go test_readme

web-serve: web/* ## Serve the web version on a local development server
	./go web_serve
web-build: web/*.go ## Build the web version
	./go web_build
web-e2e: web/*.go web/*.js ## Run end to end tests for web version (requires web-build)
	./go web_e2e

web-e2e-live: web/*.go web/*.js ## Run end to end tests for web version in live mode for development (requires web-build)
	./go web_e2e_live

web-deploy: ## Deploy the web version to GitHub pages
	./go web_deploy

web-visual-regression-test:  web/*.go web/*.js web/*.css web/*.html ## Test for changes in Web UI visuals
	./go web_visual_regression_test

web-visual-regression-approve:  web/*.go web/*.js web/*.css web/*.html ## Accept changes in Web UI visuals
	./go web_visual_regression_approve

help:
	@grep -h -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' | sort
