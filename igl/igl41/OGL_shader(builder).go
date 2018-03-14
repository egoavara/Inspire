package igl46

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/pkg/errors"
	"github.com/iamGreedy/Inspire/er"
)

type BuilderShader struct {
	FlagmentSoucrce string
	VertexSoucrce   string
}
func (s *BuilderShader) Build() (*Shader, error) {
	// Create Program
	var program = gl.CreateProgram()
	// compile vertex shader
	var vs, vserr = ShaderSource(s.FlagmentSoucrce).Compile(gl.FRAGMENT_SHADER)
	if vserr != nil {
		return nil, vserr
	}
	defer gl.DeleteShader(vs)
	// compile fragment shader
	var fs, fserr = ShaderSource(s.FlagmentSoucrce).Compile(gl.FRAGMENT_SHADER)
	if fserr != nil {
		return nil, fserr
	}
	defer gl.DeleteShader(fs)
	// Linking
	gl.AttachShader(program, fs)
	gl.AttachShader(program, vs)
	gl.LinkProgram(program)
	// Error check
	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)
		log := make([]byte, logLength+1)
		gl.GetProgramInfoLog(program, logLength, nil, (*uint8)(&log[0]))
		return nil, errors.WithMessage(er.ErrorBuildFail, string(log))
	}
	//
	return &Shader{
		program: program,
	}, nil
}

type ShaderSource string
func (s ShaderSource) Compile(shaderType uint32) (shader uint32, err error) {
	shader = gl.CreateShader(shaderType)
	defer func() {
		if err != nil {
			gl.DeleteShader(shader)
			shader = 0
		}
	}()
	var cstr = gl.Str(string(s) + "\x00")
	//
	gl.ShaderSource(shader, 1, &cstr, nil)
	gl.CompileShader(shader)
	//
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
		var log []byte
		log = make([]byte, logLength+1)
		gl.GetShaderInfoLog(shader, logLength, nil, (*uint8)(gl.Ptr(&log[0])))
		err = errors.WithMessage(er.ErrorCompileFail, string(log[:len(log)-1]))
		return
	}
	return
}
