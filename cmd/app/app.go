package main

import (
	"flag"
	"fmt"
	"github.com/gevorg-tsat/link-shortener/config"
	"github.com/gevorg-tsat/link-shortener/internal/grcpserver"
	"github.com/gevorg-tsat/link-shortener/internal/httpserver"
	pb "github.com/gevorg-tsat/link-shortener/internal/shortener_v1"
	"github.com/gevorg-tsat/link-shortener/internal/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)

	storageType := flag.String("storage", "in-memory", "type of storage that will be used. Available: in-memory, postgres")
	flag.Parse()
	if *storageType != "postgres" && *storageType != "in-memory" {
		log.Fatal("invalid storage type. Available: in-memory, postgres")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	var linkStorage storage.Storage
	if *storageType == "postgres" {
		linkStorage, err = storage.NewDB(
			cfg.DB.Host,
			cfg.DB.User,
			cfg.DB.Password,
			cfg.DB.Name,
			cfg.DB.Port,
		)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		linkStorage = storage.NewInMemory()
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%d", cfg.GRCP.Host, cfg.GRCP.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	shortenerService := grcpserver.New(linkStorage, cfg)
	pb.RegisterShortenerV1Server(s, shortenerService)
	httpServer := httpserver.New(shortenerService, cfg)

	go func() {
		log.Printf("grcp server listening to %v\n", lis.Addr())
		if err = s.Serve(lis); err != nil {
			log.Fatalf("failed to start: %v", err)
		}
		wg.Done()
	}()

	go func() {
		log.Printf("http server listening to %v\n", httpServer.S.Addr)
		if err = httpServer.S.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
		wg.Done()
	}()

	wg.Wait()
}
