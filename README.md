# Go-utils

[![CircleCI](https://circleci.com/gh/medtune/go-utils/tree/master.svg?style=svg)](https://circleci.com/gh/medtune/go-utils/tree/master) [![codecov](https://codecov.io/gh/medtune/beta-platform/branch/master/graph/badge.svg)](https://codecov.io/gh/medtune/beta-platform) [![CodeFactor](https://www.codefactor.io/repository/github/medtune/go-utils/badge)](https://www.codefactor.io/repository/github/medtune/go-utils)


A set of Go libraries that provide low-level, independent packages supplementing the Go standard libs.


## Purpose

This repository is intended to hold shared utilities
with no Medtune dependence that may be of interest
to any Go project.


## Criteria for adding code here

- Used by multiple repositories.

- Full unit test coverage.

- Go tools compliant (`go get`, `go test`, etc.).

- Complex enough to be worth vendoring, rather than copying.

- Stable, or backward compatible, API.

## Libraries


- [Errors](/errors) implementation to fit all use cases.

- [Math](/math) completing go math standard library.

- [Exec](/exec) provides an interface for `os/exec`. It makes it easier
  to mock and replace in tests, especially with
  the [FakeExec](exec/testing/fake_exec.go) struct.

- [Cert](/cert) provide base functionnalities for creating certificates.

- [Temp](/temp) provides an interface to create temporary directories. It also
  provides a [FakeDir](temp/temptest) implementation to replace in tests.

- [Clock](/clock) provides an interface for time-based operations.  It allows
  mocking time for testing.
  
- [Pointers](/pointers) provides some functions for pointer-based operations.

- [Crypto](/crypto) expose simple API for hash functionalities.

- [Random](/random) provide random generation mechanics for different types.

- [Regex](/regex) provide predefined regex matching patterns.

