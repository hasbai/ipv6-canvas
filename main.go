package main

import (
	"github.com/hasbai/ipv6-canvas/client"
	"github.com/hasbai/ipv6-canvas/server"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		server.Serve()
	} else {
		switch os.Args[1] {
		case "ping":
			if len(os.Args) < 3 {
				panic("usage: ping <addr>")
			}
			client.Ping(os.Args[2])
		case "draw":
			if len(os.Args) < 4 {
				panic("usage: draw <filepath> <cidr>")
			}
			client.Draw(os.Args[2], os.Args[3])
		default:
			panic("unknown command")
		}
	}
}
