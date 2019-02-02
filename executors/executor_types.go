package executors

import (
	"github.com/yabslabs/runner/common"
	"github.com/yabslabs/runner/executors/kubernetes"
	"github.com/yabslabs/runner/runner"
)

var types = []Executor{&kubernetes.KubernetesExecutor{}}

type Executor interface {
	GetKind() string
	GetStatus() runner.RunnerStatus
	GetJobStatus() common.JobStatus
	Run(common.Job) common.JobStatus
}
