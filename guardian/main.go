package main

import (
	"github.com/cf-guardian/guardian-backend/guardian/server"
	"github.com/cf-guardian/guardian-backend/options"
	"os"
)

// Main program for warden server with guardian backend.
func main() {
	os.Exit(<-server.StartServer(options.Parse()))
}
