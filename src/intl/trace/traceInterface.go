package trace

import "time"

type ITracer interface {
	TraceRequest(
		isParent bool,
		method string,
		path string,
		query string,
		statusCode int,
		bodySize int,
		ip string,
		userAgent string,
		startTimestamp time.Time,
		eventTimestamp time.Time,
		fields ...Field,
	)
	TraceDependency(
		spanId string,
		dependencyType string,
		serviceName string,
		commandName string,
		success bool,
		startTimestamp time.Time,
		eventTimestamp time.Time,
		fields ...Field,
	)
}

type Field struct {
	Key   string
	Value string
}

func NewField (key string, value string) Field {
	return Field{
		Key: key,
		Value: value,
	}
}
