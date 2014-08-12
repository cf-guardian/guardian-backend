package options

import (
	"testing"
)

var (
	errorOccurred bool
	errorPassed error
)

func setErrorCatcher() {
	errorOccurred = false
	errorPassed = nil
	actionOnError = testingActionOnError
}

func testingActionOnError(err error) {
	errorOccurred = true
	errorPassed = err
}

func TestDepot(t *testing.T) {
	setErrorCatcher()
	args := []string{`test`, `-depot=/var/test`}
	opts := Parse(args)
	if errorOccurred {
		t.Errorf("Failed (err=%q) to parse %q", errorPassed, args)
	}
	if opts.DepotPath != `/var/test` {
		t.Errorf("Incorrect depot path %q, expected %q", opts.DepotPath, "/var/test")
	}
}

func TestOmittedDepot(t *testing.T) {
	setErrorCatcher()
	args := []string{"test"}
	Parse(args)
	if !errorOccurred {
		t.Error("should have failed")
	}
}
