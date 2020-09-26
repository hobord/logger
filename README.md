# Logger
Log system for microservices.

The basic concept is create json formated logs what You can consume with logstash...

- You can track Your logs flow with CorelationID, and TraceIDs
- You can add fields
- You can set up global level fileds
- You can add tags
- You can set up global level tags
- You can add error stack

It has one concrete implementation with github.com/sirupsen/logrus but it is free to implement with other log systems.

## CorelationID
It is a unique uuid for the entire flow across on the microservices network.

You can generate and add into the context.

If You make log ```.WithContext(ctx).Info("message")```, the correlation Id will automatically extracted and added into the log.

## TraceID
It is like the correlationID but in the flow it is changing and using linked list to shows the flow steps.

```
type TraceID struct {
	currentID string
	prevID    string
}
```

With ```traceID.NewStep()```  You can indicate the next step in Your process flow...

You can store TraceID also in the context ```traceID.SetToContext(ctx)```, and when You make make log with context ```.WithContext(ctx).Info("message")``` the it will automatically extracted and added into the log. 

## Example use:
```

Log = lgrs.MakeLoggerWithLogrus(nil)
ctx = context.Background()
corelationID := logger.MakeCorelationID()
ctx = corelationID.SetToContext(ctx)
traceID := logger.MakeTraceID()
ctx = traceID.SetToContext(ctx)

Log.SetDefaultFields(logger.Fields{
  "Version": Version,
  "GitHash": GitHash,
  "AppPath": filepath.Dir(os.Args[0]),
  "AppBin":  filepath.Base(os.Args[0]),
})

Log.SetDefaultTags([]string{"apiserver", "orders"})

Log.WithContext(ctx).Info("App start")

...

traceID = traceID.NewStep()
ctx = traceID.SetToContext(ctx)

...
err := ...
err = errors.wrap(err,"MyCustom error msg")
Log.WithContext(ctx).WithError(err).Error("Error message")
...

```