package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/valri11/testgrpsstream/comms"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/stats"
)

type counterServer struct {
	comms.UnimplementedCounterServer
}

func main() {
	fmt.Println("test Counter server stream")

	srv := grpc.NewServer(
		grpc.StatsHandler(&Handler{}),
		grpc.KeepaliveParams(keepalive.ServerParameters{Time: 2 * time.Second, Timeout: 2 * time.Second}),
	)

	var counterSrv counterServer
	comms.RegisterCounterServer(srv, &counterSrv)

	port := 8080
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
		return
	}

	err = srv.Serve(lis)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
		return
	}
}

func (s *counterServer) Count(req *comms.CounterReq, stream comms.Counter_CountServer) error {

	log.Println("counter started")

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	count := int32(0)

	for {
		select {
		case <-stream.Context().Done():
			log.Println("stream context done")
			return stream.Context().Err()
		case <-ticker.C:
			count++
			resp := comms.CounterResp{Counter: count}
			err := stream.Send(&resp)
			if err != nil {
				log.Printf("Send err: %v", err)
				return err
			}
		}
	}
}

type Handler struct {
}

func (h *Handler) TagRPC(context.Context, *stats.RPCTagInfo) context.Context {
	log.Println("TagRPC")
	return context.Background()
}

func (h *Handler) HandleRPC(context.Context, stats.RPCStats) {
	log.Println("HandleRPC")
}

func (h *Handler) TagConn(context.Context, *stats.ConnTagInfo) context.Context {
	log.Println("Tag Conn")
	return context.Background()
}

func (h *Handler) HandleConn(c context.Context, s stats.ConnStats) {
	log.Printf("HandleConn: %v\n", s)
	switch s.(type) {
	case *stats.ConnEnd:
		log.Println("get grpc ConnEnd")
	}
}
