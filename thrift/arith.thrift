namespace go tf

struct Args {
  1: i32 A,
  2: i32 B
}

struct Reply {
  1: i32 R
}

struct Quotient {
  1: i32 Quo,
  2: i32 Rem
}

service Arith {
  Reply Multiply(1:Args arg),
  Quotient Divide(1:Args arg)
}