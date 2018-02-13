package igl

import (
	"github.com/go-gl/gl/v4.5-core/gl"
	"sync/atomic"
	"sync"
)

type GLShader struct {
	program uint32
	layout VertexLayout
	mtx *sync.RWMutex
}

func (s *GLShader) Use() {
	s.mtx.RLock()
	if s.program == 0{
		panic(CriticalProgramDeallocate)
	}
	gl.UseProgram(s.program)
}
func (s *GLShader) Release() {
	s.mtx.RUnlock()
	if s.program == 0{
		panic(CriticalProgramDeallocate)
	}
	gl.UseProgram(0)
}
func (s *GLShader) Close() {

	s.mtx.Lock()
	defer s.mtx.Unlock()
	if s.program == 0{
		panic(CriticalProgramDeallocate)
	}
	temp := s.program
	atomic.StoreUint32(&s.program, 0)
	gl.DeleteProgram(temp)
}
//
func (s *GLShader) Layout(layout VertexLayout)  {
	stride := layout.Stride()
	s.layout = layout
	for i, v := range layout{
		gl.EnableVertexAttribArray(uint32(i))
		gl.VertexAttribPointer(uint32(i), int32(v.Size()), uint32(v.Type()), false, int32(stride), gl.PtrOffset(layout.Offset(i)))
	}
}
