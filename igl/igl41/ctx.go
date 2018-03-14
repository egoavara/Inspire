package igl46

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/iamGreedy/Inspire/igl/iglcommons"
)

type Context struct {}

func (s Context) Version() (major, minor int) {
	var ma, mi int32
	gl.GetIntegerv(gl.MAJOR_VERSION, &ma)
	gl.GetIntegerv(gl.MINOR_VERSION, &mi)
	return int(ma), int(mi)
}
func (s Context) Profile() iglcommons.Profile {
	return iglcommons.CORE
}
func (s Context) Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}
