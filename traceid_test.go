package logger

import (
	"context"
	"testing"
)

func TestTraceID(t *testing.T) {
	tr := MakeTraceID()

	initID := tr.GetCurrent()

	if initID == "" {
		t.Error("The init id should be not empty")
	}

	if tr.GetPrev() != "" {
		t.Error("Prev trace ID should be empty on init")
	}

	tr = tr.NewStep()

	if initID != tr.GetPrev() {
		t.Errorf("The next step prevID shuld be equal the prev ID. Got: %v, Want:%v", tr.GetPrev(), initID)
	}
	if tr.GetCurrent() == "" {
		t.Error("The current id on the next step should be not empty")
	}
	if initID == tr.GetCurrent() {
		t.Errorf("The next step id should be different like the init id. Got: %v, Want:%v", tr.GetCurrent(), initID)
	}

	ctx := context.Background()

	ctx = tr.SetToContext(ctx)

	tr2 := GetTraceIDFromContext(ctx)
	if tr.GetCurrent() != tr2.GetCurrent() {
		t.Errorf("trace id is should be equal with in the context stored traceID. Got: %v, Want:%v", tr2.GetCurrent(), tr.GetCurrent())
	}
	if tr.GetPrev() != tr2.GetPrev() {
		t.Errorf("trace prevID is should be equal with in the context stored prevID Got: %v, Want:%v", tr2.GetPrev(), tr.GetPrev())
	}
}
