package igl45

import "github.com/veandco/go-sdl2/sdl"


func init() {
	for _, ch := range []uint32{sdl.INIT_TIMER , sdl.INIT_AUDIO , sdl.INIT_VIDEO , sdl.INIT_EVENTS , sdl.INIT_JOYSTICK , sdl.INIT_HAPTIC , sdl.INIT_GAMECONTROLLER}{
		if sdl.WasInit(ch) != ch{
			sdl.Init(ch)
		}
	}

}