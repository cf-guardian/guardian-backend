package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cf-guardian/guardian-backend/guardian_backend"
	"github.com/cloudfoundry-incubator/garden/server"
	"github.com/cloudfoundry-incubator/garden/warden"
)

type Options struct {
	DepotPath     string
	ListenNetwork string
	ListenAddr    string
	GraceTime     time.Duration
}

func StartServer(opts *Options) <-chan int {
	backend := guardian_backend.New(opts.DepotPath)
	return runServer(backend, opts)
}

func runServer(backend warden.Backend, opts *Options) <-chan int {
	wardenServer := server.New(opts.ListenNetwork, opts.ListenAddr, opts.GraceTime, backend)
	err := wardenServer.Start()
	if err != nil {
		log.Fatalln("failed to start server:", err)
	}

	log.Println("server started; listening over", opts.ListenNetwork, "on", opts.ListenAddr)

	// TODO[sp]: make runServer asynchronous
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-signals
	log.Println("stopping server...")
	wardenServer.Stop()

	exitChan := make(chan int, 1)
	exitChan <- 0
	return exitChan
}
