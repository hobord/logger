package logger

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

var (
	// ContextKeyTraceID is the key in context for TraceID
	ContextKeyTraceID = contextKey("TraceID")
	// LogFieldKeyTraceIDCurrent is the key in log for current trace id
	LogFieldKeyTraceIDCurrent = "TraceIDCurrent"
	// LogFieldKeyTraceIDPrev is key in the log for the pre trace id
	LogFieldKeyTraceIDPrev = "TraceIDPrev"
)

// TraceID is an inmutable ID for track processflow.
type TraceID struct {
	currentID string
	prevID    string
}

// TracerID is help follow the steps on the processflow
type TracerID interface {
	// Get current id of the process
	GetCurrent() string
	// Get prev process ID
	GetPrev() string
	// Create a new process step
	NewStep(current TraceID) *TraceID
}

// MakeTraceID is generate a new TraceID
func MakeTraceID() *TraceID {
	return &TraceID{
		currentID: uuid.New().String(),
		prevID:    "",
	}
}

// NewStep is create a new process steps
func (t *TraceID) NewStep() *TraceID {
	return &TraceID{
		currentID: uuid.New().String(),
		prevID:    t.GetCurrent(),
	}
}

// MarshalJSON marshal struct
func (t *TraceID) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		CurrentID string `json:"current_id"`
		PrevID    string `json:"prev_id"`
	}{
		CurrentID: t.GetCurrent(),
		PrevID:    t.GetPrev(),
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

// UnmarshalJSON is unmarshal TraceID json
func (t *TraceID) UnmarshalJSON(b []byte) error {
	type T struct {
		CurrentID string `json:"current_id"`
		PrevID    string `json:"prev_id"`
	}
	var tmp T
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	t.currentID = tmp.CurrentID
	t.prevID = tmp.PrevID

	return nil
}

// GetCurrent return with the  current id of the process
func (t *TraceID) GetCurrent() string {
	return t.currentID
}

// GetPrev is return with the prev process ID
func (t *TraceID) GetPrev() string {
	return t.prevID
}

// GetTraceIDFromContext is return a TraceID from ctx
func GetTraceIDFromContext(ctx context.Context) *TraceID {
	if c, ok := ctx.Value(ContextKeyTraceID).(TraceID); ok {
		return &c
	}
	return nil
}

// SetToContext set this TraceID into the context
func (t *TraceID) SetToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ContextKeyTraceID, *t)
}
