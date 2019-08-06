package main

import (
	"context"
	"fmt"
	"io"
	"sync"

	pb "github.com/sylba2050/bigrpc/bi/proto"
	"google.golang.org/grpc"
)

const port = "localhost:31080"

func main() {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewBidirectionClient(conn)
	stream, err := client.Bidirection(context.Background())
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer stream.CloseSend()

		for i := 1; i < 10; i++ {
			ijk := fmt.Sprintf("%d", i)
			stream.Send(&pb.Data{Data: ijk})
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
