// Command line options for main guardian executable.
package options

import (
	"flag"
	"log"
	"os"
	"time"
)

type Options struct {
	DepotPath     string
	ListenNetwork string
	ListenAddr    string
	GraceTime     time.Duration
}

var actionOnError func(error) = defaultActionOnError

func defaultActionOnError(err error) {
	os.Exit(2)
}

// Returns the options corresponding to the supplied args after applying defaults.
func Parse(args []string) *Options {
	opts := Options{}

	flagset := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flagset.StringVar(&opts.DepotPath, "depot", "", "directory in which to store containers")
	flagset.StringVar(&opts.ListenNetwork, "listenNetwork", "unix", "listener network (see net.Listen)")
	flagset.StringVar(&opts.ListenAddr, "listenAddr", "/tmp/warden.sock", "listener address (see net.Listen)")
	flagset.DurationVar(&opts.GraceTime, "containerGraceTime", 0, "time after which to destroy idle containers")

	if err := flagset.Parse(args[1:]); err != nil {
		actionOnError(err)
	}

	if opts.DepotPath == "" {
		log.Println("-depot flag must be provided")
		actionOnError(nil)
	}

	return &opts
}
