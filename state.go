package optional

import "fmt"

type state uint8

const (
	stateMissing state = iota
	stateNull
	statePresent
)

var _ fmt.Stringer = state(0)

func (s state) String() string {
	switch s {
	case stateMissing:
		return "missing"
	case stateNull:
		return "null"
	case statePresent:
		return "present"
	default:
		return "unknown"
	}
}
