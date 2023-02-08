package client

import (
	"github.com/hasbai/ipv6-canvas/lib"
	"log"
	"net"
	"time"
)

func Draw(filepath, cidr string) {
	conn := dial()
	defer conn.Close()
	prefix := parsePrefix(cidr)

	img := lib.LoadImage(filepath)
	for i := img.Bounds().Min.X; i < img.Bounds().Max.X; i++ {
		for j := img.Bounds().Min.Y; j < img.Bounds().Max.Y; j++ {
			p := lib.Pixel{X: i, Y: j, Color: img.RGBAAt(i, j)}
			if p.Color.R == 255 && p.Color.G == 255 && p.Color.B == 255 {
				continue
			}
			ip := append(prefix, lib.Pixel2IP(p)[8:]...)
			err := ping(ip, conn)
			if err != nil {
				log.Fatalf("error ping %s: %v", ip, err)
			}
			time.Sleep(time.Second / 2)
		}
	}
}

func parsePrefix(cidr string) net.IP {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		log.Fatalf("parse CIDR error, got %s, %v", cidr, err)
	}
	if len(ipNet.IP) != 16 {
		log.Fatalf("must be an ipv6 cidr, got %s", cidr)
	}
	return ipNet.IP[:8]
}
