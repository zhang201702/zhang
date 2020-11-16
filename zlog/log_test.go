package zlog

import (
	"errors"
	"testing"
)

type testS struct {
	A  string
	B  string
	As []*testS
	Bs map[int]*testS
}

func TestLog(t *testing.T) {
	Log("test", "OK")
	t.Fail()
}

func TestLogError(t *testing.T) {
	err := errors.New("tst error")
	d := testS{
		"A", "B", make([]*testS, 0), make(map[int]*testS),
	}
	a1 := &testS{A: "A1", B: "B1"}
	d.As = append(d.As, a1)
	d.As = append(d.As, &testS{A: "A2", B: "B2"})
	d.As = append(d.As, &testS{A: "A3", B: "B3"})
	d.Bs[1] = a1
	d.Bs[2] = &testS{A: "A22", B: "B22"}
	d.Bs[3] = &testS{A: "A33", B: "B33"}

	LogError(err, "test", d, &d)
}
