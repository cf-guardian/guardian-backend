package main

import (
	"github.com/cf-guardian/guardian-backend/guardian/server"
	"github.com/cf-guardian/guardian-backend/options"
	"github.com/cf-guardian/guardian-backend/utils"
	"os"
)

// Main program for warden server with guardian backend.
func main() {
	os.Exit(<-server.StartServer(options.ParseOptions()))
}

func init() {
	utils.OptimiseScheduling()
}

