package ot

import (
	"io"
)

// CompositeOp composite operation
type CompositeOp []PrimitiveOp

// NewCompositeOp new normalized composite operation
func NewCompositeOp(ops ...PrimitiveOp) CompositeOp {
	return normalize(ops)
}

// InputLength input length
func (c CompositeOp) InputLength() (length int) {
	for _, p := range c {
		length += p.InputLength()
	}
	return
}

// OutputLength output length
func (c CompositeOp) OutputLength() (length int) {
	for _, p := range c {
		length += p.OutputLength()
	}
	return
}

// Apply executes the operation
func (c CompositeOp) Apply(r io.Reader, w io.Writer) error {
	for _, p := range c {
		if err := p.Apply(r, w); err != nil {
			return err
		}
	}
	return nil
}

// Transform transformation function
func (c CompositeOp) Transform(b CompositeOp) (aPrime, bPrime CompositeOp) {
	slicedA, slicedB := slice(c, b,
		inputLengthFunc, inputLengthFunc)

	for i := range slicedA {
		aOp, bOp := slicedA[i], slicedB[i]
		aOpPrime, bOpPrime := aOp.Transform(bOp)
		aPrime = append(aPrime, aOpPrime)
		bPrime = append(bPrime, bOpPrime)
	}
	return NewCompositeOp(aPrime...), NewCompositeOp(bPrime...)
}

// Compose composition function
func (c CompositeOp) Compose(b CompositeOp) CompositeOp {
	slicedA, slicedB := slice(c, b,
		outputLengthFunc, inputLengthFunc)

	var res []PrimitiveOp
	for i := range slicedA {
		aOp, bOp := slicedA[i], slicedB[i]
		c := aOp.Compose(bOp)
		res = append(res, c)
	}
	return NewCompositeOp(res...)
}

// Compose compose multiple operations
func Compose(ops ...CompositeOp) CompositeOp {
	if l := len(ops); l == 0 {
		return CompositeOp{}
	} else if l < 2 {
		return ops[0]
	}

	op := ops[0]
	for i := 1; i < len(ops); i++ {
		op = op.Compose(ops[i])
	}
	return op
}
