package kubernetes

import (
	"fmt"
	"io"
	"os"

	"github.com/yabslabs/runner/common"
	"github.com/yabslabs/runner/runner"
	"sigs.k8s.io/controller-runtime/pkg/client/config"

	"github.com/sirupsen/logrus"
	api "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	restclient "k8s.io/client-go/rest"
)

// RemoteExecutor defines the interface accepted by the Exec command - provided for test stubbing
type RemoteExecutor interface {
	Execute(method string, url *url.URL, config *restclient.Config, stdin io.Reader, stdout, stderr io.Writer, tty bool) error
}

// DefaultRemoteExecutor is the standard implementation of remote command execution
type DefaultRemoteExecutor struct{}

func (*DefaultRemoteExecutor) Execute(method string, url *url.URL, config *restclient.Config, stdin io.Reader, stdout, stderr io.Writer, tty bool) error {
	exec, err := remotecommand.NewSPDYExecutor(config, method, url)
	if err != nil {
		return err
	}

	return exec.Stream(remotecommand.StreamOptions{
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
		Tty:    tty,
	})
}


func (e *KubernetesExecutor) GetKind() string {
	return "Kubernetes Executor"
}

func (e *KubernetesExecutor) GetStatus() runner.RunnerStatus {
	return 1
}

func (e *KubernetesExecutor) GetJobStatus() common.JobStatus {
	return 1
}

func (e *KubernetesExecutor) Run(job common.Job) common.JobStatus {
	pod, err := e.Client.CoreV1().Pods(e.Namespace).Get(e.PodName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	if pod.Status.Phase != api.PodRunning {
		return fmt.Errorf("Pod '%s' (on namespace '%s') is not running and cannot execute commands; current phase is '%s'",
			e.PodName, e.Namespace, pod.Status.Phase)
	}

	containerName := e.ContainerName
	if len(containerName) == 0 {
		logrus.Infof("defaulting container name to '%s'", pod.Spec.Containers[0].Name)
		containerName = pod.Spec.Containers[0].Name
	}

	// TODO: refactor with terminal helpers from the edit utility once that is merged
	var stdin io.Reader
	if e.Stdin {
		stdin = e.In
	}

	// TODO: consider abstracting into a client invocation or client helper
	req := e.Client.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(pod.Name).
		Namespace(pod.Namespace).
		SubResource("exec").
		Param("container", containerName)
	req.VersionedParams(&api.PodExecOptions{
		Container: containerName,
		Command:   e.Command,
		Stdin:     stdin != nil,
		Stdout:    e.Out != nil,
		Stderr:    e.Err != nil,
	}, scheme.ParameterCodec)

	return e.Executor.Execute("POST", req.URL(), e.Config, stdin, e.Out, e.Err, false)

}

type KubernetesExecutor struct {
	Namespace     string
	PodName       string
	ContainerName string
	Stdin         bool
	Command       []string

	In  io.Reader
	Out io.Writer
	Err io.Writer

	Executor RemoteExecutor
	Client   *kubernetes.Clientset
	Config   *restclient.Config
}

func CreateKubernetesExecutor()(*KubernetesExecutor, error){
	cfg, err := config.GetConfig()
		if err != nil {
			log.Error(err, "unable to set up client config")
			return nil, err
		}
	hostname, err := os.Hostname()
	if err != nil {
		log.Error(err, "unable to get hostname")
		return nil, err
	}
	clientset, err := GetClientSet(cfg)
	if err != nil {
		log.Error(err, "unable to get clientset")
		return nil, err
	}
	return KubernetesExecutor{
		PodName: hostname,
		Client: clientset,
		Config: cfg,
	}, nil
}


func GetClientSet(config *rest.Config) (*kubernetes.Clientset, error) {
	return kubernetes.NewForConfig(config)
}