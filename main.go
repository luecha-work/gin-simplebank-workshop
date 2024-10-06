package main

import (
	"context"
	"database/sql"
	"os"

	"net"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"github.com/rakyll/statik/fs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/techschool/simple_bank/api"
	db "github.com/techschool/simple_bank/db/sqlc"
	_ "github.com/techschool/simple_bank/doc/statik"
	"github.com/techschool/simple_bank/gapi"
	"github.com/techschool/simple_bank/pb"
	"github.com/techschool/simple_bank/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal().Msg("cannot load config")
	}

	//TODO: Create Logger in gRPC server 1
	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal().Msg("cannot connect to db")
	}

	// run db migration
	runDBMigration(config.MigrationURL, config.DBSource)

	store := db.NewStore(conn)
	go runGatewayServer(config, store)
	runGrpcServer(config, store)
	// runGinServer(config, store)
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		// log.Fatal().Msg().Err(err).Msg("cannot create new migrate instance")
		log.Fatal().Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		// log.Fatal().Msg().Err(err).Msg("failed to run migrate up")
		log.Fatal().Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}

func runGrpcServer(
	config utils.Config,
	store db.Store,
) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal().Msg("cannot create server")
	}

	//TODO: Create Logger in gRPC server 2
	gprcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	grpcServer := grpc.NewServer(gprcLogger)
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Msg("cannot start listener")
	}

	log.Info().Msgf("start gRPC server %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Msg("cannot start gRPC server")
	}
}

func runGatewayServer(
	config utils.Config,
	store db.Store,
) {

	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal().Msg("cannot create server")
	}

	// TODO: Custom Responst json format
	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterSimpleBankHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Msg("cannot register handler server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	//TODO: Setup Swagger
	//TODO: Use statik import `_ "github.com/techschool/simple_bank/doc/statik"`
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal().Msg("cannot create statik fs")
	}

	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
	mux.Handle("/swagger/", swaggerHandler)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Msg("cannot start listener")
	}

	log.Info().Msgf("start HTTP gateway server %s", listener.Addr().String())
	// TODO: Create Logger in HTTP gateway server
	handler := gapi.HttpLogger(mux)
	err = http.Serve(listener, handler)
	if err != nil {
		log.Fatal().Msg("cannot start HTTP gateway server")
	}
}

func runGinServer(config utils.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Msg("cannot create server")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Msg("cannot start server")
	}
}
