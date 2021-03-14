package ot

import (
	"fmt"
	"io"
)

// Retain retain operation
type Retain struct {
	Count int
}

// InputLength required input length
func (r Retain) InputLength() int {
	return r.Count
}

// OutputLength output string length
func (r Retain) OutputLength() int {
	return r.Count
}

// Slice slice the operation
func (r Retain) Slice(start, end int) PrimitiveOp {
	return Retain{Count: end - start}
}

// Apply executes the operation
func (r Retain) Apply(reader io.Reader, w io.Writer) error {
	_, err := io.CopyN(w, reader, int64(r.Count))
	return err
}

// Join joins the operation with another one if possible
func (r Retain) Join(next PrimitiveOp) PrimitiveOp {
	if nextRetain, ok := next.(Retain); ok {
		return Retain{Count: r.Count + nextRetain.Count}
	}
	return nil
}

// Compose composition function
func (r Retain) Compose(b PrimitiveOp) PrimitiveOp {
	checkComposeLength(r, b)

	switch b := b.(type) {
	case Retain:
		return r
	case Delete:
		return b
	default:
		panic(ErrUnexpectedOp)
	}
}

// Transform transformation function
func (r Retain) Transform(b PrimitiveOp) (aPrime, bPrime PrimitiveOp) {
	checkTransformLength(r, b)

	switch b := b.(type) {
	case Retain:
		return b, b
	case Delete:
		return NoOp{}, b
	default:
		panic(ErrUnexpectedOp)
	}
}

func (r Retain) String() string {
	return fmt.Sprintf("Retain(%d)", r.Count)
}
