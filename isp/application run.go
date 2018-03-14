package isp

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/iamGreedy/Inspire/igl"
)

type Worker interface {
	Before(dt int64) error
	While(app *Application, ctx igl.Context) error
	After(app *Application) error
	Handle(event sdl.Event) error
	Recover(err error) error
}

type Working struct {
	FnBefore func(dt int64) error
	FnWhile  func(app *Application, ctx igl.Context) error
	FnRecover func(err error) error
	FnHandle func(event sdl.Event) error
	FnAfter  func(app *Application) error
}


func (s *Working) Handle(event sdl.Event) error {
	if s.FnBefore == nil{
		return nil
	}
	return s.FnHandle(event)
}

func (s *Working) Before(dt int64) error {
	if s.FnBefore == nil{
		return nil
	}
	return s.FnBefore(dt)
}
func (s *Working) While(app *Application, ctx igl.Context) error {
	if s.FnWhile == nil{
		return nil
	}
	return s.FnWhile(app, ctx)
}
func (s *Working) After(app *Application) error {
	if s.FnAfter == nil{
		return nil
	}
	return s.FnAfter(app)
}
func (s *Working) Recover(err error) error {
	if s.FnRecover == nil{
		return err
	}
	return s.FnRecover(err)
}
