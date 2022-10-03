package main

import (
	"context"
	"fmt"
	"time"

	"github.com/valri11/testgrpsstream/comms"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("test client")

	host := "localhost"
	port := 8080
	addr := fmt.Sprintf("%s:%d", host, port)

	ctx := context.Background()
	ctxTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctxTimeout, addr,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
		return
	}
	c := comms.NewCounterClient(conn)

	req := comms.CounterReq{}
	countStream, err := c.Count(ctx, &req)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
		return
	}

	for {
		resp, err := countStream.Recv()
		if err != nil {
			fmt.Printf("ERR: %v\n", err)
			return
		}
		fmt.Printf("resp: %v\n", resp)
	}
}
