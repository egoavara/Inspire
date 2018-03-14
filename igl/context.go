package igl

import "github.com/iamGreedy/Inspire/igl/iglcommons"

type Context interface {
	Version() (major, minor int)
	Clear()
	Profile() iglcommons.Profile
} 
