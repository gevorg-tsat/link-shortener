package main

import (
	"context"
	"errors"
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
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

const shutdownTimeout = 5 * time.Second

func main() {
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
		linkStorage, err = storage.NewDB(cfg)
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

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("grcp server listening to %v\n", lis.Addr())
		if err = s.Serve(lis); err != nil {
			log.Fatalf("failed to start: %v", err)
		}
	}()
	go func() {
		log.Printf("http server listening to %v\n", httpServer.S.Addr)
		if err = httpServer.S.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("http listen and serve returned err: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("got interruption signal")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := httpServer.S.Shutdown(shutdownCtx); err != nil {
		log.Printf("http server shutdown returned an err: %v\n", err)
	}
	log.Println("http server is stopped gracefully")
	s.GracefulStop()
	log.Println("grpc server is stopped gracefully")
	linkStorage.Shutdown()
	log.Println("storage connection is stopped")
}
