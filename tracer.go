package tracer

import (
	"fmt"
	"io"
)

// Tracer is an interface which describes
// an object capable of tracing events
// throughout code.
type Tracer interface {
	Trace(...interface{})
}

type tracer1 struct {
	out io.Writer
}

type nilTracer struct{}

func New(w io.Writer) Tracer {
	return &tracer1{
		out: w,
	}
}

func (t *tracer1) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

func (t *nilTracer) Trace(a ...interface{}) {}

func Off() Tracer {
	return &nilTracer{}
}
