package main

import (
	"context"
	"fmt"
	"io"

	pb "github.com/sylba2050/bigrpc/se/proto"
	"google.golang.org/grpc"
)

const port = "localhost:31080"

func main() {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewSestreamClient(conn)
	data := pb.Data{Data: "connect"}
	stream, err := client.Sestream(context.Background(), &data)
	if err != nil {
		panic(err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(resp.Res)
	}
}
