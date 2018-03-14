package isp

import (
	"github.com/iamGreedy/Inspire/er"
	"github.com/iamGreedy/Inspire/igl"
	"github.com/iamGreedy/Inspire/utl"
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
	"runtime"
	"sync/atomic"
	"time"
	"fmt"
)

type Application struct {
	Header ApplicationHeader
	Work   Worker
	//
	runflag int32
	wnd     *sdl.Window
	ctx     sdl.GLContext
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
		if err != nil{
			if s.wnd != nil{
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
	fmt.Println(major, minor)
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

	return s.Header.Commit()
}
func (s *Application) Swap() {
	sdl.GLSwapWindow(s.wnd)
}
func (s *Application) Run() (err error) {
	if s.Work == nil {
		return errors.WithMessage(er.ErrorUnsatisfied, "s.Work == nil")
	}
	// runflag Set
	atomic.StoreInt32(&s.runflag, 1)
	// ticket setup
	var ticker Ticker
	if s.Header.SwapInterval == 0 {
		ticker = NewTicker(0)
	} else {
		ticker = NewTicker(time.Second / time.Duration(s.Header.SwapInterval))
	}
	ticker.Start()
	defer ticker.Stop()
	// Event
	//fmt.Println(utl.MustValue(s.wnd.GetID()).(uint32))
	//id := utl.MustValue(s.wnd.GetID()).(uint32)
	//evch := installHandler(id)
	//defer uninstallHandler(id)
	// context
	var ctx = igl.LoadContext()
	// timer values
	var prev, curr time.Time
	prev = ticker.Wait()
	// Main loop
MainLoop:
	for curr = ticker.Wait(); !atomic.CompareAndSwapInt32(&s.runflag, 0, 0); curr = ticker.Wait() {
		// Before
		err = s.Work.Before(int64(curr.Sub(prev) / time.Millisecond))
		if err != nil {
			err = s.Work.Recover(err)
			if err == nil {
				continue MainLoop
			}
			break MainLoop
		}
		// While
		err = s.Work.While(s, ctx)
		if err != nil {
			err = s.Work.Recover(err)
			if err == nil {
				continue MainLoop
			}
			break MainLoop
		}
		// Swapping
		s.Swap()
		// Event handle
		//ChannelLoop:
		//for {
		//	select {
		//	case e := <-evch:
		//		err = s.Work.Handle(e)
		//		if err != nil {
		//			err = s.Work.Recover(err)
		//			if err == nil {
		//				continue MainLoop
		//			}
		//			break MainLoop
		//		}
		//	//default:
		//	//	break ChannelLoop
		//	}
		//}
		// After
		err = s.Work.After(s)
		if err != nil {
			err = s.Work.Recover(err)
			if err == nil {
				continue MainLoop
			}
			break MainLoop
		}
		// endup
		prev = curr
	}
	return err
}
func (s *Application) Stop() {
	atomic.StoreInt32(&s.runflag, 0)
}

func (s *Application) IsRun() bool {
	return atomic.CompareAndSwapInt32(&s.runflag, 1, 1)
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
