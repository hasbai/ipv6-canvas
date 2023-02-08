package lib

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"sync"
)

const ImageSavePath = "save.png"
const SIZE = 256

type Image struct {
	*image.RGBA
	*sync.Mutex
}

func (img Image) Modify(p Pixel) {
	img.Lock()
	defer img.Unlock()
	img.Set(p.X, p.Y, color.RGBA{})
}

// Encode encodes an Image to png
func (img Image) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (img Image) Save() {
	data, err := img.Encode()
	if err != nil {
		log.Fatalf("error saving image %v", err)
	}
	os.WriteFile(ImageSavePath, data, 0644)
}

func LoadImage(path string) Image {
	fp, err := os.Open(path)
	defer fp.Close()
	if err != nil {
		log.Println("image does not exist, create one")
		return Image{
			RGBA: image.NewRGBA(image.Rectangle{
				Min: image.Point{},
				Max: image.Point{X: SIZE, Y: SIZE},
			}),
			Mutex: &sync.Mutex{},
		}
	}
	log.Println("decoding image...")
	img, err := png.Decode(fp)
	if err != nil {
		log.Fatalf("error decoding image %v", err)
	}
	return Image{
		RGBA:  img.(*image.RGBA),
		Mutex: &sync.Mutex{},
	}
}
