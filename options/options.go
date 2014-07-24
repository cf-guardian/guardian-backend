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

// Returns the options corresponding to the command line flags after applying defaults.
func Parse() *Options {
	opts := Options{}

	flag.StringVar(&opts.DepotPath, "depot", "", "directory in which to store containers")
	flag.StringVar(&opts.ListenNetwork, "listenNetwork", "unix", "listener network (see net.Listen)")
	flag.StringVar(&opts.ListenAddr, "listenAddr", "/tmp/warden.sock", "listener address (see net.Listen)")
	flag.DurationVar(&opts.GraceTime, "containerGraceTime", 0, "time after which to destroy idle containers")

	flag.Parse()

	if opts.DepotPath == "" {
		log.Fatalln("-depot must be specified")
	}

	return &opts
}
