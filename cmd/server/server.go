package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/HMasataka/bigrpc/pb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type stream struct {
	pb.UnimplementedServerStreamServer
}

func (s *stream) ServerStream(data *pb.Data, stream pb.ServerStream_ServerStreamServer) error {
	fmt.Println(data.Data)

	for i := 1; i < 10; i++ {
		res := &pb.Response{Res: fmt.Sprintf("%d", i)}

		if err := stream.Send(res); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
	}

	return nil
}

type bidirection struct {
	pb.UnimplementedBidirectionServer
}

func (b *bidirection) Bidirection(stream pb.Bidirection_BidirectionServer) error {
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

	pb.RegisterServerStreamServer(grpcServer, &stream{})
	pb.RegisterBidirectionServer(grpcServer, &bidirection{})

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
