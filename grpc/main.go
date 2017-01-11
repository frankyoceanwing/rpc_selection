package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/frankyoceanwing/rpc_selection/common"
	"github.com/frankyoceanwing/rpc_selection/grpc/pb"
	"github.com/frankyoceanwing/rpc_selection/grpc/server"
)

func main() {
	l, e := net.Listen("tcp", common.ServerAddress)
	if e != nil {
		log.Fatalf("failed to listen: %v", e)
	}
	s := grpc.NewServer()
	pb.RegisterArithServer(s, new(server.Arith))
	s.Serve(l)
}
