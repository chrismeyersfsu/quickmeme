package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/gotk3/gotk3/gtk"
)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func warnIf(err error) {
	if err != nil {
		log.Println(err)
	}
}

func CancelAll(cancel []context.CancelFunc) {
	for _, c := range cancel {
		c()
	}
}

func _ConsumeSignals(cancel []context.CancelFunc, sigchan <-chan os.Signal) {
	// TODO: Should we loop over the signals or just break after the first?
	// Maybe we should have an if condition for SIGINT
	for range sigchan {
		CancelAll(cancel)
	}
	gtk.MainQuit()
	os.Exit(0)
}

func StartSignalHandlers(cancel []context.CancelFunc) {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	go _ConsumeSignals(cancel, sigchan)
}
