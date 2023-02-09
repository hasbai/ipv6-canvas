package server

import (
	"github.com/hasbai/ipv6-canvas/lib"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Serve(httpAddr string) {
	go serveHTTP(httpAddr)
	go serveICMP()

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	s := <-c
	log.Println("received", s)
	exit()
}

func exit() {
	IMG.Save(lib.ImageSavePath)
	log.Printf("image saved to %s, exiting...", lib.ImageSavePath)
	os.Exit(0)
}
