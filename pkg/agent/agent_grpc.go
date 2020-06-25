package agent

import "google.golang.org/grpc"

func (r *Agent) RegisterGRPCHandler(_ *grpc.Server) {
	//NO-OP for now
}

func (r *Agent) GetServerOpts() []grpc.ServerOption {
	return []grpc.ServerOption{}
}
