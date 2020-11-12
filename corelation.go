package logger

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type contextKey string

func (c contextKey) String() string {
	return "github.com/hobord context_key " + string(c)
}

var (
	// ContextKeyCorelationID is the key in context for CorelationID
	ContextKeyCorelationID = contextKey("CorelationID")
	// LogFieldKeyCorelationID is the key in log for CorelationID
	LogFieldKeyCorelationID = "CorelationID"
)

// CorelationID is an inmutable ID for track processflow.
type CorelationID struct {
	corealtionID string
}

func (c *CorelationID) String() string {
	return c.corealtionID
}

// MakeCorelationID is generate a new CorelationID
func MakeCorelationID() *CorelationID {
	return &CorelationID{
		corealtionID: uuid.New().String(),
	}
}

// MarshalJSON marshal struct
func (c *CorelationID) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		ID string `json:"id"`
	}{
		ID: c.corealtionID,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

// UnmarshalJSON is unmarshal CorelationID json
func (c *CorelationID) UnmarshalJSON(b []byte) error {
	type T struct {
		ID string `json:"id"`
	}
	var tmp T
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	c.corealtionID = tmp.ID
	return nil
}

// GetCorelationIDValFromContext is return corelationID string from context
// if not extists then ""
func GetCorelationIDValFromContext(ctx context.Context) string {
	if c, ok := ctx.Value(ContextKeyCorelationID).(CorelationID); ok {
		return c.String()
	}
	return ""
}

// ContextCorelationID is return a CorelationID from the context
// if the context has no corelation then generate and and into it.
func ContextCorelationID(ctx context.Context) (*CorelationID, context.Context) {
	corelationID := GetCorelationIDFromContext(ctx)
	if corelationID == nil {
		corelationID = MakeCorelationID()
		ctx = corelationID.SetToContext(ctx)
	}
	return corelationID, ctx
}

// GetCorelationIDFromContext is get CorelationID from context
func GetCorelationIDFromContext(ctx context.Context) *CorelationID {
	if c, ok := ctx.Value(ContextKeyCorelationID).(CorelationID); ok {
		return &c
	}
	return nil
}

// SetToContext set this CorealtionID into the context
func (c *CorelationID) SetToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ContextKeyCorelationID, *c)
}
