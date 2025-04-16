package main

import (
	"io"
	"net"

	"github.com/HMasataka/bigrpc/bi/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBidirectionServer
}

func (s *server) Bidirection(stream pb.Bidirection_BidirectionServer) error {
	for {
		in, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		stream.Send(&pb.Response{Res: in.Data})
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:31080")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterBidirectionServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}
