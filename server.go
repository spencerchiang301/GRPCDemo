package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "grpcgps/proto"
	"grpcgps/utility"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server portPort ")
)

type server struct {
	pb.UnimplementedMyGpsServer
}

func (s *server) GetMyAddress(ctx context.Context, in *pb.Address) (*pb.MyPoint, error) {
	log.Printf("Received: %v", in.GetAddr())
	lat, lng, err := utility.GetLatLongFromAddress(in.GetAddr())
	if err != nil {
		println(err)
	}
	myPoint := &pb.MyPoint{
		Addr: in.GetAddr(),
		Location: &pb.Point{
			P1: lat,
			P2: lng,
		},
	}
	return myPoint, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMyGpsServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
