package framework

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/scoir/canis/pkg/steward/api"
)

func (r *Config) GetStewardClient() (steward.AdminClient, error) {
	cc, err := grpc.Dial(r.Steward.Address(), grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial grpc for steward client")
	}
	cl := steward.NewAdminClient(cc)
	return cl, nil
}

//func (r *Config) GetAgencyClient() (agency.AgencyClient, error) {
//	cc, err := grpc.Dial(r.Agency.Address(), grpc.WithInsecure())
//	if err != nil {
//		return nil, errors.Wrap(err, "failed to dial grpc for routing client")
//	}
//	cl := agency.NewAgencyClient(cc)
//	return cl, nil
//}
