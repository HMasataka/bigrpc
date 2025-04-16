package main

import (
	"errors"
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
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		if err := stream.Send(&pb.Response{Res: in.Data}); err != nil {
			return err
		}
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

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
