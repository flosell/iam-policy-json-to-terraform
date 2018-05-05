# iam-policy-json-to-terraform

[![Build Status](https://travis-ci.org/flosell/iam-policy-json-to-terraform.svg?branch=master)](https://travis-ci.org/flosell/iam-policy-json-to-terraform)[![Release](https://img.shields.io/github/release/flosell/iam-policy-json-to-terraform.svg)](https://github.com/flosell/iam-policy-json-to-terraform/releases)

Small tool to convert an IAM Policy in JSON format into a Terraform [`aws_iam_policy_document`](https://www.terraform.io/docs/providers/aws/d/iam_policy_document.html)

## Installation

Download the latest binary from the [releases page](https://github.com/flosell/iam-policy-json-to-terraform/releases) and put it into your `PATH` under the name `iam-policy-json-to-terraform`


## Usage

```
$ cat some-policy.json | iam-policy-json-to-terraform
```

## TODO:

* [x] Read JSON
* [x] Format HCL
* [x] Support all attributes of the datasource
  * [x] Sid
  * [x] Effect
  * [x] Principal
  * [x] NotPrincipal
  * [x] Action
  * [x] NotAction
  * [x] Resource
  * [x] NotResource
  * [x] Condition with single value
  * [x] Condition with multiple values (e.g. [s3:prefix](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_statement.html))
* [ ] Make data source name configurable
* [x] `./go`-script or `Makefile`
* [x] Release (and complete installation instructions)
