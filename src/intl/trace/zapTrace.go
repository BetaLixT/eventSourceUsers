package trace

import (
	"time"

	"go.uber.org/zap"
)

type ZapTracer struct {
	lgr *zap.Logger
}

var _ ITracer = (*ZapTracer)(nil)

func NewZapTracer(lgr *zap.Logger) *ZapTracer {
	return &ZapTracer{
		lgr: lgr,
	}
}

func (tracer *ZapTracer) TraceRequest(
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

) {

	latency := eventTimestamp.Sub(startTimestamp)
	other := fieldsToZapFields(fields...)
	zfields := []zap.Field{
		zap.Int("status", statusCode),
		zap.String("method", method),
		zap.String("path", path),
		zap.String("query", query),
		zap.Int("bodySize", bodySize),
		zap.String("ip", ip),
		zap.String("userAgent", userAgent),
		zap.Time("mvts", eventTimestamp),
		zap.String("pmvts", eventTimestamp.Format("2006-01-02T15:04:05-0700")),
		zap.Duration("latency", latency),
		zap.String("pLatency", latency.String()),
	}
	zfields = append(zfields, other...)
	if statusCode > 199 && statusCode < 300 {
		tracer.lgr.Info(
			"Request",
			zfields...,
		)
	} else {
		tracer.lgr.Error(
			"Failed Request",
			zfields...,
		)
	}
}

func (tracer *ZapTracer) TraceDependency(
	spanId string,
	dependencyType string,
	serviceName string,
	commandName string,
	success bool,
	startTimestamp time.Time,
	eventTimestamp time.Time,
	fields ...Field,

) {
	other := fieldsToZapFields(fields...)
	latency := eventTimestamp.Sub(startTimestamp)
	zfields := []zap.Field{
		zap.String("sid", spanId),
		zap.String("dependencyType", dependencyType),
		zap.String("serviceName", serviceName),
		zap.String("commandName", commandName),
		zap.Time("mvts", eventTimestamp),
		zap.String("pmvts", eventTimestamp.Format("2006-01-02T15:04:05-0700")),
		zap.Duration("latency", latency),
		zap.String("pLatency", latency.String()),
	}
	zfields = append(zfields, other...)
	if success {
		tracer.lgr.Info(
			"Dependency",
			zfields...,
		)
	} else {
		tracer.lgr.Error(
			"Failed Dependency",
			zfields...,
		)
	}
}

// TODO test how this works with 0 fields
func fieldsToZapFields(fields ...Field) (zf []zap.Field) {
	zf = make([]zap.Field, len(fields))
	for idx, field := range fields {
		zf[idx] = zap.String(field.Key, field.Value)
	}
	return
}
