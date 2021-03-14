package ot

import (
	"bytes"
	"errors"
	"io"
	"strings"
)

// Op operation that can be applied
type Op interface {
	InputLength() int
	OutputLength() int
	Apply(io.Reader, io.Writer) error
}

// PrimitiveOp primitive operation used in CompositeOp
type PrimitiveOp interface {
	Op
	Slice(start, end int) PrimitiveOp
	Compose(b PrimitiveOp) PrimitiveOp
	Transform(b PrimitiveOp) (aPrime, bPrime PrimitiveOp)
}

// ErrLengthMismatch length mismatch error
var ErrLengthMismatch = errors.New("length mismatch")

// ErrUnexpectedOp unexpected operation error
var ErrUnexpectedOp = errors.New("unexpected operation")

func checkComposeLength(a, b Op) {
	if a.OutputLength() != b.InputLength() {
		panic(ErrLengthMismatch)
	}
}

func checkTransformLength(a, b Op) {
	if a.InputLength() != b.InputLength() {
		panic(ErrLengthMismatch)
	}
}

// ApplyString applies the operation on a string
func ApplyString(op Op, text string) (string, error) {
	reader := bytes.NewReader([]byte(text))
	out := new(strings.Builder)
	if err := op.Apply(reader, out); err != nil {
		return "", err
	}
	return out.String(), nil
}
