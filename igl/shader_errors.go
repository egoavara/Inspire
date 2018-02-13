package igl

import "github.com/pkg/errors"

var (
	ErrorCompileFail = errors.New("Comtilation Fail")
	ErrorBuildFail   =  errors.New("Building Fail")
)
var (
	CriticalProgramDeallocate = errors.New("Program deallocate")
)
