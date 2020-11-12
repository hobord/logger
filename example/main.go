package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hobord/logger"
	"github.com/hobord/logger/lgrs"
	"github.com/pkg/errors"
)

var Log logger.Logger

func main() {
	Log = lgrs.MakeLoggerWithLogrus(nil)
	ctx := context.Background()
	corelationID := logger.MakeCorelationID()
	ctx = corelationID.SetToContext(ctx)
	traceID := logger.MakeTraceID()
	ctx = traceID.SetToContext(ctx)

	Log.SetDefaultFields(logger.Fields{
		"Version": "v1.0.0",
		"GitHash": "fa2309",
		"AppPath": filepath.Dir(os.Args[0]),
		"AppBin":  filepath.Base(os.Args[0]),
	})

	Log.SetDefaultTags([]string{"apiserver", "orders"})

	Log.WithContext(ctx).Info("App start")

	traceID = traceID.NewStep()
	ctx = traceID.SetToContext(ctx)

	Log.WithContext(ctx).
		WithTags([]string{"Tag1", "Tag2"}).
		WithFields(logger.Fields{
			"id":     15,
			"action": "order",
			"user": struct {
				ID   int
				Name string
			}{ID: 15, Name: "Bob"},
			"items": []string{"item1", "item2"},
		}).Info("New request")

	err := fnc()
	Log.WithContext(ctx).WithError(err).Error("Log message")
}

func fnc() error {
	err := fmt.Errorf("My custom error message into error field")
	err = errors.Wrap(err, "Wrap error")
	return err
}
