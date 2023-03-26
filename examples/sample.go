package main

import (
	"context"

	"github.com/RangelReale/trcache"
	trttlcache "github.com/RangelReale/trcache/cache/ttlcache"
	"github.com/RangelReale/trcache/chain"
	"github.com/jellydator/ttlcache/v3"
)

func main() {
	ctx := context.Background()

	sampleChain(ctx)
}

func sampleChain(ctx context.Context) {
	cache1 := trttlcache.New[string, string](ttlcache.New[string, string]())
	cache2 := trttlcache.New[string, string](ttlcache.New[string, string]())

	cache := chain.New[string, string]([]trcache.Cache[string, string]{
		cache1, cache2,
	})

	cache.Set(ctx, "a", "b")
}
