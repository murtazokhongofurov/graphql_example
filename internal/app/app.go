package app

import (
	"fmt"
	"log"
	"net"

	"github.com/graphql_example/pkg/config"
	"github.com/graphql_example/pkg/logger"
	"google.golang.org/grpc"
	"github.com/graphql_example/pkg/db"
	"google.golang.org/grpc/reflection"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.LogLevel)

	// postgres://user:password@host:5432/database
	pgxURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase)

	pg, err := db.New(pgxURL, db.MaxPoolSize(cfg.PGXPoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()
	groupService := service.NewGroupService(l, pg)

	lis, err := net.Listen("tcp", ":"+cfg.ProductServicePort)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - grpcclient.New: %w", err))
	}
	c := grpc.NewServer()
	reflection.Register(c)
	group.RegisterGroupServiceServer(c, groupService)

	l.Info("Server is running on" + "port" + ": " + cfg.ProductServicePort)

	if err := c.Serve(lis); err != nil {
		log.Fatal("Error while listening: ", err)
	}
}
