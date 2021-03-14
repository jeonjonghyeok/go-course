package ot

import (
	"fmt"
	"io"
)

// Insert insert operation
type Insert struct {
	Text string
}

// InputLength required input length
func (i Insert) InputLength() int {
	return 0
}

// OutputLength output string length
func (i Insert) OutputLength() int {
	return len(i.Text)
}

// Slice slice the operation
func (i Insert) Slice(start, end int) PrimitiveOp {
	return Insert{Text: i.Text[start:end]}
}

// Apply executes the operation
func (i Insert) Apply(r io.Reader, w io.Writer) error {
	_, err := io.WriteString(w, i.Text)
	return err
}

// Join joins the operation with another one if possible
func (i Insert) Join(next PrimitiveOp) PrimitiveOp {
	if nextInsert, ok := next.(Insert); ok {
		return Insert{Text: i.Text + nextInsert.Text}
	}
	return nil
}

// Compose composition function
func (i Insert) Compose(b PrimitiveOp) PrimitiveOp {
	checkComposeLength(i, b)

	switch b.(type) {
	case Delete:
		return NoOp{}
	case Retain:
		return i
	default:
		panic(ErrUnexpectedOp)
	}
}

// Transform transformation function
func (i Insert) Transform(b PrimitiveOp) (aPrime, bPrime PrimitiveOp) {
	checkTransformLength(i, b)

	switch b.(type) {
	case NoOp:
		return i, Retain{Count: i.OutputLength()}
	default:
		panic(ErrUnexpectedOp)
	}
}

func (i Insert) String() string {
	return fmt.Sprintf("Insert(%s)", i.Text)
}
