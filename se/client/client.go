package main

import (
	"context"
	"fmt"
	"io"

	"github.com/HMasataka/bigrpc/se/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = "localhost:31080"

func main() {
	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ctx := context.Background()

	client := pb.NewSestreamClient(conn)

	data := pb.Data{Data: "connect"}

	stream, err := client.Sestream(ctx, &data)
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
