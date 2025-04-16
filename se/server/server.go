package main

import (
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/HMasataka/bigrpc/se/pb"
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

func main() {
	lis, err := net.Listen("tcp", "localhost:31080")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterSestreamServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
