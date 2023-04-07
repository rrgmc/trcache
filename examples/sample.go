package main

import (
	"context"
	"fmt"

	"github.com/RangelReale/trcache"
	"github.com/RangelReale/trcache/cache/chain"
	trttlcache "github.com/RangelReale/trcache/cache/ttlcache"
)

func main() {
	ctx := context.Background()

	sampleChain(ctx)
}

func sampleChain(ctx context.Context) {
	cache1, err := trttlcache.NewDefault[string, string]()
	if err != nil {
		panic(err)
	}

	cache2, err := trttlcache.NewDefault[string, string]()
	if err != nil {
		panic(err)
	}

	cache, err := chain.New[string, string]([]trcache.Cache[string, string]{
		cache1, cache2,
	})
	if err != nil {
		panic(err)
	}

	err = cache.Set(ctx, "a", "b")
	if err != nil {
		panic(err)
	}

	v, err := cache.Get(ctx, "a")
	if err != nil {
		panic(err)
	}
	fmt.Print(v)
}
