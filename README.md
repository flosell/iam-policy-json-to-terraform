# iam-policy-json-to-terraform

Small tool to convert an IAM Policy in JSON format into a Terraform [`aws_iam_policy_document`](https://www.terraform.io/docs/providers/aws/d/iam_policy_document.html)

## Installation

TODO

## Usage

```
$ cat some-policy.json | iam-policy-json-to-terraform
```

## TODO:

* [x] Read JSON
* [x] Format HCL
* [ ] Support all attributes of the datasource
  * [x] Sid
  * [x] Effect
  * [ ] Principal
  * [ ] NotPrincipal
  * [x] Action
  * [x] NotAction
  * [x] Resource
  * [x] NotResource
  * [x] Condition with single value
  * [ ] Condition with multiple values (e.g. [s3:prefix](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_statement.html))
* [ ] Make data source name configurable
* [ ] `./go`-script or `Makefile`
* [ ] Release (and complete installation instructions)
