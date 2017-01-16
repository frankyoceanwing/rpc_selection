package server

import (
	"errors"

	"github.com/frankyoceanwing/rpc_selection/thrift/gen-go/tf"
)

type Arith int

func (t *Arith) Multiply(args *tf.Args_) (*tf.Reply, error) {
	return &tf.Reply{args.A * args.B}, nil
}

func (t *Arith) Divide(args *tf.Args_) (*tf.Quotient, error) {
	if args.B == 0 {
		return nil, errors.New("divide by zero")
	}

	return &tf.Quotient{Quo: args.A / args.B, Rem: args.A % args.B}, nil
}
