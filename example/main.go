package main

import (
	"github.com/iamGreedy/Inspire/isp"
	"github.com/iamGreedy/Inspire/utl"
	"time"
)

func main() {
	var disp = isp.Displays.GetByIndex(0)
	var app = disp.Create(nil, &isp.ApplicationHeader{
		Name: "Hello, Inspire!",
		Mode: disp.DisplayModes[len(disp.DisplayModes) - 1],
	})

	//fmt.Println(app)
	utl.Must(app.Init())
	for true{
		time.Sleep(50 * time.Millisecond)
		app.Swap()
	}
}
