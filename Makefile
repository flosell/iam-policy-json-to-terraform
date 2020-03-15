.PHONY: build test clean fmt fmtcheck tools test-readme
.DEFAULT_GOAL := help
export GOFLAGS = -mod=readonly

build: test iam-policy-json-to-terraform_amd64 iam-policy-json-to-terraform_alpine iam-policy-json-to-terraform_darwin iam-policy-json-to-terraform.exe ## Test and build the whole application

vendor: **/*.go go.* ## Install dependencies into ./vendor
	go mod vendor

clean: ## Remove build artifacts and vendored dependencies
	rm -f *_amd64 *_darwin *.exe
	rm -rf vendor

test: fmtcheck vendor **/*.go ## Run all tests
	go test -v ./...
	golint -set_exit_status ./converter
	golint -set_exit_status .
	go vet ./...

fmt: **/*.go ## Format code
	go fmt ./...

tools: ## Install additional required tooling
	go install golang.org/x/lint/golint

iam-policy-json-to-terraform_amd64: vendor **/*.go
	 GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o $@ *.go

iam-policy-json-to-terraform_alpine: vendor **/*.go
	 GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $@ *.go

iam-policy-json-to-terraform_darwin: vendor **/*.go
	GOOS=darwin go build -o $@ *.go

iam-policy-json-to-terraform.exe: vendor **/*.go
	GOOS=windows GOARCH=amd64 go build -o $@ *.go

fmtcheck: vendor **/*.go ## Run linter
	@gofmt_files=$$(gofmt -l `find . -name '*.go' | grep -v vendor`); \
    if [ -n "$${gofmt_files}" ]; then \
        echo 'gofmt needs running on the following files:'; \
        echo "$${gofmt_files}"; \
        echo "You can use the command: \`make fmt\` to reformat code."; \
        exit 1; \
    fi; \
    exit 0

test-readme: README.md scripts/test-readme.sh ## Run the commands mentioned in the README for sanity-checking
	scripts/test-readme.sh

help:
	@grep -h -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
