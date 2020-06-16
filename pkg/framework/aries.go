package framework

import (
	"log"

	"github.com/hyperledger/aries-framework-go/pkg/client/didexchange"
	"github.com/hyperledger/aries-framework-go/pkg/client/issuecredential"
	"github.com/hyperledger/aries-framework-go/pkg/client/route"
	"github.com/hyperledger/aries-framework-go/pkg/didcomm/messaging/msghandler"
	"github.com/hyperledger/aries-framework-go/pkg/didcomm/transport/ws"
	"github.com/hyperledger/aries-framework-go/pkg/framework/aries"
	"github.com/hyperledger/aries-framework-go/pkg/framework/aries/api"
	"github.com/hyperledger/aries-framework-go/pkg/framework/aries/defaults"
	"github.com/hyperledger/aries-framework-go/pkg/framework/context"
	"github.com/hyperledger/aries-framework-go/pkg/kms/legacykms"
	"github.com/hyperledger/aries-framework-go/pkg/storage"
	"github.com/hyperledger/aries-framework-go/pkg/storage/leveldb"
	indiana "github.com/hyperledger/aries-framework-go/pkg/vdri/indy"
	"github.com/pkg/errors"
)

type provider struct {
	sp  storage.Provider
	kms *legacykms.BaseKMS
}

func newProvider(dbpath string) *provider {
	r := &provider{
		sp: leveldb.NewProvider(dbpath),
	}

	r.kms, _ = legacykms.New(r)
	return r
}

func (r *provider) StorageProvider() storage.Provider {
	return r.sp
}

func (r *provider) createKMS(_ api.Provider) (api.CloseableKMS, error) {
	return r.kms, nil
}

func (r *Config) GetAriesContext() *context.Provider {
	if r.ctx == nil {
		err := r.createAriesContext()
		if err != nil {
			log.Fatalln("failed to create aries context", err)
		}

	}
	return r.ctx
}

func (r *Config) createAriesContext() error {
	wsinbound := r.WSInbound.Address()

	p := newProvider(r.DBPath)
	jones, err := indiana.New("scoir", r.LedgerURL, p.kms)
	if err != nil {
		return errors.Wrap(err, "unable to initialize indiana jones")
	}

	log.Printf("creating for %s\n", r.DBPath)
	framework, err := aries.New(
		aries.WithMessageServiceProvider(msghandler.NewRegistrar()),
		aries.WithStoreProvider(p.StorageProvider()),
		aries.WithLegacyKMS(p.createKMS),
		defaults.WithInboundWSAddr(wsinbound, wsinbound),
		aries.WithOutboundTransports(ws.NewOutbound()),
		aries.WithVDRI(jones),
		aries.WithServiceEndpoint(r.Endpoint),
	)
	if err != nil {
		return errors.Wrapf(err, "failed to start aries agent rest on port [%s], failed to initialize framework",
			wsinbound)
	}

	ctx, err := framework.Context()
	if err != nil {
		return errors.Wrapf(err, "failed to start aries agent rest on port [%s], failed to get aries context",
			wsinbound)
	}
	r.ctx = ctx

	return nil
}

func (r *Config) GetDIDClient() (*didexchange.Client, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.didcl != nil {
		return r.didcl, nil
	}

	didcl, err := didexchange.New(r.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error creating did client")
	}

	r.didcl = didcl
	return r.didcl, nil
}

func (r *Config) GetCredentialClient() (*issuecredential.Client, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.credcl != nil {
		return r.credcl, nil
	}

	credcl, err := issuecredential.New(r.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error creating credential client")
	}
	r.credcl = credcl
	return r.credcl, nil
}

func (r *Config) GetRouterClient() (*route.Client, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.routecl != nil {
		return r.routecl, nil
	}

	routecl, err := route.New(r.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create route client for college: %v\n")
	}
	r.routecl = routecl
	return r.routecl, nil
}
