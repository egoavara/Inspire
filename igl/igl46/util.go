package igl46

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"fmt"
)

func CSTR(str string) *uint8 {
	return gl.Str(str + "\x00")
}
func GLVersion() (Major, Minor, Patch int) {
	fmt.Println(gl.GetString(gl.VERSION))
	return 0,0,0
}