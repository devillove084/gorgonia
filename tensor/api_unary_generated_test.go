package tensor

import (
	"math/rand"
	"testing"
	"testing/quick"
	"time"
)

/*
GENERATED FILE. DO NOT EDIT
*/

func TestNeg(t *testing.T) {
	var r *rand.Rand
	invFn := func(q *Dense) bool {
		a := q.Clone().(*Dense)
		correct := a.Clone().(*Dense)
		we, willFailEq := willerr(a, numberTypes, nil)
		ret, err := Neg(a)
		if err, retEarly := qcErrCheck(t, "Neg", a, nil, we, err); retEarly {
			if err != nil {
				return false
			}
			return true
		}
		Neg(ret, UseUnsafe())
		if !qcEqCheck(t, a.Dtype(), willFailEq, correct.Data(), ret.Data()) {
			return false
		}
		return true
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := quick.Check(invFn, &quick.Config{Rand: r}); err != nil {
		t.Errorf("Inv tests for Neg failed: %v", err)
	}
}
func TestSquare(t *testing.T) {
	var r *rand.Rand
	invFn := func(q *Dense) bool {
		a := q.Clone().(*Dense)
		correct := a.Clone().(*Dense)
		we, willFailEq := willerr(a, numberTypes, nil)
		ret, err := Square(a)
		if err, retEarly := qcErrCheck(t, "Square", a, nil, we, err); retEarly {
			if err != nil {
				return false
			}
			return true
		}
		if err := typeclassCheck(a.Dtype(), floatcmplxTypes); err != nil {
			return true // uninvertible due to type class implementation issues
		}
		Sqrt(ret, UseUnsafe())
		if !qcEqCheck(t, a.Dtype(), willFailEq, correct.Data(), ret.Data()) {
			return false
		}
		return true
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := quick.Check(invFn, &quick.Config{Rand: r}); err != nil {
		t.Errorf("Inv tests for Square failed: %v", err)
	}
}
func TestCube(t *testing.T) {
	var r *rand.Rand
	invFn := func(q *Dense) bool {
		a := q.Clone().(*Dense)
		correct := a.Clone().(*Dense)
		we, willFailEq := willerr(a, numberTypes, nil)
		ret, err := Cube(a)
		if err, retEarly := qcErrCheck(t, "Cube", a, nil, we, err); retEarly {
			if err != nil {
				return false
			}
			return true
		}
		if err := typeclassCheck(a.Dtype(), floatTypes); err != nil {
			return true // uninvertible due to type class implementation issues
		}
		Cbrt(ret, UseUnsafe())
		if !qcEqCheck(t, a.Dtype(), willFailEq, correct.Data(), ret.Data()) {
			return false
		}
		return true
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := quick.Check(invFn, &quick.Config{Rand: r}); err != nil {
		t.Errorf("Inv tests for Cube failed: %v", err)
	}
}
func TestExp(t *testing.T) {
	var r *rand.Rand
	invFn := func(q *Dense) bool {
		a := q.Clone().(*Dense)
		correct := a.Clone().(*Dense)
		we, willFailEq := willerr(a, floatcmplxTypes, nil)
		ret, err := Exp(a)
		if err, retEarly := qcErrCheck(t, "Exp", a, nil, we, err); retEarly {
			if err != nil {
				return false
			}
			return true
		}
		Log(ret, UseUnsafe())
		if !qcEqCheck(t, a.Dtype(), willFailEq, correct.Data(), ret.Data()) {
			return false
		}
		return true
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := quick.Check(invFn, &quick.Config{Rand: r}); err != nil {
		t.Errorf("Inv tests for Exp failed: %v", err)
	}
}
func TestLog(t *testing.T) {
	var r *rand.Rand
	invFn := func(q *Dense) bool {
		a := q.Clone().(*Dense)
		correct := a.Clone().(*Dense)
		we, willFailEq := willerr(a, floatcmplxTypes, nil)
		ret, err := Log(a)
		if err, retEarly := qcErrCheck(t, "Log", a, nil, we, err); retEarly {
			if err != nil {
				return false
			}
			return true
		}
		Exp(ret, UseUnsafe())
		if !qcEqCheck(t, a.Dtype(), willFailEq, correct.Data(), ret.Data()) {
			return false
		}
		return true
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := quick.Check(invFn, &quick.Config{Rand: r}); err != nil {
		t.Errorf("Inv tests for Log failed: %v", err)
	}
}
func TestSqrt(t *testing.T) {
	var r *rand.Rand
	invFn := func(q *Dense) bool {
		a := q.Clone().(*Dense)
		correct := a.Clone().(*Dense)
		we, willFailEq := willerr(a, floatcmplxTypes, nil)
		ret, err := Sqrt(a)
		if err, retEarly := qcErrCheck(t, "Sqrt", a, nil, we, err); retEarly {
			if err != nil {
				return false
			}
			return true
		}
		Square(ret, UseUnsafe())
		if !qcEqCheck(t, a.Dtype(), willFailEq, correct.Data(), ret.Data()) {
			return false
		}
		return true
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := quick.Check(invFn, &quick.Config{Rand: r}); err != nil {
		t.Errorf("Inv tests for Sqrt failed: %v", err)
	}
}
func TestCbrt(t *testing.T) {
	var r *rand.Rand
	invFn := func(q *Dense) bool {
		a := q.Clone().(*Dense)
		correct := a.Clone().(*Dense)
		we, willFailEq := willerr(a, floatTypes, nil)
		ret, err := Cbrt(a)
		if err, retEarly := qcErrCheck(t, "Cbrt", a, nil, we, err); retEarly {
			if err != nil {
				return false
			}
			return true
		}
		Cube(ret, UseUnsafe())
		if !qcEqCheck(t, a.Dtype(), willFailEq, correct.Data(), ret.Data()) {
			return false
		}
		return true
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := quick.Check(invFn, &quick.Config{Rand: r}); err != nil {
		t.Errorf("Inv tests for Cbrt failed: %v", err)
	}
}
func TestNeg_unsafe(t *testing.T) {
	var r *rand.Rand
	invFn := func(q *Dense) bool {
		a := q.Clone().(*Dense)
		correct := a.Clone().(*Dense)
		we, willFailEq := willerr(a, numberTypes, nil)
		ret, err := Neg(a, UseUnsafe())
		if err, retEarly := qcErrCheck(t, "Neg", a, nil, we, err); retEarly {
			if err != nil {
				return false
			}
			return true
		}
		Neg(ret, UseUnsafe())
		if !qcEqCheck(t, a.Dtype(), willFailEq, correct.Data(), ret.Data()) {
			return false
		}
		if ret != a {
			t.Errorf("Expected ret to be the same as a")
			return false
		}

		return true
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := quick.Check(invFn, &quick.Config{Rand: r}); err != nil {
		t.Errorf("Inv tests for Neg failed: %v", err)
	}
}
func TestSquare_unsafe(t *testing.T) {
	var r *rand.Rand
	invFn := func(q *Dense) bool {
		a := q.Clone().(*Dense)
		correct := a.Clone().(*Dense)
		we, willFailEq := willerr(a, numberTypes, nil)
		ret, err := Square(a, UseUnsafe())
		if err, retEarly := qcErrCheck(t, "Square", a, nil, we, err); retEarly {
			if err != nil {
				return false
			}
			return true
		}
		if err := typeclassCheck(a.Dtype(), floatcmplxTypes); err != nil {
			return true // uninvertible due to type class implementation issues
		}
		Sqrt(ret, UseUnsafe())
		if !qcEqCheck(t, a.Dtype(), willFailEq, correct.Data(), ret.Data()) {
			return false
		}
		if ret != a {
			t.Errorf("Expected ret to be the same as a")
			return false
		}

		return true
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := quick.Check(invFn, &quick.Config{Rand: r}); err != nil {
		t.Errorf("Inv tests for Square failed: %v", err)
	}
}
func TestCube_unsafe(t *testing.T) {
	var r *rand.Rand
	invFn := func(q *Dense) bool {
		a := q.Clone().(*Dense)
		correct := a.Clone().(*Dense)
		we, willFailEq := willerr(a, numberTypes, nil)
		ret, err := Cube(a, UseUnsafe())
		if err, retEarly := qcErrCheck(t, "Cube", a, nil, we, err); retEarly {
			if err != nil {
				return false
			}
			return true
		}
		if err := typeclassCheck(a.Dtype(), floatTypes); err != nil {
			return true // uninvertible due to type class implementation issues
		}
		Cbrt(ret, UseUnsafe())
		if !qcEqCheck(t, a.Dtype(), willFailEq, correct.Data(), ret.Data()) {
			return false
		}
		if ret != a {
			t.Errorf("Expected ret to be the same as a")
			return false
		}

		return true
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := quick.Check(invFn, &quick.Config{Rand: r}); err != nil {
		t.Errorf("Inv tests for Cube failed: %v", err)
	}
}
func TestExp_unsafe(t *testing.T) {
	var r *rand.Rand
	invFn := func(q *Dense) bool {
		a := q.Clone().(*Dense)
		correct := a.Clone().(*Dense)
		we, willFailEq := willerr(a, floatcmplxTypes, nil)
		ret, err := Exp(a, UseUnsafe())
		if err, retEarly := qcErrCheck(t, "Exp", a, nil, we, err); retEarly {
			if err != nil {
				return false
			}
			return true
		}
		Log(ret, UseUnsafe())
		if !qcEqCheck(t, a.Dtype(), willFailEq, correct.Data(), ret.Data()) {
			return false
		}
		if ret != a {
			t.Errorf("Expected ret to be the same as a")
			return false
		}

		return true
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := quick.Check(invFn, &quick.Config{Rand: r}); err != nil {
		t.Errorf("Inv tests for Exp failed: %v", err)
	}
}
func TestLog_unsafe(t *testing.T) {
	var r *rand.Rand
	invFn := func(q *Dense) bool {
		a := q.Clone().(*Dense)
		correct := a.Clone().(*Dense)
		we, willFailEq := willerr(a, floatcmplxTypes, nil)
		ret, err := Log(a, UseUnsafe())
		if err, retEarly := qcErrCheck(t, "Log", a, nil, we, err); retEarly {
			if err != nil {
				return false
			}
			return true
		}
		Exp(ret, UseUnsafe())
		if !qcEqCheck(t, a.Dtype(), willFailEq, correct.Data(), ret.Data()) {
			return false
		}
		if ret != a {
			t.Errorf("Expected ret to be the same as a")
			return false
		}

		return true
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := quick.Check(invFn, &quick.Config{Rand: r}); err != nil {
		t.Errorf("Inv tests for Log failed: %v", err)
	}
}
func TestSqrt_unsafe(t *testing.T) {
	var r *rand.Rand
	invFn := func(q *Dense) bool {
		a := q.Clone().(*Dense)
		correct := a.Clone().(*Dense)
		we, willFailEq := willerr(a, floatcmplxTypes, nil)
		ret, err := Sqrt(a, UseUnsafe())
		if err, retEarly := qcErrCheck(t, "Sqrt", a, nil, we, err); retEarly {
			if err != nil {
				return false
			}
			return true
		}
		Square(ret, UseUnsafe())
		if !qcEqCheck(t, a.Dtype(), willFailEq, correct.Data(), ret.Data()) {
			return false
		}
		if ret != a {
			t.Errorf("Expected ret to be the same as a")
			return false
		}

		return true
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := quick.Check(invFn, &quick.Config{Rand: r}); err != nil {
		t.Errorf("Inv tests for Sqrt failed: %v", err)
	}
}
func TestCbrt_unsafe(t *testing.T) {
	var r *rand.Rand
	invFn := func(q *Dense) bool {
		a := q.Clone().(*Dense)
		correct := a.Clone().(*Dense)
		we, willFailEq := willerr(a, floatTypes, nil)
		ret, err := Cbrt(a, UseUnsafe())
		if err, retEarly := qcErrCheck(t, "Cbrt", a, nil, we, err); retEarly {
			if err != nil {
				return false
			}
			return true
		}
		Cube(ret, UseUnsafe())
		if !qcEqCheck(t, a.Dtype(), willFailEq, correct.Data(), ret.Data()) {
			return false
		}
		if ret != a {
			t.Errorf("Expected ret to be the same as a")
			return false
		}

		return true
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := quick.Check(invFn, &quick.Config{Rand: r}); err != nil {
		t.Errorf("Inv tests for Cbrt failed: %v", err)
	}
}
func TestNeg_reuse(t *testing.T) {
	var r *rand.Rand
	invFn := func(q *Dense) bool {
		a := q.Clone().(*Dense)
		reuse := New(Of(a.t), WithShape(a.Shape().Clone()...))
		correct := a.Clone().(*Dense)
		we, willFailEq := willerr(a, numberTypes, nil)
		ret, err := Neg(a, WithReuse(reuse))
		if err, retEarly := qcErrCheck(t, "Neg", a, nil, we, err); retEarly {
			if err != nil {
				return false
			}
			return true
		}
		Neg(ret, UseUnsafe())
		if !qcEqCheck(t, a.Dtype(), willFailEq, correct.Data(), ret.Data()) {
			return false
		}
		if reuse != ret {
			t.Errorf("Expected reuse to be the same as retVal")
			return false
		}

		return true
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := quick.Check(invFn, &quick.Config{Rand: r}); err != nil {
		t.Errorf("Inv tests for Neg failed: %v", err)
	}
}
func TestSquare_reuse(t *testing.T) {
	var r *rand.Rand
	invFn := func(q *Dense) bool {
		a := q.Clone().(*Dense)
		reuse := New(Of(a.t), WithShape(a.Shape().Clone()...))
		correct := a.Clone().(*Dense)
		we, willFailEq := willerr(a, numberTypes, nil)
		ret, err := Square(a, WithReuse(reuse))
		if err, retEarly := qcErrCheck(t, "Square", a, nil, we, err); retEarly {
			if err != nil {
				return false
			}
			return true
		}
		if err := typeclassCheck(a.Dtype(), floatcmplxTypes); err != nil {
			return true // uninvertible due to type class implementation issues
		}
		Sqrt(ret, UseUnsafe())
		if !qcEqCheck(t, a.Dtype(), willFailEq, correct.Data(), ret.Data()) {
			return false
		}
		if reuse != ret {
			t.Errorf("Expected reuse to be the same as retVal")
			return false
		}

		return true
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := quick.Check(invFn, &quick.Config{Rand: r}); err != nil {
		t.Errorf("Inv tests for Square failed: %v", err)
	}
}
func TestCube_reuse(t *testing.T) {
	var r *rand.Rand
	invFn := func(q *Dense) bool {
		a := q.Clone().(*Dense)
		reuse := New(Of(a.t), WithShape(a.Shape().Clone()...))
		correct := a.Clone().(*Dense)
		we, willFailEq := willerr(a, numberTypes, nil)
		ret, err := Cube(a, WithReuse(reuse))
		if err, retEarly := qcErrCheck(t, "Cube", a, nil, we, err); retEarly {
			if err != nil {
				return false
			}
			return true
		}
		if err := typeclassCheck(a.Dtype(), floatTypes); err != nil {
			return true // uninvertible due to type class implementation issues
		}
		Cbrt(ret, UseUnsafe())
		if !qcEqCheck(t, a.Dtype(), willFailEq, correct.Data(), ret.Data()) {
			return false
		}
		if reuse != ret {
			t.Errorf("Expected reuse to be the same as retVal")
			return false
		}

		return true
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := quick.Check(invFn, &quick.Config{Rand: r}); err != nil {
		t.Errorf("Inv tests for Cube failed: %v", err)
	}
}
func TestExp_reuse(t *testing.T) {
	var r *rand.Rand
	invFn := func(q *Dense) bool {
		a := q.Clone().(*Dense)
		reuse := New(Of(a.t), WithShape(a.Shape().Clone()...))
		correct := a.Clone().(*Dense)
		we, willFailEq := willerr(a, floatcmplxTypes, nil)
		ret, err := Exp(a, WithReuse(reuse))
		if err, retEarly := qcErrCheck(t, "Exp", a, nil, we, err); retEarly {
			if err != nil {
				return false
			}
			return true
		}
		Log(ret, UseUnsafe())
		if !qcEqCheck(t, a.Dtype(), willFailEq, correct.Data(), ret.Data()) {
			return false
		}
		if reuse != ret {
			t.Errorf("Expected reuse to be the same as retVal")
			return false
		}

		return true
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := quick.Check(invFn, &quick.Config{Rand: r}); err != nil {
		t.Errorf("Inv tests for Exp failed: %v", err)
	}
}
func TestLog_reuse(t *testing.T) {
	var r *rand.Rand
	invFn := func(q *Dense) bool {
		a := q.Clone().(*Dense)
		reuse := New(Of(a.t), WithShape(a.Shape().Clone()...))
		correct := a.Clone().(*Dense)
		we, willFailEq := willerr(a, floatcmplxTypes, nil)
		ret, err := Log(a, WithReuse(reuse))
		if err, retEarly := qcErrCheck(t, "Log", a, nil, we, err); retEarly {
			if err != nil {
				return false
			}
			return true
		}
		Exp(ret, UseUnsafe())
		if !qcEqCheck(t, a.Dtype(), willFailEq, correct.Data(), ret.Data()) {
			return false
		}
		if reuse != ret {
			t.Errorf("Expected reuse to be the same as retVal")
			return false
		}

		return true
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := quick.Check(invFn, &quick.Config{Rand: r}); err != nil {
		t.Errorf("Inv tests for Log failed: %v", err)
	}
}
func TestSqrt_reuse(t *testing.T) {
	var r *rand.Rand
	invFn := func(q *Dense) bool {
		a := q.Clone().(*Dense)
		reuse := New(Of(a.t), WithShape(a.Shape().Clone()...))
		correct := a.Clone().(*Dense)
		we, willFailEq := willerr(a, floatcmplxTypes, nil)
		ret, err := Sqrt(a, WithReuse(reuse))
		if err, retEarly := qcErrCheck(t, "Sqrt", a, nil, we, err); retEarly {
			if err != nil {
				return false
			}
			return true
		}
		Square(ret, UseUnsafe())
		if !qcEqCheck(t, a.Dtype(), willFailEq, correct.Data(), ret.Data()) {
			return false
		}
		if reuse != ret {
			t.Errorf("Expected reuse to be the same as retVal")
			return false
		}

		return true
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := quick.Check(invFn, &quick.Config{Rand: r}); err != nil {
		t.Errorf("Inv tests for Sqrt failed: %v", err)
	}
}
func TestCbrt_reuse(t *testing.T) {
	var r *rand.Rand
	invFn := func(q *Dense) bool {
		a := q.Clone().(*Dense)
		reuse := New(Of(a.t), WithShape(a.Shape().Clone()...))
		correct := a.Clone().(*Dense)
		we, willFailEq := willerr(a, floatTypes, nil)
		ret, err := Cbrt(a, WithReuse(reuse))
		if err, retEarly := qcErrCheck(t, "Cbrt", a, nil, we, err); retEarly {
			if err != nil {
				return false
			}
			return true
		}
		Cube(ret, UseUnsafe())
		if !qcEqCheck(t, a.Dtype(), willFailEq, correct.Data(), ret.Data()) {
			return false
		}
		if reuse != ret {
			t.Errorf("Expected reuse to be the same as retVal")
			return false
		}

		return true
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := quick.Check(invFn, &quick.Config{Rand: r}); err != nil {
		t.Errorf("Inv tests for Cbrt failed: %v", err)
	}
}
