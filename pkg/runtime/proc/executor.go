package proc

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"syscall"

	"github.com/pkg/errors"

	"github.com/scoir/canis/pkg/datastore"
	"github.com/scoir/canis/pkg/runtime"
)

const (
	StewardPIDFile = "%s/steward_pid"
	StewardConfig  = "%s/steward_config.yaml"
	StewardOut     = "%s/steward_out"
	StewardErr     = "%s/steward_err"
	StewardExec    = "%s/steward"
)

type Config struct {
	ExecPath string `mapstructure:"path"`
	HomeDir  string `mapstructure:"home"`
}

type Executor struct {
	path, home string
	steward    *os.Process
}

func New(conf *Config) runtime.Executor {
	r := &Executor{
		path: conf.ExecPath,
		home: conf.HomeDir,
	}

	if r.home == "" {
		r.home = "/tmp"
	}

	return r
}

func (r *Executor) LaunchSteward(conf []byte) (string, error) {
	pfilename := fmt.Sprintf(StewardPIDFile, r.home)
	p, err := ioutil.ReadFile(pfilename)
	if err == nil {
		pid, _ := strconv.ParseInt(string(p), 10, 64)
		if isProcAlive(pid) {
			return string(p), errors.New("steward is already running")
		} else {
			_ = os.Remove(pfilename)
		}
	}

	stewardConfigFile := fmt.Sprintf(StewardConfig, r.home)
	err = ioutil.WriteFile(stewardConfigFile, conf, 0644)
	if err != nil {
		return "", errors.Wrap(err, "unable to write config for steward")
	}

	n := os.ExpandEnv(fmt.Sprintf(StewardExec, r.path))
	cmd := exec.Command(n, "start", "--config", stewardConfigFile)
	cmd.Stdout, err = os.Create(fmt.Sprintf(StewardOut, r.home))
	if err != nil {
		return "", errors.Wrap(err, "unable to create std output for steward")
	}
	cmd.Stderr, err = os.Create(fmt.Sprintf(StewardErr, r.home))
	if err != nil {
		return "", errors.Wrap(err, "unable to create std error for steward")
	}
	err = cmd.Start()
	if err != nil {
		return "", err
	}

	pid := strconv.Itoa(cmd.Process.Pid)
	err = ioutil.WriteFile(pfilename, []byte(pid), 0644)
	if err != nil {
		return "", errors.Wrap(err, "error writing steward pid file.")
	}

	return pid, nil
}

func (r *Executor) LaunchAgent(agent *datastore.Agent) (string, error) {
	panic("implement me")
}

func (r *Executor) AgentStatus(pID string) (runtime.Process, error) {
	panic("implement me")
}

func (r *Executor) ShutdownAgent(pID string) error {
	panic("implement me")
}

func (r *Executor) Watch(pID string) (runtime.Watcher, error) {
	panic("implement me")
}

func (r *Executor) StreamLogs(pID string) (io.ReadCloser, error) {
	panic("implement me")
}

func (r *Executor) PS() {
	panic("implement me")
}

func (r *Executor) Describe() {
	panic("implement me")
}

func isProcAlive(pid int64) bool {
	process, err := os.FindProcess(int(pid))
	if err != nil {
		return false
	} else {
		err := process.Signal(syscall.Signal(0))
		return err == nil
	}
}
