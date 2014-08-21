package server

import (
	"log"
	"os"
	"syscall"

	"github.com/cf-guardian/guardian-backend/guardian_backend"
	"github.com/cf-guardian/guardian-backend/options"
	"github.com/cf-guardian/guardian-backend/utils"
	"github.com/cloudfoundry-incubator/garden/server"
	"github.com/cloudfoundry-incubator/garden/warden"
	"github.com/pivotal-golang/lager"
)

func StartServer(opts *options.Options) <-chan int {
	backend, err := guardian_backend.Wire(opts.DepotPath, opts.DepotPath)
	if err != nil {
		log.Printf("StartServer failed to wire GuardianBackend: ", err)
		exitChan := make(chan int, 1)
		exitChan<-1
		return exitChan
	}
	return runServer(backend, opts)
}

func runServer(backend warden.Backend, opts *options.Options) <-chan int {
	logger := lager.NewLogger("guardian")
	wardenServer := server.New(opts.ListenNetwork, opts.ListenAddr, opts.GraceTime, backend, logger)
	err := wardenServer.Start()
	if err != nil {
		log.Fatalln("failed to start server:", err)
	}

	log.Println("server started; listening over", opts.ListenNetwork, "on", opts.ListenAddr)

	return stopOnSignal(wardenServer)
}

func stopOnSignal(wardenServer *server.WardenServer) <-chan int {
	exitChan := make(chan int, 1)
	go utils.HandleSignals(func(os.Signal) int {
			log.Println("stopping server...")
			wardenServer.Stop()
			return 0
		}, exitChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	return exitChan
}

func init() {
	utils.OptimiseScheduling()
}
