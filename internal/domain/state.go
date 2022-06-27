package domain

// State of a character in a word
type State int64

const (
	// NotPresent state
	NotPresent State = iota
	// Present state
	Present
	// InPlace state
	InPlace
)

func (s State) String() string {
	switch s {
	case NotPresent:
		return "not present"
	case Present:
		return "present"
	case InPlace:
		return "in place"
	}
	return "unknown"
}
