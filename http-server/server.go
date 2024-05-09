package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := os.Remove("/tmp/server.sock")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("socket file not exist")
		} else {
			panic(err)
		}
	}
	listener, err := net.Listen("unix", "/tmp/server.sock")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	mutex := http.NewServeMux()
	mutex.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello, world!"))
		if err != nil {
			panic(err)
		}
		fmt.Println("hello world")
	})

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		// block here
		fmt.Println("blocking...")
		<-sig
		// cleanup socket resource
		fmt.Println("cleanup socket file")
		os.Remove("/tmp/server.sock")
		// exit
		os.Exit(0)
	}()

	err = http.Serve(listener, mutex)
	if err != nil {
		panic(err)
	}
}
