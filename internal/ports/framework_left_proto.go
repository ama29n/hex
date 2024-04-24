package ports

import (
	"context"

	"github.com/ama29n/hex/internal/adapters/framework/left/grpc/pb"
)

// left side framework port for gRPC
// on driving side, adapter depends upon port which is implemented by application service

type GRPCPort interface {
	RUN()
	GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
	GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
	GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
	GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
}
