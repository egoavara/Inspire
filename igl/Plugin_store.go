package igl

import (
	"sync"
	"github.com/iamGreedy/Inspire/er"
)

//
var (
	ctx Context
	fninit func() error
	mtx sync.Mutex
)
type glversion struct {
	Major, Minor int
}
func PLGContext(context Context, init func() error)  {
	mtx.Lock()
	defer mtx.Unlock()
	if ctx != nil{
		panic(er.CriticalViolation)
	}
	ctx = context
	fninit = init
}
func LoadContext() Context {
	return ctx
}
//
