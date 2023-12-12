package main

import (
	"fmt"
	"github.com/gevorg-tsat/link-shortener/config"
	"github.com/gevorg-tsat/link-shortener/internal/server"
	pb "github.com/gevorg-tsat/link-shortener/internal/shortener_v1"
	"github.com/gevorg-tsat/link-shortener/internal/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	dbhost   = "localhost"
	dbport   = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
	appHost  = "localhost"
	appPort  = 8080
)

func main() {
	db := storage.NewInMemory()
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%d", appHost, appPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	cfg := config.Config{}
	cfg.Port = appPort
	cfg.Host = appHost
	pb.RegisterShortenerV1Server(s, server.New(db, cfg))

	log.Printf("server listening to %v\n", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to start: %v", err)
	}
}
