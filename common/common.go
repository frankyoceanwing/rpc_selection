package common

const ServerAddress = ":1234"

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith interface {
	Multiply(args *Args, reply *int) error
	Divide(args *Args, quo *Quotient) error
}
