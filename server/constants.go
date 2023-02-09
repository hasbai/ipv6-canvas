package server

import (
	"github.com/hasbai/ipv6-canvas/lib"
)

var IMG lib.Image

func init() {
	IMG = lib.LoadImage(lib.ImageSavePath)
}
