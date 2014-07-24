package main

import (
	"flag"
	"github.com/cf-guardian/guardian-backend/guardian/server"
	"github.com/cf-guardian/guardian-backend/utils"
	"log"
	"os"
)

// Main program for warden server with guardian backend.
func main() {
	os.Exit(<-server.StartServer(parseOptions()))
}

func parseOptions() *server.Options {
	opts := server.Options{}

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

func init() {
	utils.OptimiseScheduling()
}

