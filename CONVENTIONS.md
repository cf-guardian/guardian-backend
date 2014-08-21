# Conventions used in this repository

The Guardian Backend team members adopt a number of conventions in coding
style and structure. The conventions are outlined here for the benefit of
contributors, and for our own benefit.

For setting up a development environment, see [Developing in this
repository](DEVELOPMENT.md).

## Build constraints

There are a few places where we may implement only Linux versions of (usually
low-level) functions. These are in `*_linux.go` files. The aim is to develop
code that is (at least) testable on any platform, which may mean implementing
linux and non-linux versions of some low-level code.

Every such file will also record its constraints in build constraints in the
file content. For example, even if a file is called `fred_linux.go` it should
still contain this build constraint:

```go
// +build linux
```

The rule to follow is: the build constriants are definitive and complete; the
filenames are used to indicate the constraints. (***Note***: the build constraint
`// +build !linux` must *not* be indicated by a file name ending in
`_non_linux.go`. This will not work.)

## Code formatting

We follow the usual convention: before committing any code the `*.go` files are
all formatted by `go fmt`.

This can be automated by putting a symbolic link to
`development/pre-commit-hook/pre-commit` in the `.git/hooks` directory for the
repository. (See step 5 in [Setting up a development
environment](DEVELOPMENT.md#setting-up-a-development-environment).)

## Diagnostics

### Errors

Errors returned from Guardian backend functions and methods include error
identifiers to uniquely identify each kind of failure and stack traces so that
the point of failure can easily be determined. The [gerror](gerror) package is
used to construct errors.

*This principle is an aim, not an achieved goal.*

### Logging

Logging is performed using the [glog](https://github.com/golang/glog) package
(an external dependency).

`glog` has four logging levels:

*   informational—`glog.Info()`,
*   warning—`glog.Warning()`,
*   error—`glog.Error()`, and
*   fatal—`glog.Fatal()`.

Additionally, `glog` has a verbosity level  which is a non-negative integer
with a default value of 0. We use informational logs with verbosity
level 0 to log important events and informational logs with verbosity levels
1 and higher for detailed debugging.

Logs may be directed to standard error by setting the flag `logtostderr` to
`true` on the go invocation, as in this example:

```bash
go test -logtostderr=true -vmodule=*=2
```

***Note***: the glog `-v` flag clashes with the boolean `-v` flag of `go test`
and so the logging verbosity should be set during testing using `-vmodule=*=`.

See the [glog documentation](http://godoc.org/github.com/golang/glog) for further information.

## Independent unit testing

Unit tests are (supposed to be) comprehensive and based around the interface types.

We intend to run ‘independent’ tests, that is, to test an implementation
without driving the full implementation of the code it depends upon.

How is this achieved? By designing the mainline code with [Dependency
injection](#dependency-injection) and exploiting that in our tests.

To get this to work, we need to have test versions of the
*dependencies* of something we are testing, so that we can check they are
driven correctly (or provoke the right sort of errors) in the code under test.

We can then inject the test versions of the dependencies when testing a component.

In this repository we exploit [`gomock`](#gomock) to help build mock versions of
some of our interfaces, and to use them. The mock objects need only be
constructed once, and are checked in so that builders need not regenerate them
every time. They will break if their interfaces change.

### `gomock`

We use [gomock](http://godoc.org/code.google.com/p/gomock/gomock) to generate
mock implementations of some interfaces to enable packages to be easily tested
in isolation from one another.

( Once you have installed Go, run these commands to install the gomock package
  and the mockgen tool:

```bash
go install code.google.com/p/gomock/gomock
go install code.google.com/p/gomock/mockgen
```

See [gomock README](https://code.google.com/p/gomock/source/browse/README.md)
for the rest of the instructions. )

Mocks are stored in a subdirectory of the directory containing the mocked
interface. For example, the mocks for the `kernel/fileutils` package are
stored in `kernel/fileutils/mock_fileutils`.

To re-generate the generated mock implementations of interfaces, ensure
`mockgen` is on the path and then run the script
`development/scripts/update_mocks`. If you add a generated mock, don't forget
to add it to `update_mocks`.

### Dependency injection

In order to do black-box unit testing of individual interface implementations
we introduce a convention for “wiring” implementations together at create
time.

This is a lightweight version of ‘dependency injection’ which Spring Framework
implements in Java.  We use no language add-ons, pre-processing, annotations
nor non-standard syntax.

The convention provides a standard way of constructing interface values and
“wiring them together” with their dependencies. Runtime code has a simple
(standard) way of obtaining a runtime value of an interface and test code can
inject mocks and stubs for dependencies as necessary.

The convention involves some "wiring" functions typically defined in
a file called `wiring.go`.

The convention associates an interface with a pair of constructor functions of
the form:

```go
Wire(parms, depParms) (InterfaceType, gerror.Gerror)
WireWith(parms, dependencies) (InterfaceType, gerror.Gerror)
```

where `InterfaceType` is the interface type value being wired (constructed).
We mandate `gerror.Gerror` to be returned, although the caller may treat that
as an ordinary `error` value.

If there is a single interface in a package, then the function names are
`Wire` and `WireWith` as shown. If there are multiple interfaces defined in a
package, the function names can be decorated, for example `WireFS` and
`WireFSWith`.

#### The wire functions

```go
Wire(parms, depParms)
```

constructs a default value of the interface with dependencies injected.
`parms` indicates zero or more construction parameters, usually not of
interface type. `depParms` indicates zero or more additional parameters needed
to construct the dependencies. `Wire()` typically constructs any dependencies
by passing `depParms` to the corresponding `Wire()` functions of its
dependencies and then calls `WireWith()` passing `parms` and the constructed
dependencies.

```go
WireWith(parms, dependencies)
```

constructs a specific value of the interface with the given dependencies. This
is used by `Wire()` to create the “standard” value. It may also be used by
test code to create a value in which some or all of its dependencies are mocks
or stubs.

`WireWith` should not normally be used outside its package or the test
package, but for consistency, and for use by so-called black-box test
packages, this function, as well as `Wire()`, should be exported.

The `Wire` and `WireWith` functions are placed by convention in the file
`wiring.go`, although this convention is not enforced. The complexity of
`wiring.go` can be minimised by having `WireWith` call a package private function (in
another file) to construct the value. The private function is typically
of the form:

```go
newXXX(parms, dependencies) (InterfaceType, gerror.Gerror)
```

where `XXX` is usually the interface type name, but more generally
identifies the concrete type which implements the interface, for
example when the package provides more than one implementation of
the interface.

For a prototypical example of the use of `Wire` and `WireWith`, with dependencies, see
[guardian_backend](https://github.com/cf-guardian/guardian-backend/blob/master/guardian_backend/wiring.go).
