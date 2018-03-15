package igl

import "github.com/iamGreedy/Inspire/igl/iglcommons"

type Context interface {
	Version() (major, minor int)
	VersionString() (string)
	Clear()
	Profile() iglcommons.Profile
} 
