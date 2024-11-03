.PHONY: build test clean fmt fmtcheck tools test-readme
.DEFAULT_GOAL := help
export GOFLAGS = -mod=readonly

build: test iam-policy-json-to-terraform_amd64 iam-policy-json-to-terraform_alpine iam-policy-json-to-terraform_darwin iam-policy-json-to-terraform_darwin_arm iam-policy-json-to-terraform.exe ## Test and build the whole application

clean: ## Remove build artifacts
	rm -f *_amd64 *_darwin *_alpine *.exe
	rm -rf vendor
	rm -f web/web.js*
	rm -rf web/node_modules
	rm -rf web/screenshots

test: fmtcheck seccheck **/*.go ## Run all tests
	go test -v ./...
	golint -set_exit_status ./converter
	golint -set_exit_status .
	go vet ./...

fmt: **/*.go ## Format code
	go fmt ./...

tools: tools-main tools-web ## Install additional required tooling

tools-web: ## Install additional required tooling for the web version
	test -z "${NO_TOOLS_WEB}" && (cd web && npm install) || echo "skipping tools web because of environment variable (only for testing readme)"
	test -z "${NO_TOOLS_WEB}" && cp $$(tinygo env TINYGOROOT)/targets/wasm_exec.js ./web

tools-main:  ## Install additional required tooling for the main version
	cat tools.go | grep _ | cut -f2 -d '_' | xargs -n1 go install

iam-policy-json-to-terraform_amd64: **/*.go
	 GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o $@

iam-policy-json-to-terraform_alpine: **/*.go
	 GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $@

iam-policy-json-to-terraform_darwin: **/*.go
	GOOS=darwin GOARCH=amd64 go build -o $@

iam-policy-json-to-terraform_darwin_arm: **/*.go
	GOOS=darwin GOARCH=arm64 go build -o $@

iam-policy-json-to-terraform.exe: **/*.go
	GOOS=windows GOARCH=amd64 go build -o $@

fmtcheck: **/*.go ## Run linter
	@gofmt_files=$$(gofmt -l `find . -name '*.go' | grep -v vendor`); \
    if [ -n "$${gofmt_files}" ]; then \
        echo 'gofmt needs running on the following files:'; \
        echo "$${gofmt_files}"; \
        echo "You can use the command: \`make fmt\` to reformat code."; \
        exit 1; \
    fi; \
    exit 0

seccheck: **/*.go ## Run security checks
	gosec -exclude G104 ./...

test-readme: README.md scripts/test-readme.sh ## Run the commands mentioned in the README for sanity-checking
	scripts/test-readme.sh

web-serve: web/* ## Serve the web version on a local development server
	cd web && python -m http.server 8080 && cd ..
	# TODO: rebuild tinygo on changes

web-build: web/*.go ## Build the web version using tinygo
	# TODO: ensure tinygo is installed or defer to docker?
	cd web && GOOS=js GOARCH=wasm tinygo build -o wasm.wasm ./web.go

web-e2e: web/*.go web/*.js ## Run end to end tests for web version (requires web-build)
	cd web && npm test

web-e2e-live: web/*.go web/*.js ## Run end to end tests for web version in live mode for development (requires web-build)
	cd web && npm run test-live

web-deploy: ## Deploy the web version to GitHub pages
	scripts/deploy-github-pages.sh

web-visual-regression-test:  web/*.go web/*.js web/*.css web/*.html ## Test for changes in Web UI visuals
	cd web && npx backstop test --docker

web-visual-regression-approve:  web/*.go web/*.js web/*.css web/*.html ## Accept changes in Web UI visuals
	cd web && npx backstop approve

help:
	@grep -h -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' | sort
