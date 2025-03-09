# iam-policy-json-to-terraform
[![Build Status](https://github.com/flosell/iam-policy-json-to-terraform/actions/workflows/main.yml/badge.svg)](https://github.com/flosell/iam-policy-json-to-terraform/actions/workflows/main.yml)

Small tool to convert an IAM Policy in JSON format into a Terraform [`aws_iam_policy_document`](https://www.terraform.io/docs/providers/aws/d/iam_policy_document.html)

## Web Version

Check out a web version of the tool [here](https://flosell.github.io/iam-policy-json-to-terraform/).

For command line usage and automation, check out the instructions below. 

## Installation

### OSX

```bash
$ brew install iam-policy-json-to-terraform
```

### [asdf](https://github.com/asdf-vm/asdf)
```bash
$ asdf plugin add iam-policy-json-to-terraform https://github.com/carlduevel/asdf-iam-policy-json-to-terraform.git
```    

### Other

Download the latest binary from the [releases page](https://github.com/flosell/iam-policy-json-to-terraform/releases) and put it into your `PATH` under the name `iam-policy-json-to-terraform`

### Developer

If you're a go developer and have your `GOPATH` defined and have added your `$GOPATH/bin` directory to your path, you can simply run this command.
```bash testcase=usage
$ go install github.com/flosell/iam-policy-json-to-terraform@latest
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

If a video demo is more your thing, checkout [this nice 2min introduction](https://www.youtube.com/watch?v=AhtpJII6eaw) by the folks at env0. 

## Local development

### Prerequisites

* Clone the repository to a location of your choosing: 
  ```bash testcase=building
  $ git clone git@github.com:flosell/iam-policy-json-to-terraform.git
  ```

* Install dependencies and tools: 
  ```bash testcase=building
  $ cd iam-policy-json-to-terraform
  $ ./go tools
  ```
  
### Development

#### Entry point 

`./go` is your primary point of entry for any development activity. Call it without arguments to learn more: 

```bash testcase=building
$ ./go
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
web-e2e-live                   Run end to end tests for web version in live mode for development (requires web-build)
web-serve                      Serve the web version on a local development server
web-visual-regression-approve  Accept changes in Web UI visuals
web-visual-regression-test     Test for changes in Web UI visuals
```

#### Web Development

To develop the web-frontend, you'll need to first compile the JavaScript version of `iam-policy-json-to-terraform`.
`./go web-build` will do that, generating a `web.js` file. 
Include it and it'll expose a `convert(policyName,jsonString)` function in the global namespace. 

Currently, the complete web-frontend is plain HTML, JS and CSS, all within `web/index.html`. 
Edit or refine as needed.

End-To-End Tests for the web frontend exist as [TestCafe](https://testcafe.io/) tests in `web_test.js` and can be run using `./go web-e2e`.
