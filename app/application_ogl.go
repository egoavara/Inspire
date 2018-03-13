package app

import (
	"github.com/iamGreedy/Inspire/er"
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type Application struct {
	Header ApplicationHeader
	//
	wnd *sdl.Window
	ctx sdl.GLContext
}

func NewApplication(header *ApplicationHeader) *Application {
	return &Application{
		Header: *header,
	}
}

type ApplicationHeader struct {
	app           *Application
	Name          string
	Width, Height int
	SwapInterval  int
	Borderless    bool
	Top           bool
	Fullscreen    bool
}

func (s ApplicationHeader) sdlFlags() (flag uint32) {
	flag = sdl.WINDOW_ALLOW_HIGHDPI | sdl.WINDOW_OPENGL
	if s.Borderless {
		flag |= sdl.WINDOW_BORDERLESS
	}
	if s.Top {
		flag |= sdl.WINDOW_ALWAYS_ON_TOP
	}
	if s.Fullscreen {
		flag |= sdl.WINDOW_FULLSCREEN
	}
	return flag
}
func (s ApplicationHeader) Commit() error {
	s.app.wnd.Set
}

func (s *Application) Init() (err error) {
	defer func() {
		if err != nil {
			s.Clear()
		}
	}()
	s.wnd, err = sdl.CreateWindow(s.Header.Name, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, int32(s.Header.Width), int32(s.Header.Height), s.Header.sdlFlags())
	if err != nil {
		return errors.WithMessage(er.ErrorInitialization, err.Error())
	}
	s.ctx, err = sdl.GLCreateContext(s.wnd)
	if err != nil {
		return errors.WithMessage(er.ErrorInitialization, err.Error())
	}
	sdl.GLMakeCurrent(s.wnd, s.ctx)
	sdl.GLSetSwapInterval(s.Header.SwapInterval)
	return nil
}
func (s *Application) Clear() {
}
func (s *Application) Swap() {
	sdl.GLSwapWindow(s.wnd)
}
