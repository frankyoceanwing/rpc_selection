package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/frankyoceanwing/rpc_selection/common"
	"github.com/frankyoceanwing/rpc_selection/standard/server"
)

func main() {
	rpc.Register(new(server.Arith))
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", common.ServerAddress)
	if e != nil {
		log.Fatal("listen error:", e)
	}

	e = http.Serve(l, nil)
	if e != nil {
		log.Fatal("serve error:", e)
	}
}
