package rpc

import (
	"log"
	"net"

	"github.com/ama29n/hex/internal/adapters/framework/left/grpc/pb"
	"github.com/ama29n/hex/internal/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (grpca Adapter) RUN() {
	var err error

	listen, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	arithmaticServiceServer := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterArithmaticServiceServer(grpcServer, arithmaticServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal("Failed to serve gRPC server over port :9000")
	}
}
