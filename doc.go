// Package trcache implements strongly-typed generics-based cache library wrappers.
//
// The wrappers are highly customizable and allows accessing the underlying cache functionality if desired.
//
// A loadable (refresh) implementation is available for all caches, which provides a "GetOrRefresh" method which
// calls a refresh function if the data was not found in "Get".
//
// Example:
//
//	ctx := context.Background()
//
//	cache, err := trmap.NewDefault[string, string]()
//
//	err = cache.Set(ctx, "a", "12")
//	if err != nil {
//		panic(err)
//	}
//
//	v, err := cache.Get(ctx, "a")
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(v)
//
//	err = cache.Delete(ctx, "a")
//	if err != nil {
//		panic(err)
//	}
//
//	_, err = cache.Get(ctx, "a")
//	fmt.Println(err)
//	// Output:
//	// 12
//	// not found
//
// Example with refresh:
//
//	ctx := context.Background()
//
//	cache, err := trmap.NewRefreshDefault[string, string](
//		trcache.WithDefaultRefreshFunc[string, string](func(ctx context.Context, key string, options trcache.RefreshFuncOptions) (string, error) {
//			return fmt.Sprintf("abc%d", options.Data), nil
//		}),
//	)
//	if err != nil {
//		panic(err)
//	}
//
//	err = cache.Set(ctx, "a", "12")
//	if err != nil {
//		panic(err)
//	}
//
//	v, err := cache.Get(ctx, "a")
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(v)
//
//	err = cache.Delete(ctx, "a")
//	if err != nil {
//		panic(err)
//	}
//
//	_, err = cache.Get(ctx, "a")
//	fmt.Println(err)
//
//	v, err = cache.GetOrRefresh(ctx, "b", trcache.WithRefreshData[string, string](123))
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(v)
//
//	// Output:
//	// 12
//	// not found
//	// abc123
package trcache
