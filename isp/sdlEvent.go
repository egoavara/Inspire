package isp

import (
	"github.com/veandco/go-sdl2/sdl"
	"sync"
	"fmt"
)

var (
	sdlEventChannelsLock = new(sync.Mutex)
	sdlEventChannels   = make(map[uint32]chan sdl.Event)
)

func installHandler(id uint32) chan sdl.Event {
	sdlEventChannelsLock.Lock()
	defer sdlEventChannelsLock.Unlock()
	fmt.Println("install")
	ch := make(chan sdl.Event, 32)
	sdlEventChannels[id] = ch
	return ch
}
func uninstallHandler(id uint32) {
	sdlEventChannelsLock.Lock()
	defer sdlEventChannelsLock.Unlock()
	//
	close(sdlEventChannels[id])
	delete(sdlEventChannels, id)
}
func broadcast(event sdl.Event) {
	for _, v := range sdlEventChannels {
		v <- event
	}
}
func send(to uint32, event sdl.Event) {

	if ch, ok := sdlEventChannels[to]; ok {
		ch <- event
	}
}

type cbfilter struct{}

func initSDLEventHandler() {
	sdl.AddEventWatchFunc(func(event sdl.Event, i interface{}) bool {
		sdlEventChannelsLock.Lock()
		defer sdlEventChannelsLock.Unlock()

		fmt.Println("watch", event)
		switch e := event.(type) {
		case *sdl.UserEvent:
			send(e.WindowID, e)
		case *sdl.QuitEvent:
			broadcast(e)
		case *sdl.OSEvent:
			broadcast(e)
		case *sdl.AudioDeviceEvent:
			broadcast(e)
		case *sdl.TouchFingerEvent:
			broadcast(e)
		case *sdl.MultiGestureEvent:
			broadcast(e)
		case *sdl.DollarGestureEvent:
			broadcast(e)
		case *sdl.DropEvent:
			send(e.WindowID, e)
		case *sdl.RenderEvent:
			broadcast(e)
		case *sdl.ClipboardEvent:
			broadcast(e)


		case *sdl.SysWMEvent:
			broadcast(e)

		case *sdl.ControllerAxisEvent:
			broadcast(e)
		case *sdl.ControllerButtonEvent:
			broadcast(e)
		case *sdl.ControllerDeviceEvent:
			broadcast(e)
			//
		case *sdl.JoyAxisEvent:
			broadcast(e)
		case *sdl.JoyBallEvent:
			broadcast(e)
		case *sdl.JoyHatEvent:
			broadcast(e)
		case *sdl.JoyButtonEvent:
			broadcast(e)
		case *sdl.JoyDeviceEvent:
			broadcast(e)

		case *sdl.KeyboardEvent:
			send(e.WindowID, e)
		case *sdl.MouseMotionEvent:
			send(e.WindowID, e)
		case *sdl.MouseWheelEvent:
			send(e.WindowID, e)
		case *sdl.MouseButtonEvent:
			send(e.WindowID, e)

		case *sdl.TextEditingEvent:
			send(e.WindowID, e)
		case *sdl.TextInputEvent:
			send(e.WindowID, e)

		}
		return true
	}, nil)
}
