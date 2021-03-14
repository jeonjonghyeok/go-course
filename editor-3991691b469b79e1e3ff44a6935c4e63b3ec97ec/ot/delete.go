package ot

import (
	"fmt"
	"io"
	"io/ioutil"
)

// Delete delete operation
type Delete struct {
	Count int
}

// InputLength required input length
func (d Delete) InputLength() int {
	return d.Count
}

// OutputLength output string length
func (d Delete) OutputLength() int {
	return 0
}

// Slice slice the operation
func (d Delete) Slice(start, end int) PrimitiveOp {
	return Delete{Count: end - start}
}

// Apply executes the operation
func (d Delete) Apply(r io.Reader, w io.Writer) error {
	_, err := io.CopyN(ioutil.Discard, r, int64(d.Count))
	return err
}

// Join joins the operation with another one if possible
func (d Delete) Join(next PrimitiveOp) PrimitiveOp {
	if nextDelete, ok := next.(Delete); ok {
		return Delete{Count: d.Count + nextDelete.Count}
	}
	return nil
}

// Swap returns true it should be swapped with the next operation
func (d Delete) Swap(next PrimitiveOp) bool {
	if _, ok := next.(Insert); ok {
		return true
	}
	return false
}

// Compose composition function
func (d Delete) Compose(b PrimitiveOp) PrimitiveOp {
	checkComposeLength(d, b)

	switch b.(type) {
	case NoOp:
		return d
	default:
		panic(ErrUnexpectedOp)
	}
}

// Transform transformation function
func (d Delete) Transform(b PrimitiveOp) (aPrime, bPrime PrimitiveOp) {
	checkTransformLength(d, b)

	switch b := b.(type) {
	case Delete:
		return NoOp{}, NoOp{}
	case Retain:
		return Delete{Count: b.Count}, NoOp{}
	default:
		panic(ErrUnexpectedOp)
	}
}

func (d Delete) String() string {
	return fmt.Sprintf("Delete(%d)", d.Count)
}
