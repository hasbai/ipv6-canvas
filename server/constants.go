package server

import (
	"github.com/hasbai/ipv6-canvas/lib"
)

var IMG lib.Image

const HttpAddr = "localhost:8000"

func init() {
	IMG = lib.LoadImage(lib.ImageSavePath)
}
