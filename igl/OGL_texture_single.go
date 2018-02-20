package igl

import (
	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/iamGreedy/Inspire/bto"
	"image"
	"image/draw"
)

type SingleTexture struct {
	ptr uint32
}

func NewSingleTexture(w, h int) *SingleTexture {

	temp := &SingleTexture{}
	// Create texture
	gl.CreateTextures(gl.TEXTURE_2D, 1, &temp.ptr)
	// Texture size
	gl.TextureStorage2D(temp.ptr, 1, gl.RGBA8, int32(w), int32(h))
	// Texture parameter
	gl.TextureParameteri(temp.ptr, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TextureParameteri(temp.ptr, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TextureParameteri(temp.ptr, gl.TEXTURE_WRAP_S, gl.CLAMP_READ_COLOR)
	gl.TextureParameteri(temp.ptr, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	return temp
}

func (s *SingleTexture) Get() image.Image {
	var w, h = s.Size()
	var dst = image.NewRGBA(image.Rect(0, 0, int(w), int(h)))
	gl.GetTextureImage(s.ptr, 1, gl.RGBA, gl.UNSIGNED_BYTE, int32(len(dst.Pix)), gl.Ptr(dst.Pix))
	return dst
}
func (s *SingleTexture) Read(dst draw.Image) error {
	var w, h = s.Size()
	var bd = dst.Bounds()
	if bd.Dy() != int(h) || bd.Dx() != int(w) {
		return ErrorInvalidSize
	}
	if v := bto.IsRootRGBA(dst); v != nil {
		gl.GetTextureImage(s.ptr, 1, gl.RGBA, gl.UNSIGNED_BYTE, int32(len(v.Pix)), gl.Ptr(v.Pix))
		return nil
	}
	temp := s.Get()
	draw.Draw(dst, dst.Bounds().Intersect(temp.Bounds()), temp, temp.Bounds().Min, draw.Src)
	return nil
}
func (s *SingleTexture) Write(src image.Image) error {
	var w, h = s.Size()
	var bd = src.Bounds()
	if bd.Dy() != int(h) || bd.Dx() != int(w) {
		return ErrorInvalidSize
	}
	pix := bto.Image(src)
	gl.TextureSubImage2D(s.ptr, 1, 0, 0, w, h, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(pix))
	return nil
}
func (s *SingleTexture) Size() (w, h int32) {
	gl.GetTextureParameteriv(s.ptr, gl.TEXTURE_WIDTH, &w)
	gl.GetTextureParameteriv(s.ptr, gl.TEXTURE_HEIGHT, &h)
	return w, h
}
func (s *SingleTexture) Use(unit uint32) {
	gl.BindTextureUnit(unit, s.ptr)
}

