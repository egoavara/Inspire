package igl45

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Global int32

func (s Global) Mat4(mat4 mgl32.Mat4) {
	gl.UniformMatrix4fv(int32(s), 1, false, &mat4[0])
}
func (s Global) Mat3(mat3 mgl32.Mat3) {
	gl.UniformMatrix3fv(int32(s), 1, false, &mat3[0])
}
func (s Global) Mat2(mat2 mgl32.Mat2) {
	gl.UniformMatrix2fv(int32(s), 1, false, &mat2[0])
}
func (s Global) Vec4(vec4 mgl32.Vec4) {
	gl.Uniform4fv(int32(s), 1, &vec4[0])

}
func (s Global) Vec3(vec3 mgl32.Vec3) {
	gl.Uniform3fv(int32(s), 1, &vec3[0])
}
func (s Global) Vec2(vec2 mgl32.Vec2) {
	gl.Uniform2fv(int32(s), 1, &vec2[0])
}
func (s Global) Float64(f64 float64) {
	gl.Uniform1d(int32(s), f64)
}
func (s Global) Float32(f32 float32) {
	gl.Uniform1f(int32(s), f32)
}
func (s Global) Int32(i int32) {
	gl.Uniform1i(int32(s), i)
}
func (s Global) uInt32(ui uint32) {
	gl.Uniform1ui(int32(s), ui)
}
func (s Global) Float64s(f64s ...float64) {
	gl.Uniform1dv(int32(s), int32(len(f64s)), &f64s[0])
}
func (s Global) Float32s(f32s ...float32) {
	gl.Uniform1fv(int32(s), int32(len(f32s)), &f32s[0])
}
func (s Global) Int32s(is ...int32) {
	gl.Uniform1iv(int32(s), int32(len(is)), &is[0])
}
func (s Global) uInt32s(uis ...uint32) {
	gl.Uniform1uiv(int32(s), int32(len(uis)), &uis[0])
}