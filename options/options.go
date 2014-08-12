// Command line options for main guardian executable.
package options

import (
	"time"
	"log"
	"flag"
)

type Options struct {
	DepotPath     string
	ListenNetwork string
	ListenAddr    string
	GraceTime     time.Duration
}

// Returns the options corresponding to the supplied args after applying defaults.
func Parse(args []string) *Options {
	opts := Options{}

	flagset := flag.NewFlagSet(args[0], flag.ExitOnError)
	flagset.StringVar(&opts.DepotPath, "depot", "", "directory in which to store containers")
	flagset.StringVar(&opts.ListenNetwork, "listenNetwork", "unix", "listener network (see net.Listen)")
	flagset.StringVar(&opts.ListenAddr, "listenAddr", "/tmp/warden.sock", "listener address (see net.Listen)")
	flagset.DurationVar(&opts.GraceTime, "containerGraceTime", 0, "time after which to destroy idle containers")

	flagset.Parse(args[1:])

	if opts.DepotPath == "" {
		log.Fatalln("-depot flag must be provided")
	}

	return &opts
}
