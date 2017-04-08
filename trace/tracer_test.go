package trace

import (
	"bytes"
	"testing"

	"github.com/nitinstp23/chat_app/trace"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := trace.New(&buf)

	if tracer == nil {
		t.Error("Return from New should not be nil")
	} else {
		tracer.Trace("Hello trace package.")
		if buf.String() != "Hello trace package.\n" {
			t.Errorf("Trace should not write '%s'.", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	var silentTracer trace.Tracer = trace.Off()
	silentTracer.Trace("something")
}
