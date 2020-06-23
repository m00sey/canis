package datastore

type SchemaList struct {
	Count  int
	Schema []*Schema
}

type Schema struct {
	ID         string
	Name       string
	Version    string
	Attributes []*Attribute
}

type Attribute struct {
	Name string
	Type int32
}

type AgentList struct {
	Count  int
	Agents []*Agent
}

type StatusType string

var (
	Starting   StatusType = "STARTING"
	Running    StatusType = "RUNNING"
	Error      StatusType = "ERROR"
	Terminated StatusType = "TERMINATED"
)

type Agent struct {
	ID                  string
	Name                string
	AssignedSchemaId    string
	ConnectionID        string
	ConnectionState     string
	DID                 string
	EndorsableSchemaIds []string
	Status              StatusType
	PID                 string
}

type SchemaCriteria struct {
	Start, PageSize int
	Name            string
}

type AgentCriteria struct {
	Start, PageSize int
	Name            string
}

type DIDCriteria struct {
	Start, PageSize int
}

type DID struct {
	DID, Verkey, Endpoint string
	Public                bool
}

type DIDList struct {
	Count int
	DIDs  []*DID
}
