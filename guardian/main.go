package main

import (
	"flag"
	"github.com/cf-guardian/guardian-backend/guardian/server"
	"log"
	"os"
	"runtime"
)

var (
	listenNetwork = flag.String(
		"listenNetwork",
		"unix",
		"listener protocol (unix, tcp, etc. - see net.Listen)",
	)

	listenAddr = flag.String(
		"listenAddr",
		"/tmp/warden.sock",
		"listener address (see net.Listen)",
	)

	depotPath = flag.String(
		"depot",
		"",
		"directory in which to store containers",
	)

	containerGraceTime = flag.Duration(
		"containerGraceTime",
		0,
		"time after which to destroy idle containers",
	)
)

func main() {
	optimiseScheduling()

	flag.Parse()

	if *depotPath == "" {
		log.Fatalln("must specify -depot with guardian backend")
	}

	os.Exit(<-server.StartServer(*depotPath, *listenNetwork, *listenAddr, *containerGraceTime))
}

// Use all CPUs for scheduling goroutines. The default in Go 1.3 is to use only one CPU.
func optimiseScheduling() {
	cpus := runtime.NumCPU()
	prevMaxProcs := runtime.GOMAXPROCS(cpus)
	log.Println("GOMAXPROCS set to", cpus, "from", prevMaxProcs)
}
