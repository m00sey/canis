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
	LaunchAgent(agent *datastore.Agent) (string, error)
	AgentStatus(pID string) (Process, error)
	ShutdownAgent(pID string) error
	WatchAgent(pID string) (Watcher, error)
	StreamAgentLogs(pID string) (io.ReadCloser, error)
}
