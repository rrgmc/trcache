package trmap_test

import (
	"context"
	"fmt"

	"github.com/RangelReale/trcache"
	trmap "github.com/RangelReale/trcache/cache/map"
)

func ExampleCache() {
	ctx := context.Background()

	cache, err := trmap.NewDefault[string, string]()

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

func ExampleRefreshCache() {
	ctx := context.Background()

	cache, err := trmap.NewRefreshDefault[string, string](
		trcache.WithDefaultRefreshFunc[string, string](func(ctx context.Context, key string, options trcache.RefreshFuncOptions) (string, error) {
			return fmt.Sprintf("abc%d", options.Data), nil
		}),
	)
	if err != nil {
		panic(err)
	}

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

	v, err = cache.GetOrRefresh(ctx, "b", trcache.WithRefreshData[string, string](123))
	if err != nil {
		panic(err)
	}
	fmt.Println(v)

	// Output:
	// 12
	// not found
	// abc123
}
