# Developing in this repository

This code is written (almost exclusively) in Go. To make changes or to
contribute to this code you should read this file, which explains how to set
up a development environment, and [Conventions used in this
repository](CONVENTIONS.md) which describes the design and coding rules which
we have tried to follow.

## Setting up a development environment

1.  Ensure the following pre-requisites are installed:
    * [git](http://git-scm.com/downloads)
    * [Go](http://golang.org/) (1.3 or later):
        - either [download](http://golang.org/doc/install) a specific version
        - or use [gvm](https://github.com/moovweb/gvm)
        - or even `port install go` with [MacPorts](http://www.macports.org/).

2.  Create a Go [workspace](http://golang.org/doc/code.html#Organization)
    directory, such as `$HOME/go`, and add the path of this directory to the
    beginning of a new environment variable called `GOPATH`. You might want to put
    this last step in your profile (or in a file `$HOME/.go-setup`, see the
    scripts section below).

    ```bash
    $ mkdir $HOME/go
    $ export GOPATH=$HOME/go
    ```

3.  Get this repository into your workspace (under the `src` directory) by issuing:

    ```bash
    $ go get github.com/cf-guardian/guardian
    ```

4.  Change directory to `<workspace dir>/src/github.com/cf-guardian/guardian`.

5.  Install the [pre-commit hook](https://github.com/jbrukh/git-gofmt) as follows:

    ```bash
    cd .git/hooks
    ln -s ../../development/pre-commit-hook/pre-commit .
    ```

    After installing the hook, if you need to skip reformatting for a
    particular commit, use the `--no-verify` option on the `git commit`
    command.

6.  For optional setup, see the following [Development
    scripts](#developmentscripts) section, and the sections on
    [`gomock`](#gomock) and [Editing](#editing).

### Development scripts

The directory `development/scripts` includes some simple (bash) shell scripts
that may make life a little easier:

#### `gosub`

The **`gosub`** script is designed to run a `go` command in all subdirectories
of the current directory that have a `*.go` file in them. Typical usage is
`gosub build`, or `gosub fmt build`. This is not very useful, as `go xxx
./...` does (almost) the same thing.

The `gosub` script is limited to single word go commands. `gosub fmt build`
would issue `go fmt; go build` in each directory in turn. `gosub test` will
run tests in immediate subdirectories.

#### `govet`

The **`govet`** script runs `go tool vet` with appropriate options against a
single directory passed as a parameter.

#### `goroot`

The **`goroot`** script is suitable for running the `go` tool as root. It
"sources" the `$HOME/.go-setup` file, which can set up the Go environment
variables, and then executes the go tool with the parameters provided.
Something like this is necessary because root doesn't normally have a go tool
environment in its shell.

#### `update_mocks`

The **`update_mocks`** script will update the mocked interfaces used in
the unit tests. See the [`gomock`](CONVENTIONS.md#gomock) section in [Testing](#testing) below.

## Testing

To run the tests in a specific directory, issue:

```bash
go test
```

If the tests succeed, this should print `PASS`.

For test construction and some guidelines, see the [Conventions used in this
repository](CONVENTIONS.md).

## Editing

If your favourite text editor is not sufficient, try
[Eclipse](http://www.eclipse.org/downloads/) with the [goclipse
plugin](https://github.com/sesteel/goclipse) or [IntelliJ
IDEA](http://www.jetbrains.com/idea/) with the
[go plugin](https://github.com/go-lang-plugin-org/go-lang-idea-plugin).

Source code is formatted according to standard Go conventions. To re-format
all the code, issue the following from the root:

```
go fmt ./...
```

To reformat code files automatically whenever they are committed to git, install the
`pre-commit` hook as described in step 5 above.

Also, you can run `go vet` (a `govet` script is provided in `development/scripts`) and
[lint](http://go-lint.appspot.com/github.com/cf-guardian/guardian) against the code if you like.
