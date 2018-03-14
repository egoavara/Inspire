package igl

import (
	"image"
	"image/draw"
)

type Texture interface {
	Use(unit uint32)
}

type SingleTexture interface {
	Texture
	Get() image.Image
	Read(dst draw.Image) error
	Write(src image.Image) error
	Size() (w, h int32)
}
type StreamTexture interface {
	Texture
	Skip()
	Write(src image.Image) error
	Size() (w, h int32)
}