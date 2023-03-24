package trcache

import "context"

type Metric interface {
	Hit(ctx context.Context, name string)
	Miss(ctx context.Context, name string)
	Error(ctx context.Context, name string, errorType MetricErrorType)
}

type MetricErrorType int

const (
	MetricErrorTypeGet MetricErrorType = iota
	MetricErrorTypePut
	MetricErrorTypeDecode
	MetricErrorTypeEncode
	MetricErrorTypeRefresh
)
