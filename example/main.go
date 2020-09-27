package main

import (
	"context"
	"os"
	"path/filepath"

	"github.com/hobord/logger"
	"github.com/hobord/logger/lgrs"
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
		"Version": "1",
		"GitHash": "fa2309",
		"AppPath": filepath.Dir(os.Args[0]),
		"AppBin":  filepath.Base(os.Args[0]),
	})

	Log.SetDefaultTags([]string{"apiserver", "orders"})

	Log.WithContext(ctx).Info("App start")

	traceID = traceID.NewStep()
	ctx = traceID.SetToContext(ctx)

	Log.WithContext(ctx).WithFields(logger.Fields{
		"id":     15,
		"action": "order",
		"user": struct {
			ID   int
			Name string
		}{ID: 15, Name: "Bob"},
		"items": []string{"item1", "item2"},
	}).Info("New request")
}
