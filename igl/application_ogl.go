package igl

import (
	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type Application struct {
	Header ApplicationHeader
	//
	wnd           *sdl.Window
	ctx           sdl.GLContext
	currentShader Shader
}
type ApplicationHeader struct {
	Name          string
	Width, Height int
	SwapInterval  int
	Borderless    bool
	Top           bool
	Fullscreen    bool
}
func (s *Application) Init() (err error) {
	var flag uint32 = sdl.WINDOW_ALLOW_HIGHDPI | sdl.WINDOW_OPENGL
	if s.Header.Borderless {
		flag |= sdl.WINDOW_BORDERLESS
	}
	if s.Header.Top {
		flag |= sdl.WINDOW_ALWAYS_ON_TOP
	}
	if s.Header.Fullscreen {
		flag |= sdl.WINDOW_FULLSCREEN
	}
	s.wnd, err = sdl.CreateWindow(s.Header.Name, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, int32(s.Header.Width), int32(s.Header.Height), flag)
	if err != nil {
		return errors.WithMessage(ErrorInit, err.Error())
	}
	s.ctx, err = sdl.GLCreateContext(s.wnd)
	if err != nil {
		return errors.WithMessage(ErrorInit, err.Error())
	}
	sdl.GLMakeCurrent(s.wnd, s.ctx)
	sdl.GLSetSwapInterval(s.Header.SwapInterval)
	//
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	return nil
}
func (s *Application) Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)
}
func (s *Application) Swap() {
	sdl.GLSwapWindow(s.wnd)
}
