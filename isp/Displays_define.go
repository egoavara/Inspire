package isp

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"image"
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
	Format      string           `json:"Format"`
	W           int32            `json:"Width"`
	H           int32            `json:"Height"`
	RefreshRate int32            `json:"RefreshRate"`
	origin      *sdl.DisplayMode `json:"-"`
}

func newDisplayMode(mode sdl.DisplayMode) displayMode {
	return displayMode{
		Format:      SDLPixelToString(mode.Format),
		W:           mode.W,
		H:           mode.H,
		RefreshRate: mode.RefreshRate,
		origin: &mode,
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
func (s *display) Create(work Worker, header *ApplicationHeader) *Application {
	if header == nil{
		header = &ApplicationHeader{}
	}
	header.Display = s
	return createApplication(work, header)
}
