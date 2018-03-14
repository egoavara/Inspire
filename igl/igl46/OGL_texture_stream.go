package igl46

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/iamGreedy/Inspire/bto"
	"image"
	"github.com/iamGreedy/Inspire/er"
)

type StreamTexture struct {
	tex        uint32
	pbo        [2]uint32
	writepoint int
	inited bool
}

func NewStreamTexture(w, h int) *StreamTexture {
	temp := &StreamTexture{}
	// Create texture
	gl.CreateTextures(gl.TEXTURE_2D, 1, &temp.tex)
	gl.CreateBuffers(2, &temp.pbo[0])
	// Texture parameter
	gl.TextureParameteri(temp.tex, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TextureParameteri(temp.tex, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TextureParameteri(temp.tex, gl.TEXTURE_WRAP_S, gl.CLAMP_READ_COLOR)
	gl.TextureParameteri(temp.tex, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	// buffer
	gl.NamedBufferData(temp.pbo[0], w * h * 4, gl.PtrOffset(0), gl.STREAM_DRAW)
	gl.NamedBufferData(temp.pbo[1], w * h * 4, gl.PtrOffset(0), gl.STREAM_DRAW)
	return temp
}
func (s *StreamTexture) Skip() {
	s.writepoint = (s.writepoint + 1) % 2
}
func (s *StreamTexture) Write(src image.Image) error {
	var w, h = s.Size()
	var bd = src.Bounds()
	var dst  = bto.IsRootRGBA(src)
	if bd.Dy() != int(h) || bd.Dx() != int(w) || dst == nil{
		return er.ErrorInvalidSize
	}
	//
	current := s.writepoint
	next := (s.writepoint + 1) % 2
	defer func() {
		s.writepoint = next
	}()
	//
	gl.BindBuffer(gl.PIXEL_UNPACK_BUFFER, s.pbo[current])
	defer gl.BindBuffer(gl.PIXEL_UNPACK_BUFFER, 0)
	if !s.inited{
		gl.NamedBufferData(s.pbo[current], len(dst.Pix), gl.Ptr(dst.Pix), gl.STREAM_DRAW)
	}
	gl.TextureSubImage2D(s.tex, 1, 0, 0, w, h, gl.RGBA, gl.UNSIGNED_BYTE, gl.PtrOffset(0))
	gl.NamedBufferData(s.pbo[next], len(dst.Pix), gl.Ptr(dst.Pix), gl.STREAM_DRAW)
	return nil
}
func (s *StreamTexture) Size() (w, h int32) {
	gl.GetTextureParameteriv(s.tex, gl.TEXTURE_WIDTH, &w)
	gl.GetTextureParameteriv(s.tex, gl.TEXTURE_HEIGHT, &h)
	return w, h
}
func (s *StreamTexture) Use(unit uint32) {
	gl.BindTextureUnit(unit, s.tex)
}

