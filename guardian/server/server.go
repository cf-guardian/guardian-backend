package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cf-guardian/guardian-backend/guardian_backend"
	"github.com/cf-guardian/guardian-backend/options"
	"github.com/cf-guardian/guardian-backend/utils"
	"github.com/cloudfoundry-incubator/garden/server"
	"github.com/cloudfoundry-incubator/garden/warden"
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

	exitChan := make(chan int, 1)
	go stopOnSignals(exitChan, func() int {
		log.Println("stopping server...")
		wardenServer.Stop()
		return 0
	}, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	return exitChan
}

func init() {
	utils.OptimiseScheduling()
}

func stopOnSignals(exitChan chan int, stopFunc func() int, signals ...Signal) {
	signalsChan := make(chan os.Signal, 1)
	signal.Notify(signalsChan, signals...)
	<-signalsChan

	exitChan <- stopFunc()
}
