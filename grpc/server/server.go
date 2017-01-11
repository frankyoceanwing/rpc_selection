package server

import (
	"errors"

	"golang.org/x/net/context"

	"github.com/frankyoceanwing/rpc_selection/grpc/pb"
)

type Arith int

func (t *Arith) Multiply(ctx context.Context, args *pb.Args) (*pb.Reply, error) {
	return &pb.Reply{args.A * args.B}, nil
}

func (t *Arith) Divide(ctx context.Context, args *pb.Args) (*pb.Quotient, error) {
	if args.B == 0 {
		return nil, errors.New("divide by zero")
	}

	return &pb.Quotient{Quo: args.A / args.B, Rem: args.A % args.B}, nil
}
