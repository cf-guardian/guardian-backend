[![GoDoc](https://godoc.org/github.com/cf-guardian/guardian-backend?status.svg)](https://godoc.org/github.com/cf-guardian/guardian-backend)

# Guardian Backend

This repository is an attempt to write the backend of Warden/Garden, entirely in Go,
based upon the docker/libcontainer low-level container library.

It is “work-in-progress” and is not available for redistribution.

- The contents of this repository are copyright the authors.

    We don't want you copying our half-baked prototypes and then blaming us.

    We'll open source this when we are good and ready, thank you.

In this document we outline:

* [Objectives](#objectives) of the development
* where to get the [Documentation](#documentation), and
* how the repository is [layed out](#repository-layout).

We refer you to other documents for:

* how to [contribute](CONTRIBUTING.md) to this repository
* how to set up your environment to [develop](DEVELOPMENT.md) code in this repository, and
* what [coding conventions](CONVENTIONS.md) we are adopting.

## Objectives

1. *Understandability*: the code should be clearly structured and documented so that a newcomer can understand the rationale and be able to propose changes.
1. *Robustness*: the code should function correctly or fail with meaningful diagnostics.
1. *Maintainability*: it should be straightforward to fix bugs and add new features.
1. *Testability*: the runtime code should be thoroughly exercised by the tests.
1. *Portability*: it should be straightforward to port the code to other Linux distributions.

These objectives will be achieved through the following practices:

1. Construct separately testable components with documented interfaces, injecting stubbed or mocked dependencies for testing.
1. Test each component including error paths.
1. Separate runtime and test code.
1. Code to fail rather than degrade function.
1. Preserve integrity even when returning errors by the use of transactional (atomic) programming.
1. Isolate operating system dependencies to simplify porting (and testing).
1. Use pure Go for maintainability. Avoid scripting and C code ***whenever possible***.
1. Capture failure diagnostics including typed error identifiers and stack traces.

These practices (and some compromises) are described in the [coding
conventions](CONVENTIONS.md) document.

## Documentation

Automatically generated documentation, taken from the source code, is available at
[godoc.org](http://godoc.org/github.com/cf-guardian/guardian-backend) (also
linked to by the `godoc|reference` icon at the top of this README).

If this hasn't been refreshed for a while, feel free to click the "Refresh now" link.

## Repository layout

The directories in this repository consist mostly of runtime code, tests (`*_test.go`), and examples (`*example*.go`). Directories named `mock_*` contain mock implementations of interfaces for use in testing.

In addition, the `development` directory contains scripts used during development and the `test_support` directory contains shared functions and custom matchers used by tests.

For the design and coding conventions used in this repository, see
[Conventions](CONVENTIONS.md).

## Contributing

[Pull requests](http://help.github.com/send-pull-requests) are welcome; see the
[contributor guidelines](CONTRIBUTING.md) for details.
