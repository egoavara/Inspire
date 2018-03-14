package igl46

import (
	"github.com/iamGreedy/Inspire/igl"
	"github.com/go-gl/gl/v4.6-core/gl"
)

func init() {
	igl.PLGContext(Context{}, glInit)
}
func glInit() error {
	if err := gl.Init(); err != nil{
		return err
	}
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0.0, 0.0, 1.0, 1.0)
	return nil
}
