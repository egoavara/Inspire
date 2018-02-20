package igl

import "github.com/pkg/errors"

// Error.* is return value
var (
	ErrorCompileFail = errors.New("Comtilation Fail")
	ErrorBuildFail   = errors.New("Building Fail")
	ErrorInvalidSize = errors.New("Invalid TypeCount")
	ErrorInit = errors.New("Application Init Error")
)

// Critical only use for panic()
var (
	CriticalProgramDeallocate = errors.New("Program deallocate")
)
