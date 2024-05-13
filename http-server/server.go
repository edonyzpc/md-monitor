package main

import (
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/edony-ink/log"
)

const (
	LogFile  = "./bin/server.log"
	SockFile = "./bin/server.sock"
)

func init() {
	log.SWLogger.Init(LogFile, log.DebugLevel, true)
}

func main() {
	err := os.Remove(SockFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Info("socket file not exist")
		} else {
			log.Panic(err)
		}
	}
	listener, err := net.Listen("unix", SockFile)
	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()

	mutex := http.NewServeMux()
	mutex.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello, world!"))
		if err != nil {
			log.Panic(err)
		}
		log.Info("hello world")
	})

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		// block here
		log.Info("blocking...")
		<-sig
		// cleanup socket resource
		log.Info("cleanup socket file")
		os.Remove(SockFile)
		// exit
		os.Exit(0)
	}()

	err = http.Serve(listener, mutex)
	if err != nil {
		log.Panic(err)
	}
}
