package utils

import (
	"os"
	"os/signal"
)

// Handle a given collection of signals by waiting for one of them to occur, passing that signal
// to the given function, and sending the output of the function to the given channel.
func HandleSignals(onSignal func(os.Signal) int, exitChan chan<- int, signals ...os.Signal) {
	signalsChan := make(chan os.Signal, 1)
	signal.Notify(signalsChan, signals...)
	exitChan<-onSignal(<-signalsChan)
}
