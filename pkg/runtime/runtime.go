package runtime

import (
	"io"
	"time"

	"github.com/scoir/canis/pkg/datastore"
)

type Watcher interface {
	Stop()
	ResultChan() <-chan AgentEvent
}
type Process interface {
	Status() datastore.StatusType
	Exited() bool
	Time() time.Duration
	Tail() []byte
}

type AgentEvent struct {
	RuntimeContext Process
}

//go:generate mockery -name=Executor
type Executor interface {
	LaunchSteward([]byte) (string, error)
	LaunchAgent(agent *datastore.Agent) (string, error)
	AgentStatus(pID string) (Process, error)
	ShutdownAgent(pID string) error

	Watch(pID string) (Watcher, error)
	StreamLogs(pID string) (io.ReadCloser, error)
	PS()
	Describe()
}
