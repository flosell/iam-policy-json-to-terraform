# Changelog

This changelog contains a loose collection of changes in every release. I will also try and document all breaking changes to the API.

The format is based on [Keep a Changelog](http://keepachangelog.com/).

## 1.0.0 - 2018-04-28

Initial Release

### Added

* [x] Read IAM Policy JSON from STDIN and write Terraform HCL to STDOUT
* [ ] Support for the following properties
  * [x] Sid
  * [x] Effect
  * [x] Action
  * [x] NotAction
  * [x] Resource
  * [x] NotResource
  * [x] Condition with single value
