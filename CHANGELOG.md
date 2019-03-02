# Changelog

This changelog contains a loose collection of changes in every release. I will also try and document all breaking changes to the API.

The format is based on [Keep a Changelog](http://keepachangelog.com/).

## Unreleased

### Added

* Support for wildcard-principal `"Principal": "*"` (#2)

## 1.1.0 - 2018-05-05

### Added 

* Support for `Condition` with multiple values
* Support for `Pricipal` and `NotPrincipal`
* Flag `-name` to specify the name of the generated policy

## 1.0.0 - 2018-04-28

Initial Release

### Added

* Read IAM Policy JSON from STDIN and write Terraform HCL to STDOUT
* Support for the following properties
  * `Sid`
  * `Effect`
  * `Action`
  * `NotAction`
  * `Resource`
  * `NotResource`
  * `Condition` with single value
