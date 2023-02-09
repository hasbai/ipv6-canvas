package lib

import (
	"fmt"
	"image"
	"log"
	"net"
)

func ParsePoint(s string) image.Point {
	if s == "" {
		return image.Point{}
	}
	var x, y int
	_, err := fmt.Sscanf(s, "%d,%d", &x, &y)
	if err != nil {
		log.Fatalf("point must be in the format <x>,<y>")
	}
	if x <= 0 || y <= 0 {
		log.Fatalf("point must be positive")
	}
	return image.Point{X: x, Y: y}
}

func ParsePrefix(cidr string) net.IP {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		log.Fatalf("parse CIDR error, got %s, %v", cidr, err)
	}
	if len(ipNet.IP) != 16 {
		log.Fatalf("must be an ipv6 cidr, got %s", cidr)
	}
	return ipNet.IP[:8]
}
