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
	runServer(backend, listenNetwork, listenAddr, graceTime)

	exitChan := make(chan int, 1)
	exitChan <- 0
	return exitChan
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
