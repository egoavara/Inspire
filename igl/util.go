package igl

import "github.com/go-gl/gl/v4.5-core/gl"

func CSTR(str string) *uint8 {
	return gl.Str(str + "\x00")
}