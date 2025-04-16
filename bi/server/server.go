package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/HMasataka/bigrpc/bi/pb"
	"github.com/joho/godotenv"
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

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func getListener() (net.Listener, error) {
	port := os.Getenv("PORT")
	address := fmt.Sprintf("127.0.0.1:%v", port)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	fmt.Println("Listening on", address)

	return listener, nil
}

func main() {
	listener, err := getListener()
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterBidirectionServer(grpcServer, &server{})

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
