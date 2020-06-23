package framework

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/hyperledger/aries-framework-go/pkg/client/didexchange"
	"github.com/hyperledger/aries-framework-go/pkg/client/issuecredential"
	"github.com/hyperledger/aries-framework-go/pkg/client/route"
	"github.com/hyperledger/aries-framework-go/pkg/framework/context"
	"github.com/spf13/viper"

	"github.com/scoir/canis/pkg/datastore"
)

var configFileName string

func init() {
	flag.StringVar(&configFileName, "config", "agent-config", "override default config file name")

}

type Config struct {
	LedgerURL     string   `yaml:"ledgerURL"redis:"-"`
	ID            string   `yaml:"ID"`
	SCID          string   `yaml:"SCID"`
	Name          string   `yaml:"Name"`
	DID           string   `redis:"DID"`
	Verkey        string   `redis:"Verkey"`
	Confirmed     bool     `redis:"Confirmed"`
	KubeConfig    string   `yaml:"kubeConfig"`
	Namespace     string   `yaml:"namespace"`
	FQDN          string   `yaml:"FQDN"`
	ImageRegistry string   `yaml:"imageRegistry"`
	DockerTag     string   `yaml:"dockerTag"`
	DBPath        string   `yaml:"dbpath"`
	Endpoint      string   `yaml:"endpoint"`
	VDR           Endpoint `yaml:"vdr"`
	WSInbound     Endpoint `yaml:"wsinbound" redis:"-"`
	GRPC          Endpoint `yaml:"grpc" redis:"-"`
	Agency        Endpoint `yaml:"agency" redis:"-"`
	Steward       Endpoint `yaml:"steward" redis:"-"`
	GRPCBridge    Endpoint `yaml:"grpcBridge" redis:"-"`

	Database string `yaml:"database"`
	Mongo    *Mongo `yaml:"mongo"`

	lock sync.Mutex

	ds      datastore.Store
	ctx     *context.Provider
	didcl   *didexchange.Client
	credcl  *issuecredential.Client
	routecl *route.Client
}

type Mongo struct {
	URL      string `yaml:"url"`
	Database string `yaml:"database"`
}

type Endpoint struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (r Endpoint) Address() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}

//NewFileConfig constructs the agent config from filesystem
//noinspection GoUnusedExportedFunction
func NewFileConfig() (*Config, error) {
	flag.Parse()
	viper.SetConfigName(configFileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/canis/")
	//maybe
	viper.AddConfigPath("./config/proc/")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalln("Fatal error config file", err)
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatalln("failed to unmarshal config", err)
	}

	return config, nil
}
