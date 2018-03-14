package isp

import (
	"github.com/iamGreedy/Inspire/er"
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
	"runtime"
	"github.com/iamGreedy/Inspire/utl"
	"sync/atomic"
	"time"
)

type Application struct {
	Header ApplicationHeader
	Work Worker
	Ticker Ticker
	//
	runflag int32
	wnd *sdl.Window
	ctx sdl.GLContext
}


func createApplication(work Worker, header *ApplicationHeader) *Application {
	temp := &Application{
		Work:work,
	}
	header.app = temp
	temp.Header = *header
	runtime.SetFinalizer(temp, finalizeApplication)
	return temp
}
func finalizeApplication(s *Application)  {
	if s.wnd != nil{
		sdl.GLDeleteContext(s.ctx)
		utl.Must(s.wnd.Destroy())
	}
}

// SDL, GL functions
func (s *Application) Init() (err error) {
	s.wnd, err = sdl.CreateWindow(s.Header.Name, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, int32(s.Header.Mode.W), int32(s.Header.Mode.H), s.Header.sdlFlags())
	if err != nil {
		return errors.WithMessage(er.ErrorInitialization, err.Error())
	}
	s.ctx, err = sdl.GLCreateContext(s.wnd)
	if err != nil {
		return errors.WithMessage(er.ErrorInitialization, err.Error())
	}
	sdl.GLMakeCurrent(s.wnd, s.ctx)
	return s.Header.Commit()
}
func (s *Application) Swap() {
	sdl.GLSwapWindow(s.wnd)
}
func (s *Application) Run() (err error){
	atomic.StoreInt32(&s.runflag, 1)
	defer atomic.StoreInt32(&s.runflag, 0)
	s.Ticker.Start()
	defer s.Ticker.Stop()

	var prev, curr time.Time
	prev = s.Ticker.Wait()
	for curr = s.Ticker.Wait();err == nil;curr = s.Ticker.Wait(){
		dt := int64(prev.Sub(curr)/time.Millisecond)
		err := s.Work.Work(s, dt)
		if err != nil {
			break
		}

		// endup
		s.Swap()
		prev = curr
	}
	return nil
}
func (s *Application) Stop() error{
	for true{
		err := s.Work.Work(s)
		if err != nil {
			return err
		}
		s.Swap()
	}
	return nil
}

func (s *Application ) IsRun() bool {
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
func (s *Application) SetGrab(isgrab bool)  {
	s.wnd.SetGrab(isgrab)
}
