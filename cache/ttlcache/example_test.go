package trttlcache

import (
	"context"
	"fmt"
	"time"

	"github.com/jellydator/ttlcache/v3"
)

func ExampleCache() {
	ctx := context.Background()

	tcache := ttlcache.New[string, string]()

	cache, err := New[string, string](tcache,
		WithDefaultDuration[string, string](time.Minute),
	)

	err = cache.Set(ctx, "a", "12")
	if err != nil {
		panic(err)
	}

	v, err := cache.Get(ctx, "a")
	if err != nil {
		panic(err)
	}
	fmt.Println(v)

	err = cache.Delete(ctx, "a")
	if err != nil {
		panic(err)
	}

	_, err = cache.Get(ctx, "a")
	fmt.Println(err)
	// Output:
	// 12
	// not found
}
