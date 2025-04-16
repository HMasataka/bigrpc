package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/HMasataka/bigrpc/se/pb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSestreamServer
}

func (s *server) Sestream(data *pb.Data, stream pb.Sestream_SestreamServer) error {
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

	pb.RegisterSestreamServer(grpcServer, &server{})

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
