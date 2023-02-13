package client

import (
	"github.com/hasbai/ipv6-canvas/lib"
	"image"
	"log"
)

func Draw(filepath, cidr string, resize, offset image.Point) {
	conn := dial()
	defer conn.Close()
	prefix := lib.ParsePrefix(cidr)

	img := lib.LoadImage(filepath)
	img.Resize(resize)
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			p := lib.Pixel{Point: image.Point{X: x + offset.X, Y: y + offset.Y}, Color: img.NRGBAAt(x, y)}
			if p.Color.R == 255 && p.Color.G == 255 && p.Color.B == 255 {
				continue
			}
			ip := append(prefix, p.Marshal()[1:]...)
			err := ping(ip, conn)
			if err != nil {
				log.Fatalf("error ping %s: %v", ip, err)
			}
		}
	}
}
