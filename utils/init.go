package utils

import (
	"log"
	"runtime"
)

// init.go is a collection of utilities useful in init() functions for packages.

// OptimiseScheduling by using all CPUs for scheduling goroutines. The default in Go 1.3 is to use only one CPU.
func OptimiseScheduling() {
	cpus := runtime.NumCPU()
	prevMaxProcs := runtime.GOMAXPROCS(cpus)
	// TODO: the following is too chatty. Consider using glog for log levels.
	log.Println("utils: GOMAXPROCS set to", cpus, "from", prevMaxProcs)
}
