# Changelog

This changelog contains a loose collection of changes in every release. I will also try and document all breaking changes to the API.

The format is based on [Keep a Changelog](http://keepachangelog.com/).

## Unreleased

### Changed

* Minor dependency updates

### Fixed

* Allow multiple Terraform variables in the same line (#59)
* Include more helpful error messages on parsing errors (#49)

## 1.8.0 - 2021-07-07

### Added

* Usage message if STDIN is a terminal (#18)

## 1.7.0 - 2021-05-03

### Added

* Support for Mac ARM CPUs

## 1.6.0 - 2020-12-22

### Added

* Support for input that contains HCL expressions that break JSON syntax (#16) - this should make it easier to convert from heredoc to terraform

## 1.5.0 - 2020-05-24

### Fixed

* Limit escaping of dollar signs to IAM policy variables, don't escape terraform interpolations (#13)

## 1.4.0 - 2020-04-19

### Fixed

* Parsing of single-statement policies when expressed as a JSON object instead of a JSON object wrapped into an array (#10)

## 1.3.0 - 2020-01-26

### Added

* Add `-version` flag that returns the current version. Thanks @nitrocode for the contribution!

## 1.2.0 - 2019-03-02

### Added

* Support for wildcard-principal `"Principal": "*"` (#2)

### Fixed

* Booleans in JSON are converted to empty-string instead of their real string representation (#2) 

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
