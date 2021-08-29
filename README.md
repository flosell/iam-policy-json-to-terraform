# iam-policy-json-to-terraform
[![Build Status](https://github.com/flosell/iam-policy-json-to-terraform/actions/workflows/main.yml/badge.svg)](https://github.com/flosell/iam-policy-json-to-terraform/actions/workflows/main.yml)

Small tool to convert an IAM Policy in JSON format into a Terraform [`aws_iam_policy_document`](https://www.terraform.io/docs/providers/aws/d/iam_policy_document.html)

## Installation

### OSX

```bash
$ brew install iam-policy-json-to-terraform
```
    
### Other

Download the latest binary from the [releases page](https://github.com/flosell/iam-policy-json-to-terraform/releases) and put it into your `PATH` under the name `iam-policy-json-to-terraform`

### Developer

If you're a go developer and have your `GOPATH` defined and have added your `$GOPATH/bin` directory to your path, you can simply run this command.
```bash testcase=usage
$ go get github.com/flosell/iam-policy-json-to-terraform
```

## Usage

From raw JSON

```bash testcase=usage
$ echo '{"Statement":[{"Effect":"Allow","Action":["ec2:Describe*"],"Resource":"*"}]}' | iam-policy-json-to-terraform
data "aws_iam_policy_document" "policy" {
  statement {
    sid       = ""
    effect    = "Allow"
    resources = ["*"]
    actions   = ["ec2:Describe*"]
  }
}
```

From a JSON policy file

```bash
$ iam-policy-json-to-terraform < some-policy.json
```

## Local development

### Prerequisites

* Clone the repository to a location of your choosing: 
  ```bash testcase=building
  $ git clone git@github.com:flosell/iam-policy-json-to-terraform.git
  ```

* Install dependencies and tools: 
  ```bash testcase=building
  $ cd iam-policy-json-to-terraform
  $ make tools
  ```
  
### Development

`make` is your primary point of entry for any development activity. Call it without arguments to learn more: 

```bash testcase=building
$ make
build                          Test and build the whole application
clean                          Remove build artifacts
fmt                            Format code
fmtcheck                       Run linter
seccheck                       Run security checks
test                           Run all tests
test-readme                    Run the commands mentioned in the README for sanity-checking
tools                          Install additional required tooling
tools-main                     Install additional required tooling for the main version
tools-web                      Install additional required tooling for the web version
web-build                      Build the web version
web-deploy                     Deploy the web version to GitHub pages
web-e2e                        Run end to end tests for web version (requires web-build)
web-serve                      Serve the web version on a local development server
```
