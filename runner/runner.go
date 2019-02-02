package runner

// RunnerStatus that represents the Status of a Runner
type RunnerStatus int

// RunnerStatus implementation as enum
const (
	UNKNOWN RunnerStatus = iota
	READY
	BUSY
	UNREACHABLE
	ERROR
	STUCK
	PAUSED
	SUSPENDED
)
