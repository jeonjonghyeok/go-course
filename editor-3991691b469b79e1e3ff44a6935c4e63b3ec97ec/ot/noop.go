package ot

import "io"

// NoOp empty operation
type NoOp struct{}

// InputLength required input length
func (n NoOp) InputLength() int {
	return 0
}

// OutputLength output string length
func (n NoOp) OutputLength() int {
	return 0
}

// Slice slice the operation
func (n NoOp) Slice(start, end int) PrimitiveOp {
	return NoOp{}
}

// Apply executes the operation
func (n NoOp) Apply(io.Reader, io.Writer) error {
	return nil
}

// Compose composition function
func (n NoOp) Compose(b PrimitiveOp) PrimitiveOp {
	switch b := b.(type) {
	case Insert:
		return b
	default:
		panic(ErrUnexpectedOp)
	}
}

// Transform transformation function
func (n NoOp) Transform(b PrimitiveOp) (aPrime, bPrime PrimitiveOp) {
	switch b := b.(type) {
	case Insert:
		return Retain{Count: b.OutputLength()}, b
	default:
		panic(ErrUnexpectedOp)
	}
}

func (n NoOp) String() string {
	return "NoOp"
}
