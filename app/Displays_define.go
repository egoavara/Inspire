package app

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"image"
	"unsafe"
)

type display struct {
	index int `json:"-"`
	//
	Name          string          `json:"Name"`
	Bound         image.Rectangle `json:"Bound"`
	UsableBound   image.Rectangle `json:"UsableBound"`
	DiagonalDPI   float32         `json:"DiagonalDPI"`
	VerticalDPI   float32         `json:"VerticalDPI"`
	HorizontalDPI float32         `json:"HorizontalDPI"`
	DisplayModes  []displayMode   `json:"DisplayModes"`
}
//
type displayMode struct {
	Format      string         `json:"Format"`
	W           int32          `json:"Width"`
	H           int32          `json:"Height"`
	RefreshRate int32          `json:"RefreshRate"`
	DriverData  unsafe.Pointer `json:"-"`
}
func newDisplayMode(mode sdl.DisplayMode) displayMode {
	return displayMode{
		Format:      SDLPixelToString(mode.Format),
		W:           mode.W,
		H:           mode.H,
		RefreshRate: mode.RefreshRate,
		DriverData:  mode.DriverData,
	}
}
func (s displayMode) convert() sdl.DisplayMode {
	return sdl.DisplayMode{
		Format:      SDLPixelFromString(s.Format),
		W:           s.W,
		H:           s.H,
		RefreshRate: s.RefreshRate,
		DriverData:  s.DriverData,
	}
}
//
const displayStringerFormat = `Display(name:"%s", bound:"%s", DPI:%2.5f, Mode:<%d>")`
func (s display) String() string {
	return fmt.Sprintf(
		displayStringerFormat,
		s.Name,
		s.Bound,
		s.DiagonalDPI,
		len(s.DisplayModes),
	)
}
func (s display) CreateWindow() string {
	return fmt.Sprintf(
		displayStringerFormat,
		s.Name,
		s.Bound,
		s.DiagonalDPI,
		len(s.DisplayModes),
	)
}


