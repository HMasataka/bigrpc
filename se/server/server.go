package main

import (
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
		ijk := fmt.Sprintf("%d", i)
		res := &pb.Response{Res: ijk}
		err := stream.Send(res)
		if err == io.EOF {
			return nil
		}
		if err != nil {
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
	grpcServer.Serve(lis)
}
