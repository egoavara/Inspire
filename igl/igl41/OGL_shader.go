package igl46

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"sync/atomic"
	"sync"
	"github.com/iamGreedy/Inspire/er"
)

type Shader struct {
	program uint32
	layout VertexLayout
	mtx *sync.RWMutex
}

func (s *Shader) Use() {
	s.mtx.RLock()
	if s.program == 0{
		panic(er.CriticalProgramDeallocate)
	}
	gl.UseProgram(s.program)
}
func (s *Shader) Release() {
	s.mtx.RUnlock()
	if s.program == 0{
		panic(er.CriticalProgramDeallocate)
	}
	gl.UseProgram(0)
}
func (s *Shader) Close() {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if s.program == 0{
		panic(er.CriticalProgramDeallocate)
	}
	temp := s.program
	atomic.StoreUint32(&s.program, 0)
	gl.DeleteProgram(temp)
}
//func (s *Shader) Layout(layout VertexLayout)  {
//	s.mtx.Lock()
//	defer s.mtx.Unlock()
//	//
//	stride := layout.Stride()
//	s.layout = layout
//	for i, v := range layout{
//		gl.EnableVertexAttribArray(uint32(i))
//		gl.VertexAttribPointer(uint32(i), int32(v.TypeCount()), uint32(v.Type()), false, int32(stride), gl.PtrOffset(layout.Offset(i)))
//	}
//}
func (s *Shader) Global(name string) Global {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	//
	loc := gl.GetUniformLocation(s.program, CSTR(name))
	if loc < 0{
		return -1
	}
	return Global(loc)
}
//func (s *Shader) Local(name string) Local{
//
//}
func (s *Shader) Fragment(name string)  {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	gl.BindFragDataLocation(s.program, 0, CSTR(name))

}

func Exist(global Global) bool {
	return global > 0
}