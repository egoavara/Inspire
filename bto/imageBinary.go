package bto

import (
	"image"
	"image/draw"
)

func Image(src image.Image) (pix []uint8) {
	if v := IsRootRGBA(src); v != nil{
		return v.Pix
	}
	temp := image.NewRGBA(src.Bounds())
	draw.Draw(temp, temp.Rect, src, src.Bounds().Min, draw.Src)
	return temp.Pix
}
func IsRootRGBA(src image.Image) *image.RGBA {
	if v, ok := src.(*image.RGBA); ok &&
		v.Stride/4 == v.Rect.Dx() &&
		len(v.Pix) / v.Stride == v.Rect.Dy() {
		return v
	}
	return nil
}