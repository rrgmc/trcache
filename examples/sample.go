package main

import (
	"context"

	"github.com/RangelReale/trcache"
	"github.com/RangelReale/trcache/cache/ttlcache"
	"github.com/RangelReale/trcache/chain"
)

func main() {
	ctx := context.Background()

	sampleChain(ctx)
}

func sampleChain(ctx context.Context) {
	cache1 := trttlcache.NewDefault[string, string]()
	cache2 := trttlcache.NewDefault[string, string]()

	cache := chain.New[string, string]([]trcache.Cache[string, string]{
		cache1, cache2,
	})

	cache.Set(ctx, "a", "b")
}
