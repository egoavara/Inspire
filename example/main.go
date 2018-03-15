package main

import (
	"github.com/iamGreedy/Inspire/isp"
	"github.com/iamGreedy/Inspire/utl"
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
	"github.com/iamGreedy/Inspire/er"
	"github.com/iamGreedy/Inspire/igl"
	_ "github.com/iamGreedy/Inspire/igl/igl41"
	"runtime"
)

func main() {
	runtime.LockOSThread()
	runtime.GOMAXPROCS(runtime.NumCPU())
	//
	var disp = isp.Displays.GetByIndex(0)
	var app = disp.Create(nil, &isp.ApplicationHeader{
		Name: "Hello, Inspire!",
		Mode: disp.DisplayModes[len(disp.DisplayModes) - 1],
	})
	utl.Must(app.Init())
	app.Work = &isp.Working{
		FnBefore: func(dt int64) error {
			return nil
		},
		FnHandle: func(event sdl.Event) error {
			fmt.Println("Get", event)
			switch event.(type) {
			case *sdl.QuitEvent:
				return er.NotifyExit
			default:
				fmt.Println(event)
			}
			return nil
		},
		FnWhile: func(app *isp.Application, ctx igl.Context) error {
			ctx.Clear()
			sdl.Delay(5)
			return nil
		},
	}

	isp.Main()
}
