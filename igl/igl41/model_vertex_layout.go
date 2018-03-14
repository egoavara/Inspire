package igl46

import "github.com/go-gl/gl/v4.1-core/gl"

//
type VertexData uint8
const (
	// XYZ
	POSITION VertexData = iota
	// UV
	TEXTURE VertexData = iota
	// NORMAL
	NORMAL VertexData = iota
	// SKELLETON

)
func (s VertexData) Type() int {
	switch s {
	default:
		return 0
	case POSITION:
		return gl.FLOAT
	case NORMAL:
		return gl.FLOAT
	case TEXTURE:
		return gl.FLOAT
	}
}
func (s VertexData) TypeByteLength() int {
	switch s {
	default:
		return 0
	case POSITION:
		return 4
	case NORMAL:
		return 4
	case TEXTURE:
		return 4
	}
}
func (s VertexData) TypeCount() int {
	switch s {
	default:
		return 0
	case POSITION:
		return 3
	case NORMAL:
		return 3
	case TEXTURE:
		return 2
	}
}
func (s VertexData) TotalSize() int {
	return s.TypeByteLength() * s.TypeCount()
}

type VertexLayout []VertexData
func (s VertexLayout) Stride() int {
	res := 0
	for _, v:= range s {
		res += v.TypeCount() * v.TypeByteLength()
	}
	return res
}
func (s VertexLayout) Offset(start int) int {
	res := 0
	for i := 0; i < start; i ++{
		res += s[i].TypeCount() * s[i].TypeByteLength()
	}
	return res
}
func BuildVertexLayout(vds ... VertexData) VertexLayout {
	var temp []VertexData
	for _, v := range vds{
		temp = append(temp, v)
	}
	return VertexLayout(temp)
}