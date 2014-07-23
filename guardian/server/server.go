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

func StartServer(depotPath string, listenNetwork string, listenAddr string, graceTime time.Duration) <-chan int {
	backend := guardian_backend.New(depotPath)
	return runServer(backend, listenNetwork, listenAddr, graceTime)
}

func runServer(backend warden.Backend, listenNetwork string, listenAddr string, graceTime time.Duration) <-chan int {
	wardenServer := server.New(listenNetwork, listenAddr, graceTime, backend)
	err := wardenServer.Start()
	if err != nil {
		log.Fatalln("failed to start server:", err)
	}

	log.Println("server started; listening over", listenNetwork, "on", listenAddr)

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
