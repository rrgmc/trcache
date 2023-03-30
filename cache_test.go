package trcache

import "testing"

//go:generate go run github.com/vektra/mockery/v2@v2.23.1 --name Cache
//go:generate go run github.com/vektra/mockery/v2@v2.23.1 --name RefreshCache
//go:generate go run github.com/vektra/mockery/v2@v2.23.1 --name Validator
//go:generate go run github.com/vektra/mockery/v2@v2.23.1 --name Codec
//go:generate go run github.com/vektra/mockery/v2@v2.23.1 --name CacheRefreshFunc

func TestCache(t *testing.T) {
	
}
