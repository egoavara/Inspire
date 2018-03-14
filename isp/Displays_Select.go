package isp

import (
	"image"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/iamGreedy/Inspire/utl"
	"fmt"
)


var Displays _Displays
func initDisplays()  {
	Displays = nil
	var err error
	var displ = utl.MustValue(sdl.GetNumVideoDisplays()).(int)
	var tempsdlrect sdl.Rect
	for i := 0; i < displ; i++ {
		model := utl.MustValue(sdl.GetNumDisplayModes(i)).(int)
		modes := make([]displayMode, model)
		disp := display{}
		for j := 0; j < model; j++{
			var temp sdl.DisplayMode
			utl.Must(sdl.GetDisplayMode(i, j, &temp))
			modes[j] = newDisplayMode(temp)
		}
		disp.Name = utl.MustValue(sdl.GetDisplayName(i)).(string)
		disp.DiagonalDPI, disp.VerticalDPI, disp.HorizontalDPI, err = sdl.GetDisplayDPI(i)
		utl.Must(err)
		utl.Must(sdl.GetDisplayBounds(i, &tempsdlrect))
		disp.Bound = ConvertRect(tempsdlrect)
		utl.Must(sdl.GetDisplayUsableBounds(i, &tempsdlrect))
		disp.UsableBound = ConvertRect(tempsdlrect)
		disp.DisplayModes = modes
		Displays = append(Displays, &disp)
	}
}
func ConvertRect(from sdl.Rect) (to image.Rectangle) {
	return image.Rect(int(from.X), int(from.Y), int(from.X + from.W), int(from.Y + from.H))
}

type _Displays []*display

func (s _Displays) GetByName(name string) *display {
	for _, v := range s{
		if v.Name == name{
			return v
		}
	}
	return nil
}
func (s _Displays) List() []string {
	var temp = make([]string, s.Length())
	for i, v := range s{
		temp[i] = v.Name
	}
	return temp
}
func (s _Displays) GetByIndex(idx int) *display {
	idx = s.indexize(idx)
	if idx < 0{
		return nil
	}
	return s[idx]
}
func (s _Displays) Length() int {
	return len(s)
}
func (s _Displays) indexize(i int) int {
	if i < 0{
		if -i <= len(s){
			return len(s) + i
		}
	}else{
		if i > len(s){
			return -1
		}
		return i
	}
	return i
}
func (s _Displays) String() string {
	switch len(s) {
	case 0:
		return fmt.Sprint("Displays(No Available)")
	case 1, 2, 3, 4, 5:
		var listr = ""
		for _, v := range s{
			listr += v.Name
			listr += ", "
		}
		listr = listr[:len(listr) - 2]
		return fmt.Sprintf("Displays(%s)", listr)
	}
	return fmt.Sprint("Displays(...)")
}


