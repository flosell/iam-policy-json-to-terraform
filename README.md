# iam-policy-json-to-terraform

[![Build Status](https://travis-ci.org/flosell/iam-policy-json-to-terraform.svg?branch=master)](https://travis-ci.org/flosell/iam-policy-json-to-terraform)[![Release](https://img.shields.io/github/release/flosell/iam-policy-json-to-terraform.svg)](https://github.com/flosell/iam-policy-json-to-terraform/releases)

Small tool to convert an IAM Policy in JSON format into a Terraform [`aws_iam_policy_document`](https://www.terraform.io/docs/providers/aws/d/iam_policy_document.html)

## Installation

Download the latest binary from the [releases page](https://github.com/flosell/iam-policy-json-to-terraform/releases) and put it into your `PATH` under the name `iam-policy-json-to-terraform`


## Usage

```
$ cat some-policy.json | iam-policy-json-to-terraform
```
