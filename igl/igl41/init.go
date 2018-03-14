package igl46

import (
	"github.com/iamGreedy/Inspire/igl"
	"github.com/go-gl/gl/v4.1-core/gl"
)

func init() {
	//fmt.Println("v4.1")
	igl.PLGContext(Context{}, glInit)
}
func glInit() error {
	if err := gl.Init(); err != nil{
		return err
	}
	//fmt.Println(gl.GoStr(gl.GetString(gl.VERSION)))
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0.0, 0.0, 1.0, 1.0)
	return nil
}
