package common

// JobStatus that represents the Status of a Job
type JobStatus int

// JobStatus implementation as enum
const (
	UNKNOWN JobStatus = iota
	CREATED
	PENDING
	RUNNING
	FAILED
	SUCCESS
	CANCELED
	SKIPPED
	MANUAL
	STUCK
	RETRIED
	PAUSED
	SUSPENDED
)

type Job struct {
}
