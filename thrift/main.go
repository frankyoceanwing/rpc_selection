package main

import (
	"log"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/frankyoceanwing/rpc_selection/common"
	"github.com/frankyoceanwing/rpc_selection/thrift/gen-go/tf"
	"github.com/frankyoceanwing/rpc_selection/thrift/server"
)

func main() {
	var trans thrift.TServerTransport
	var e error
	trans, e = thrift.NewTServerSocket(common.ServerAddress)
	if e != nil {
		log.Fatal("creating socket error:", e)
	}
	processor := tf.NewArithProcessor(new(server.Arith))
	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	server := thrift.NewTSimpleServer4(processor, trans, transportFactory, protocolFactory)
	e = server.Serve()
	if e != nil {
		log.Fatal("serve error:", e)
	}
}
