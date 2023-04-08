[![GoDoc](https://godoc.org/github.com/RangelReale/trcache?status.png)](https://godoc.org/github.com/RangelReale/trcache)

trcache
=======

Package trcache implements strongly-typed generics-based cache library wrappers.

The wrappers are highly customizable and allows accessing the underlying cache functionality if desired.

A loadable (refresh) implementation is available for all caches, which provides a "GetOrRefresh" method which
calls a refresh function if the data was not found in "Get".

## Installation

```go
go get github.com/RangelReale/trcache
```

## Examples

```go
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
```

With refresh:

```go
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
```

## Implementations

- Go `map`: [https://pkg.go.dev/github.com/RangelReale/trcache/cache/map](https://pkg.go.dev/github.com/RangelReale/trcache/cache/map)
- Chain (uses an ordered list of caches): [https://pkg.go.dev/github.com/RangelReale/trcache/cache/chain](https://pkg.go.dev/github.com/RangelReale/trcache/cache/chain)
- bigcache: [https://pkg.go.dev/github.com/RangelReale/trcache/cache/bigcache](https://pkg.go.dev/github.com/RangelReale/trcache/cache/bigcache)
- freecache: [https://pkg.go.dev/github.com/RangelReale/trcache/cache/freecache](https://pkg.go.dev/github.com/RangelReale/trcache/cache/freecache)
- gocache: [https://pkg.go.dev/github.com/RangelReale/trcache/cache/gocache](https://pkg.go.dev/github.com/RangelReale/trcache/cache/gocache)
- ristretto: [https://pkg.go.dev/github.com/RangelReale/trcache/cache/ristretto](https://pkg.go.dev/github.com/RangelReale/trcache/cache/ristretto)
- ttlcache: [https://pkg.go.dev/github.com/RangelReale/trcache/cache/ttlcache](https://pkg.go.dev/github.com/RangelReale/trcache/cache/ttlcache)
- goredis: [https://github.com/RangelReale/trcache-goredis](https://github.com/RangelReale/trcache-goredis)
- rueidis: [https://github.com/RangelReale/trcache-rueidis](https://github.com/RangelReale/trcache-rueidis)

## Author

Rangel Reale (rangelreale@gmail.com)
