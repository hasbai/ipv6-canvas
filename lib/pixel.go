package lib

import (
	"encoding/json"
	"image"
	"image/color"
	"log"
	"net"
)

type Pixel struct {
	Point image.Point `json:"point"`
	Color color.NRGBA `json:"color"`
}

func IP2Pixel(ip net.IP) Pixel {
	return Pixel{
		Point: image.Point{
			X: int(ip[8])<<8 + int(ip[9]),
			Y: int(ip[10])<<8 + int(ip[11]),
		},
		Color: color.NRGBA{
			R: ip[12],
			G: ip[13],
			B: ip[14],
			A: ip[15],
		},
	}
}

func Pixel2IP(p Pixel) net.IP {
	ip := make(net.IP, 16)
	ip[8] = byte(p.Point.X >> 8)
	ip[9] = byte(p.Point.X)
	ip[10] = byte(p.Point.Y >> 8)
	ip[11] = byte(p.Point.Y)
	ip[12] = p.Color.R
	ip[13] = p.Color.G
	ip[14] = p.Color.B
	ip[15] = p.Color.A
	return ip
}

func (p *Pixel) String() string {
	data, err := json.Marshal(p)
	if err != nil {
		log.Fatalf("Error marshalling pixel: %s", err)
	}
	return string(data)
}
