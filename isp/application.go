package isp

import (
	"fmt"
	"github.com/iamGreedy/Inspire/er"
	"github.com/iamGreedy/Inspire/igl"
	"github.com/iamGreedy/Inspire/utl"
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
	"runtime"
	"sync/atomic"
)

type Application struct {
	Header ApplicationHeader
	Work   Worker
	//
	wnd *sdl.Window
	ctx sdl.GLContext
}

func createApplication(work Worker, header *ApplicationHeader) *Application {
	temp := &Application{
		Work: work,
	}
	header.app = temp
	temp.Header = *header
	runtime.SetFinalizer(temp, finalizeApplication)
	return temp
}
func finalizeApplication(s *Application) {
	if s.wnd != nil {
		utl.Must(s.wnd.Destroy())
	}
}

// SDL, GL functions
func (s *Application) Init() (err error) {
	defer func() {
		if err != nil {
			if s.wnd != nil {
				s.wnd.Destroy()
			}
		}
	}()
	s.wnd, err = sdl.CreateWindow(s.Header.Name, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, int32(s.Header.Mode.W), int32(s.Header.Mode.H), s.Header.sdlFlags())
	if err != nil {
		return errors.WithMessage(er.ErrorInitialization, err.Error())
	}
	s.ctx, err = sdl.GLCreateContext(s.wnd)
	if err != nil {
		return errors.WithMessage(er.ErrorInitialization, err.Error())
	}
	sdl.GLMakeCurrent(s.wnd, s.ctx)
	sdl.GLSetSwapInterval(0)
	//
	igl.Init()
	ctx := igl.LoadContext()
	major, minor := ctx.Version()
	fmt.Println(ctx.VersionString())
	//
	if err = sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, major); err != nil {
		return err
	}
	if err = sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, minor); err != nil {
		return err
	}
	if err = sdl.GLSetAttribute(sdl.GL_CONTEXT_FORWARD_COMPATIBLE_FLAG, 1); err != nil {
		return err
	}
	if err = sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, helpIGLProfile(ctx.Profile())); err != nil {
		return err
	}
	if err = sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1); err != nil {
		return err
	}
	if err = s.wnd.SetDisplayMode(s.Header.Mode.origin); err != nil {
		return err
	}

	return nil
}
func (s *Application) Swap() {
	sdl.GLSwapWindow(s.wnd)
}
func (s *Application) Run() error {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent:
			atomic.StoreInt32(&running, 0)
		default:
		}
	}
	var err error
	// Before
	err = s.Work.Before(dt)
	if err != nil {
		err = s.Work.Recover(err)
		if err == nil {
			return true
		}
		return false
	}
	// While
	err = s.Work.While(s, ctx)
	if err != nil {
		err = s.Work.Recover(err)
		if err == nil {
			return true
		}
		return false
	}
	// Swapping
	s.Swap()
	// Event handle

	for _, e := range evs {
		err = s.Work.Handle(e)
		if err != nil {
			err = s.Work.Recover(err)
			if err == nil {
				return true
			}
			return false
		}
	}
	// After
	err = s.Work.After(s)
	if err != nil {
		err = s.Work.Recover(err)
		if err == nil {
			return true
		}
		return false
	}
	return true
}

// Window functions
func (s *Application) Restore() {
	s.wnd.Restore()
}
func (s *Application) Raise() {
	s.wnd.Raise()
}
func (s *Application) Hide() {
	s.wnd.Hide()
}
func (s *Application) Show() {
	s.wnd.Show()
}

// Mouse Functions
func (s *Application) GetGrab() bool {
	return s.wnd.GetGrab()
}
func (s *Application) SetGrab(isgrab bool) {
	s.wnd.SetGrab(isgrab)
}
