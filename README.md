[![GoDoc](https://godoc.org/github.com/cf-guardian/guardian-backend?status.png)](https://godoc.org/github.com/cf-guardian/guardian-backend)

## Guardian Backend

This repository is an attempt to write the backend of Warden/Garden, entirely in Go,
and based upon the docker/libcontainer low-level container library.

It is “work-in-progress” and is not available for redistribution.

- The contents of this repository are copyright the authors.

    We don't want you copying our half-baked prototypes and then
blaming us.

    We'll open source this when we are ready, thank you.

## Objectives

1. Understandability: the code should be clearly structured and documented so that a newcomer can understand the rationale and be able to propose changes.
1. Robustness: the code should function correctly or fail with meaningful diagnostics.
1. Maintainability: it should be straightforward to fix bugs and add new features.
1. Testability: the runtime code should be thoroughly exercised by the tests.
1. Portability: it should be straightforward to port the code to other Linux distributions.

These objectives will be achieved through the following practices:

1. Construct separately testable components with documented interfaces.
1. Test each component including error paths.
1. Keep runtime and test code separate.
1. Fail rather than degrade function.
1. Operating system dependencies should be isolated and carefully managed to simplify porting.
1. Use pure Go for maintainability. Avoid scripting (even in Go) and C code.
1. Instrument code to capture failure diagnostics including error identifiers (described below) and stack traces.

## Documentation

Documentation is available at [godoc.org](http://godoc.org/github.com/cf-guardian/guardian-backend).

If this hasn't been refreshed for a while, feel free to click the "Refresh now" link.

## Diagnostics

### Errors

Errors returned from Guardian backend functions and methods include error identifiers to uniquely identify each kind of failure and stack traces so that the point of failure can easily be determined. The [gerror](gerror) package is used to construct errors.

*This principle is an aim, not an achieved goal.*

### Logging

Logging is performed using the [glog](https://github.com/golang/glog) package (an external dependency).

`glog` has four logging levels: informational (`glog.Info()`), warning (`glog.Warning()`), error (`glog.Error()`), and fatal (`glog.Fatal()`). Additionally, `glog` has a verbosity level  which is a non-negative integer with a default value of `0`. Guardian uses informational logs with verbosity level `0` to log important events and informational logs with verbosity levels `1` and higher for debugging.

Logs may be directed to standard error by setting the flag `logtostderr` to `true` on the go invocation, as in this example:

```
go test -logtostderr=true -vmodule=*=2
```

Note: the glog `-v` flag clashes with the boolean `-v` flag of `go test` and so the logging verbosity should be set during testing using `-vmodule=*=`.

See the [glog documentation](http://godoc.org/github.com/golang/glog) for further information.

## Repository Layout

The directories in this repository consist mostly of runtime code, tests (`*_test.go`), and examples (`*example*.go`). Directories named `mock_*` contain mock implementations of interfaces for use in testing.

In addition, the `development` directory contains scripts used during development and the `test_support` directory contains shared functions and custom matchers used by tests.

## Build constraints

There are a few places where we may implement only Linux versions of (usually low-level) functions. These are in
`*_linux.go` files. The aim is to develop code that is (at least) testable on any platform, which may mean implementing linux and non-linux versions of some low-level code.

Every such file will also record its constraints in build constraints in the file content. For example, even if a file is called `fred_linux.go` it will still contain this build constraint:

```
// +build linux
```

The rule to follow is: the build constriants are definitive and complete; the filenames are used to indicate the constraints. (Note: the build constraint `// +build !linux` must *not* be indicated by a file name ending in `_non_linux.go`. This will not work.)

## Development Environment Setup

1. Ensure the following pre-requisites are installed:
    * [git](http://git-scm.com/downloads)
    * [Go](http://golang.org/) 1.2.1 or later:

        - either [download](http://golang.org/doc/install) a specific version
        - or use [gvm](https://github.com/moovweb/gvm)
        - or even `port install go` with [MacPorts](http://www.macports.org/).

2. Create a Go [workspace](http://golang.org/doc/code.html#Organization) directory, such as `$HOME/go`, and add the path of this directory to the
beginning of a new environment variable called `GOPATH`. You might want to put this last step in your profile.

    ```
    $ mkdir $HOME/go
    $ export GOPATH=$HOME/go
    ```

3. Get this repository into your workspace (`src` directory) by issuing:

    ```
    $ go get github.com/cf-guardian/guardian
    ```

4. Change directory to `<workspace dir>/src/github.com/cf-guardian/guardian`.

5. Install the [pre-commit hook](https://github.com/jbrukh/git-gofmt) as follows:

    ```
    cd .git/hooks
    ln -s ../../development/pre-commit-hook/pre-commit .
    ```

    After installing the hook, if you need to skip reformatting for a particular commit, use `git commit --no-verify`.

### Development scripts

The directory `development/scripts` is there for simple (bash) shell scripts that may make life a little easier in our context:

* `gosub` is designed to run a `go` command in all subdirectories of the current directory that have a
  `*.go` file in them. Typical usage is `gosub build`, or `gosub fmt build`. This saves being driven by a
  top-level `go` package with explicit dependencies.

    `gosub` is limited to single word go commands. `gosub fmt build` would issue `go fmt; go build` in each
    directory in turn. `gosub test` is quite useful to run all tests in immediate subdirectories.

* `govet` runs `go tool vet` with appropriate options against a single directory passed as a parameter to the
  script.

* `goroot` is a simple script suitable for running the `go` tool as root. It simply sources the
  `$HOME/.go-setup` file, which sets up the Go environment vars, and then executes the go tool with the
  parms provided. This is necessary since root doesn't normally have a go tool environment in its shell.

## Testing

To run the tests in a specific directory, issue:

```
go test
```

If the tests succeed, this should print `PASS`.

### gomock

Unit testing is performed on packages. [gomock](http://godoc.org/code.google.com/p/gomock/gomock) is
used as a mocking framework and to generate mock implementations of interfaces and thereby enable
packages to be tested in isolation from each other.

( Once you have installed Go, run these commands to install the gomock package and the mockgen tool:

```
go install code.google.com/p/gomock/gomock
go install code.google.com/p/gomock/mockgen
```

See [gomock README](https://code.google.com/p/gomock/source/browse/README.md) for the rest of the instructions. )

Mocks are stored in a subdirectory of the directory containing the mocked interface.
For example, the mocks for the `kernel/fileutils` package are stored in `kernel/fileutils/mock_fileutils`.

To re-generate the generated mock implementations of interfaces, ensure `mockgen` is on the path and then run the script `development/scripts/update_mocks`. If you add a generated mock, don't forget to add it to `update_mocks`.

## Editing

If your favourite text editor is not sufficient, try [Eclipse](http://www.eclipse.org/downloads/) with the [goclipse plugin](https://github.com/sesteel/goclipse) or [IntelliJ IDEA](http://www.jetbrains.com/idea/) with the [go plugin](https://github.com/go-lang-plugin-org/go-lang-idea-plugin).

Source code is formatted according to standard Go conventions. To re-format the code, and all its
dependencies issue:

```
go fmt ./...
```

To reformat code automatically whenever it is committed to git, install the pre-commit hook as described above.

Also, you can run `go vet` (a `govet` script is provided in `development/scripts`) and
[lint](http://go-lint.appspot.com/github.com/cf-guardian/guardian) against the code if you like.

## Contributing

[Pull requests](http://help.github.com/send-pull-requests) are welcome; see the
[contributor guidelines](CONTRIBUTING.md) for details.
