package main

import (
	"fmt"
	"github.com/hasbai/ipv6-canvas/client"
	"github.com/hasbai/ipv6-canvas/lib"
	"github.com/hasbai/ipv6-canvas/server"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "ipv6-canvas",
		Usage: "ping ipv6 address to draw on canvas",
		Action: func(*cli.Context) error {
			fmt.Println("boom! I say!")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "serve",
				Aliases: []string{},
				Usage:   "serve",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "listen",
						Aliases:  []string{"l"},
						Value:    "localhost:8000",
						Usage:    "http address to listen on",
						Category: "serve",
					},
				},
				Action: func(ctx *cli.Context) error {
					httpAddr := ctx.String("listen")
					server.Serve(httpAddr)
					return nil
				},
			},
			{
				Name:    "ping",
				Aliases: []string{},
				Usage:   "ping <addr>",
				Action: func(ctx *cli.Context) error {
					client.Ping(ctx.Args().First())
					return nil
				},
			},
			{
				Name:    "draw",
				Aliases: []string{},
				Usage:   "draw <filepath> <cidr>",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "resize",
						Aliases:  []string{"r"},
						Usage:    "resize the image in the format of <width>,<height>",
						Category: "draw",
					},
					&cli.StringFlag{
						Name:     "offset",
						Aliases:  []string{"o"},
						Usage:    "the start point to draw, in the format of <x>,<y>",
						Category: "draw",
					},
				},
				Action: func(ctx *cli.Context) error {
					if ctx.Args().Len() < 2 {
						return fmt.Errorf("usage: draw <filepath> <cidr>")
					}
					client.Draw(
						ctx.Args().Get(0),
						ctx.Args().Get(1),
						lib.ParsePoint(ctx.String("resize")),
						lib.ParsePoint(ctx.String("offset")),
					)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
