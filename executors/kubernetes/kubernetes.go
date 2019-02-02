package kubernetes

import (
	"github.com/yabslabs/runner/common"
	"github.com/yabslabs/runner/runner"
)

func (e *KubernetesExecutor) GetKind() string {
	return ""
}

func (e *KubernetesExecutor) GetStatus() runner.RunnerStatus {
	return 1
}

func (e *KubernetesExecutor) GetJobStatus() common.JobStatus {
	return 1
}

func (e *KubernetesExecutor) Run(common.Job) common.JobStatus {
	return 1
}

type KubernetesExecutor struct{}
