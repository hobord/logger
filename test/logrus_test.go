package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hobord/logger"
	"github.com/hobord/logger/lgrs"
	"github.com/pkg/errors"
)

func TestLogrus(t *testing.T) {
	ctx := context.Background()
	corelationID, ctx := logger.ContextCorelationID(ctx)

	var log logger.Logger
	log = lgrs.MakeLoggerWithLogrus(nil)

	ctxTes := log.WithContext(ctx)
	ctxTes = ctxTes.WithField("test", "ok")
	ctxTes.Info("Hello")

	corelationID = logger.MakeCorelationID()
	log.WithCorelationID(*corelationID).Info()

	err := fmt.Errorf("My custom error with var:%v", 42)
	err = errors.WithStack(err)
	err = errors.Wrap(err, "level 1")
	err = errors.Wrap(err, "level 2")

	log.WithError(err).Error()

	log.SetDefaultTags([]string{"tag1", "tag1", "tag2", "tag3"})

	log.Info("hello")

	log.WithTags([]string{"custom1", "custom1", "custom2", "tag2"}).WithContext(ctx).Info("hello")
	log.SetDefaultContext(ctx)
	log.WithTags([]string{"asd"}).Info("hello")

	// t.Error()
}
