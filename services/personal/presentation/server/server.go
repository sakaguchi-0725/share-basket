package server

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	pb "share-basket/personal/proto/gen"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
	addr string
	*grpc.Server
}

func New(addr string) *Server {
	return &Server{
		addr:   addr,
		Server: grpc.NewServer(),
	}
}

func (s *Server) MapServices(services Services) {
	pb.RegisterAuthServiceServer(s.Server, services.AuthService)
	pb.RegisterPersonalShoppingServiceServer(s.Server, services.ShoppingService)
}

func (s *Server) Run() {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		log.Println("gRPC server is running on", s.addr)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	done := make(chan struct{})
	go func() {
		s.GracefulStop()
		close(done)
	}()

	select {
	case <-ctx.Done():
		log.Println("shutdown timeout exceeded")
		s.Stop()
	case <-done:
		log.Println("gRPC server shutdown completed")
	}
}
