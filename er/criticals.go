package er

import "github.com/pkg/errors"

// Critical only use for panic()
var (
	CriticalProgramDeallocate = errors.New("Program deallocate")
	CriticalViolation  = errors.New("Rule violation")
)
