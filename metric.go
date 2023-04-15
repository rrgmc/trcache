package trcache

import "context"

type Metrics interface {
	Hit(ctx context.Context, name string, cacheName string, key any)
	Miss(ctx context.Context, name string, cacheName string, key any)
	Error(ctx context.Context, name string, cacheName string, key any, errorType MetricsErrorType)
}

type MetricsErrorType int

const (
	MetricsErrorTypeError MetricsErrorType = iota
	MetricsErrorTypeGet
	MetricsErrorTypeSet
	MetricsErrorTypeDecode
	MetricsErrorTypeEncode
	MetricsErrorTypeRefresh
)
