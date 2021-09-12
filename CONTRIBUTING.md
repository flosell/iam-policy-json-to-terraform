# Contribution Guide

Thanks for wanting to help out! This project only succeeds because people like you keep on improving it.

This guide is trying to get you started. It contains helpful guidelines, not hard rules.
If you aren't sure about something, just ask!

## I want to help but I don't know what to do

Have a look at the [issue tracker](https://github.com/flosell/iam-policy-json-to-terraform/issues), maybe you'll find something interesting there.
The [help wanted](https://github.com/flosell/iam-policy-json-to-terraform/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22) and
[good first issue](https://github.com/flosell/iam-policy-json-to-terraform/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22)
labels track issues that make good starting points.

## How to open the perfect issue

* Be specific and as detailed as you feel is necessary to understand the topic
* Provide context (what were you trying to achive, what were you expecting, ...)
* Code samples and logs can be really helpful. Consider [Gists](https://gist.github.com/) or links to other Github repos
  for larger pieces.
* If the UI behaves in a strange way, have a look at your browsers development tools. The console and network traffic might give you some insight that's valuable for a bug report.
* If you are reporting a bug, add steps to reproduce it.

## How to create the perfect pull request

* Have a look into the [`README`](README.md#local-development) for details on how to work with the code
* Follow the usual best practices for pull requests:
    * use a branch,
    * make sure you have pulled changes from upstream so that your change is easy to merge
    * follow the conventions in the code
    * keep a tidy commit history that speaks for itself, consider squashing commits where appropriate
* Run all the tests: `make test`
* Add tests where possible
* Add an entry in [`CHANGELOG.md`](CHANGELOG.md) if you add new features, fix bugs or otherwise change things in a way that you want
  users to be aware of. The entry goes into the `Unreleased` section, usually the top one. If that section doesn't exist yet, add it. 
