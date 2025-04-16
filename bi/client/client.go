package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/HMasataka/bigrpc/bi/pb"
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

	client := pb.NewBidirectionClient(conn)

	stream, err := client.Bidirection(ctx)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			if err := stream.CloseSend(); err != nil {
				panic(err)
			}
		}()

		for i := 1; i < 10; i++ {
			if err := stream.Send(&pb.Data{Data: fmt.Sprintf("%d", i)}); err != nil {
				panic(err)
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

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
	}()

	wg.Wait()
}
