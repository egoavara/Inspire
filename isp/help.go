package isp

import (
	"github.com/iamGreedy/Inspire/igl/iglcommons"
	"github.com/veandco/go-sdl2/sdl"
)

func helpIGLProfile(profile iglcommons.Profile) int {
	switch profile {
	case iglcommons.CORE:
		return sdl.GL_CONTEXT_PROFILE_CORE
	case iglcommons.COMPATIBLE:
		return sdl.GL_CONTEXT_PROFILE_COMPATIBILITY
	}
	return 0
}