package igl

import (
	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/pkg/errors"
	"sync/atomic"
)

type OGLShader struct {
	program uint32
}

func (s *OGLShader) Use() {
	if atomic.CompareAndSwapUint32(&s.program, 0, 0){
		panic(CriticalProgramDeallocate)
	}
	gl.UseProgram(s.program)
}

func (s *OGLShader) Release() {
	if atomic.CompareAndSwapUint32(&s.program, 0, 0){
		panic(CriticalProgramDeallocate)
	}
	gl.UseProgram(0)
}
func (s *OGLShader) Close() {
	temp := s.program
	atomic.StoreUint32(&s.program, 0)
	gl.DeleteProgram(temp)
}

