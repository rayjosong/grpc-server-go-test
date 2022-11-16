package main

import (
	"bookshop/server/bookshop/pb"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Business problem :
// 	- We have a bookstore
// 	- We want to expose its inventory via an RPC function

type server struct {
	pb.UnimplementedInventoryServer
}

// can be called by clients. This method handles that request
// the context gives us access to auth tokens, heeaders, etc
func (s *server) GetBookList(ctx context.Context, in *pb.GetBookListRequest) (*pb.GetBookListResponse, error) {
	return &pb.GetBookListResponse{
		Books: getSampleBooks(),
	}, nil
}

func main() {
	// Listening on TPC port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	// Initialise new gRPC server instance
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterInventoryServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getSampleBooks() []*pb.Book {
	sampleBooks := []*pb.Book{
		{
			Title:     "The Hithiker's Guide to the galaxy",
			Author:    "Douglas Adams",
			PageCount: 42,
		},
		{
			Title:     "The Lord of the Rings",
			Author:    "J.R.R Tolkien",
			PageCount: 1234,
		},
	}
	return sampleBooks
}
