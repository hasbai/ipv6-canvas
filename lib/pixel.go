package lib

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"log"
	"net"
)

type Pixel struct {
	Point image.Point `json:"point"`
	Color color.NRGBA `json:"color"`
}

func (p *Pixel) Marshal() []byte {
	data := make([]byte, 9)
	data[0] = MessageTypePixel
	data[1] = byte(p.Point.X >> 8)
	data[2] = byte(p.Point.X)
	data[3] = byte(p.Point.Y >> 8)
	data[4] = byte(p.Point.Y)
	data[5] = p.Color.R
	data[6] = p.Color.G
	data[7] = p.Color.B
	data[8] = p.Color.A
	return data
}

func (p *Pixel) Unmarshal(data []byte) error {
	if len(data) != 9 || data[0] != MessageTypePixel {
		return fmt.Errorf("invalid pixel data: %v", data)
	}
	p.Point.X = int(data[1])<<8 + int(data[2])
	p.Point.Y = int(data[3])<<8 + int(data[4])
	p.Color.R = data[5]
	p.Color.G = data[6]
	p.Color.B = data[7]
	p.Color.A = data[8]
	if p.Point.X < 0 || p.Point.Y < 0 || p.Point.X >= SIZE || p.Point.Y >= SIZE {
		return fmt.Errorf("pixel point should between 0 and %d", SIZE)
	}
	return nil
}

func IP2Pixel(ip net.IP) (Pixel, error) {
	data := ip[7:16]
	data[0] = MessageTypePixel
	p := Pixel{}
	err := p.Unmarshal(data)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (p *Pixel) String() string {
	data, err := json.Marshal(p)
	if err != nil {
		log.Fatalf("Error marshalling pixel: %s", err)
	}
	return string(data)
}
