package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cf-guardian/guardian-backend/guardian_backend"
	"github.com/cf-guardian/guardian-backend/options"
	"github.com/cloudfoundry-incubator/garden/server"
	"github.com/cloudfoundry-incubator/garden/warden"
	"github.com/cf-guardian/guardian-backend/utils"
)

func StartServer(opts *options.Options) <-chan int {
	backend := guardian_backend.New(opts.DepotPath)
	return runServer(backend, opts)
}

func runServer(backend warden.Backend, opts *options.Options) <-chan int {
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

func init() {
	utils.OptimiseScheduling()
}
