package runner

// RunnerStatus that represents the Status of a Runner
type RunnerStatus int

// RunnerStatus implementation as enum
const (
	UNKNOWN RunnerStatus = iota
	CREATED
	PENDING // or RUNNING?
	READY
	BUSY
	UNREACHABLE
	ERROR // or FAILED?
	STUCK
	PAUSED
	SUSPENDED
	// are these states of a runner or a job?
	// SUCCEEDED
	// CANCELED
	SKIPPED
	RETRIED
)
