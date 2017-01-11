package main

import (
	"log"

	"github.com/frankyoceanwing/rpc_selection/common"
	"github.com/frankyoceanwing/rpc_selection/rpcx/server"
	"github.com/smallnest/rpcx"
)

func main() {
	s := rpcx.NewServer()
	s.RegisterName("Arith", new(server.Arith))
	e := s.Serve("tcp", common.ServerAddress)
	if e != nil {
		log.Fatal("serve error:", e)
	}
}
