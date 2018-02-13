package igl

import (
	"image"
	"image/draw"
)

type Texture interface {
	Size() (w, h int32)
	Use(unit uint32)
}
type TextureWriter interface {
	Write(src image.Image) error
	Texture
}
type TextureReader interface {
	Read(dst draw.Image) error
	Texture
}
type TextureReadWriter interface {
	Read(dst draw.Image) error
	Write(src image.Image) error
	Texture
}