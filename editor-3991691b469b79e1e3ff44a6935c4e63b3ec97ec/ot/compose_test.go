package ot_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/learningspoons-go/editor/ot"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func testCompose(t *testing.T, input string, ops ...ot.CompositeOp) {
	composed := ot.Compose(ops...)

	afterOp := input
	var err error
	for _, op := range ops {
		afterOp, err = ot.ApplyString(op, afterOp)
		if err != nil {
			t.Error(err)
			return
		}

	}

	afterComposed, err := ot.ApplyString(composed, input)
	if err != nil {
		t.Error(err)
		return
	}

	if afterOp != afterComposed {
		t.Errorf("afterOp(%s) != afterComposed(%s), ops: %v",
			afterOp, afterComposed, ops)
	}
}

func TestCompose(t *testing.T) {
	for i := 0; i < 10000; i++ {
		t.Run(fmt.Sprintf("rand-%d", i), func(t *testing.T) {
			t.Parallel()
			inputLength := rand.Intn(20) + 5
			inputStr := randString(inputLength)

			opCount := rand.Intn(8) + 2
			ops := make([]ot.CompositeOp, opCount)
			outLength := inputLength + rand.Intn(10) - 5
			for i := 0; i < opCount; i++ {
				ops[i] = randomCompositeOp(inputLength, outLength)
				inputLength = outLength
				outLength = outLength + rand.Intn(5)
			}

			testCompose(t, inputStr, ops...)
		})
	}
}
