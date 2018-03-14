package isp

import "github.com/veandco/go-sdl2/sdl"

type ApplicationHeader struct {
	app           *Application
	Display       *display
	Mode          displayMode
	Name          string
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
func (s ApplicationHeader) Commit() (err error) {
	// TODO : Top

	// Border
	s.app.wnd.SetBordered(s.Borderless)

	// Swap Interval
	sdl.GLMakeCurrent(s.app.wnd, s.app.ctx)
	sdl.GLSetSwapInterval(s.SwapInterval)

	// FullScreen
	if s.Fullscreen {
		err = s.app.wnd.SetFullscreen(sdl.WINDOW_FULLSCREEN)
	} else {
		err = s.app.wnd.SetFullscreen(0)
	}
	if err != nil {
		return err
	}

	// Mode
	s.app.wnd.SetDisplayMode(s.Mode.origin)
	s.app.wnd.SetPosition(sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED)
	// Name
	s.app.wnd.SetTitle(s.Name)
	return nil
}
