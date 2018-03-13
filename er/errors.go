package er

import "github.com/pkg/errors"

// Error.* is return value
var (
	ErrorCompileFail = errors.New("Comtilation Fail")
	ErrorBuildFail   = errors.New("Building Fail")
	ErrorInvalidSize = errors.New("Invalid TypeCount")
	ErrorInitialization = errors.New("Initialization Fail")
)

