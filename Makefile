.PHONY: build test clean fmt fmtcheck tools

build: iam-policy-json-to-terraform_amd64 iam-policy-json-to-terraform_alpine iam-policy-json-to-terraform_darwin iam-policy-json-to-terraform.exe

vendor: **/*.go Gopkg.*
	dep ensure

clean:
	rm -f *_amd64 *_darwin *.exe
	rm -rf vendor

test: fmtcheck vendor **/*.go
	go test -v ./...
	golint -set_exit_status ./converter
	golint -set_exit_status .
	go vet ./...

fmt: **/*.go
	go fmt ./...

tools:
	go get -u golang.org/x/lint/golint

iam-policy-json-to-terraform_amd64: vendor *.go
	 GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o $@ *.go

iam-policy-json-to-terraform_alpine: vendor *.go
	 GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $@ *.go

iam-policy-json-to-terraform_darwin: vendor *.go
	GOOS=darwin go build -o $@ *.go

iam-policy-json-to-terraform.exe: vendor *.go
	GOOS=windows GOARCH=amd64 go build -o $@ *.go

fmtcheck: vendor **/*.go
	@gofmt_files=$$(gofmt -l `find . -name '*.go' | grep -v vendor`); \
    if [ -n "$${gofmt_files}" ]; then \
        echo 'gofmt needs running on the following files:'; \
        echo "$${gofmt_files}"; \
        echo "You can use the command: \`make fmt\` to reformat code."; \
        exit 1; \
    fi; \
    exit 0
