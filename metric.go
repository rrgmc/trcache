package trcache

import "context"

type Metrics interface {
	Hit(ctx context.Context, name string)
	Miss(ctx context.Context, name string)
	Error(ctx context.Context, name string, errorType MetricsErrorType)
}

type MetricsErrorType int

const (
	MetricsErrorTypeError MetricsErrorType = iota
	MetricsErrorTypeGet
	MetricsErrorTypePut
	MetricsErrorTypeDecode
	MetricsErrorTypeEncode
	MetricsErrorTypeRefresh
)
