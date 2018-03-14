package igl

import (
	"github.com/iamGreedy/Inspire/utl"
	"sync"
)

var one = new(sync.Once)
func Init() {
	one.Do(func() {
		utl.Must(fninit())
	})
}
