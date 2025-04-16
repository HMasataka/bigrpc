package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/HMasataka/bigrpc/pb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func getConnection() (*grpc.ClientConn, error) {
	port := os.Getenv("PORT")
	address := fmt.Sprintf("127.0.0.1:%v", port)

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	fmt.Println("Connecting to", address)

	return conn, nil
}

func main() {
	conn, err := getConnection()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ctx := context.Background()

	client := pb.NewServerStreamClient(conn)

	data := pb.Data{Data: "connect"}

	stream, err := client.ServerStream(ctx, &data)
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
