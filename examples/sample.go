package main

import (
	"context"

	trttlcache "github.com/RangelReale/trcache/cache/ttlcache"
	"github.com/jellydator/ttlcache/v3"
)

func main() {
	ctx := context.Background()

	c := trttlcache.New[string, string](ttlcache.New[string, string]())

	c.Set(ctx, "a", "b")
}
