package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/cloudfoundry-incubator/garden/server"
	"github.com/cf-guardian/guardian-backend/guardian_backend"
	"github.com/cloudfoundry-incubator/garden/warden"
)

var listenNetwork = flag.String(
	"listenNetwork",
	"unix",
	"listener protocol (unix, tcp, etc. - see net.Listen)",
)


var listenAddr = flag.String(
	"listenAddr",
	"/tmp/warden.sock",
	"listener address (see net.Listen)",
)

var depotPath = flag.String(
	"depot",
	"",
	"directory in which to store containers",
)

var containerGraceTime = flag.Duration(
	"containerGraceTime",
	0,
	"time after which to destroy idle containers",
)

func main() {
	optimiseScheduling()

	flag.Parse()

	if *depotPath == "" {
		log.Fatalln("must specify -depot with linux backend")
	}

	backend := guardian_backend.New(*depotPath)

	runServer(backend, *listenNetwork, *listenAddr, *containerGraceTime)
}

func runServer(backend warden.Backend, listenNetwork string, listenAddr string, graceTime time.Duration) {
	wardenServer := server.New(listenNetwork, listenAddr, graceTime, backend)

	err := wardenServer.Start()
	if err != nil {
		log.Fatalln("failed to start:", err)
	}

	log.Println("server started; listening with", listenNetwork, "on", listenAddr)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-signals
	log.Println("stopping server...")
	wardenServer.Stop()
}

func optimiseScheduling() {
	cpus := runtime.NumCPU()
	prevMaxProcs := runtime.GOMAXPROCS(cpus)
	log.Println("GOMAXPROCS set to", cpus, "from", prevMaxProcs)
}
