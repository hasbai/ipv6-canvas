package lib

import (
	"bytes"
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
	*image.NRGBA
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
	img.NRGBA = imaging.Resize(img, p.X, p.Y, imaging.Lanczos)
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
	var out *image.NRGBA
	if err != nil {
		return NewImage()
	}

	log.Println("decoding image...")
	img, err := png.Decode(fp)
	if err != nil {
		log.Fatalf("error decoding image %v", err)
	}
	switch img := img.(type) {
	case *image.RGBA:
		out = (*image.NRGBA)(img)
	case *image.NRGBA:
		out = img
	default:
		panic("image format not supported")
	}
	return Image{NRGBA: out}
}

func NewImage() Image {
	img := image.NewNRGBA(image.Rectangle{
		Min: image.Point{},
		Max: image.Point{X: SIZE, Y: SIZE},
	})
	for x := 0; x < SIZE; x++ {
		for y := 0; y < SIZE; y++ {
			img.Set(x, y, image.White)
		}
	}
	return Image{NRGBA: img}
}
