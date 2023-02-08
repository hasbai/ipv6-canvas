package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Serve() {
	go serveHTTP()
	go serveICMP()

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	s := <-c
	fmt.Println("exiting... ", s)
	exit()
}

func exit() {
	os.Exit(0)
}
