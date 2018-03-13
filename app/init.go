package app

import (
	"github.com/iamGreedy/Inspire/utl"
	"github.com/veandco/go-sdl2/sdl"
)

func init() {
	utl.Must(sdl.Init(sdl.INIT_EVERYTHING ^ sdl.WasInit(sdl.INIT_EVERYTHING)))
	initDisplays()
}
