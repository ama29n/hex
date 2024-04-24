package main

import (
	"log"
	"os"

	"github.com/ama29n/hex/internal/adapters/app/api"
	"github.com/ama29n/hex/internal/adapters/core/arithmatic"
	db "github.com/ama29n/hex/internal/adapters/framework/right"
	"github.com/ama29n/hex/internal/ports"

	grpc "github.com/ama29n/hex/internal/adapters/framework/left/grpc"
)

func main() {
	var err error

	// ports
	var dbAdapter ports.DBPort
	var coreAdapter ports.ArithmaticPort
	var appAdapter ports.APIPort
	var grpcAdapter ports.GRPCPort

	// env
	dbDriver := os.Getenv("DB_DRIVER")
	dsName := os.Getenv("DS_NAME")

	dbAdapter, err = db.NewAdapter(dbDriver, dsName)

	if err != nil {
		log.Fatalf("DB Connection Failure %v", err)
	}

	defer dbAdapter.CloseDBConnection()

	coreAdapter = arithmatic.NewAdapter()

	appAdapter = api.NewAdapter(dbAdapter, coreAdapter)

	grpcAdapter = grpc.NewAdapter(appAdapter)

	grpcAdapter.RUN()

}
