package lib

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/png"
	"log"
	"os"
	"sync"
)

const ImageSavePath = "save.png"
const SIZE = 256

type Image struct {
	*image.RGBA
	sync.Mutex
}

func (img *Image) Modify(p Pixel) {
	img.Lock()
	defer img.Unlock()
	img.Set(p.X, p.Y, p.Color)
}

func (img *Image) Resize(p image.Point) {
	if p.X <= 0 || p.Y <= 0 {
		return
	}
	img.Lock()
	defer img.Unlock()
	img.RGBA = (*image.RGBA)(imaging.Resize(img, p.X, p.Y, imaging.Lanczos))
}

// Encode encodes an Image to png
func (img *Image) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (img *Image) Save(path string) {
	data, err := img.Encode()
	if err != nil {
		log.Fatalf("error saving image %v", err)
	}
	os.WriteFile(path, data, 0644)
}

func LoadImage(path string) Image {
	fp, err := os.Open(path)
	defer fp.Close()
	if err != nil {
		return Image{
			RGBA: image.NewRGBA(image.Rectangle{
				Min: image.Point{},
				Max: image.Point{X: SIZE, Y: SIZE},
			}),
		}
	}
	log.Println("decoding image...")
	img, err := png.Decode(fp) // returns NRGBA
	if err != nil {
		log.Fatalf("error decoding image %v", err)
	}
	return Image{
		RGBA: (*image.RGBA)(img.(*image.NRGBA)),
	}
}

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
