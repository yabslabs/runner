package runner

// RunnerStatus that represents the Status of a Runner
type RunnerStatus int

// RunnerStatus implementation as enum
const (
	UNKNOWN RunnerStatus = iota
	ACTIVE
	PAUSED
	ONLINE
	OFFLINE
)
