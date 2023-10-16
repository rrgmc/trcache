# trcache [![GoDoc](https://godoc.org/github.com/rrgmc/trcache?status.png)](https://godoc.org/github.com/rrgmc/trcache)

Package trcache implements strongly-typed generics-based cache library wrappers.

```go
type Cache[K comparable, V any] interface {
	Name() string
	Get(ctx context.Context, key K, options ...GetOption) (V, error)
	Set(ctx context.Context, key K, value V, options ...SetOption) error
	Delete(ctx context.Context, key K, options ...DeleteOption) error
}

type RefreshCache[K comparable, V any] interface {
	Cache[K, V]
	GetOrRefresh(ctx context.Context, key K, options ...RefreshOption) (V, error)
}
```

The wrappers are highly customizable and allows accessing the underlying cache functionality if desired.

A loadable (refresh) implementation is available for all caches, which provides a "GetOrRefresh" method which
calls a refresh function if the data was not found in "Get".

A [chain](https://pkg.go.dev/github.com/rrgmc/trcache/cache/chain) implementation is available to chain
multiples caches in sequence, for example to use an in-memory cache with a lower TTL before calling a Redis backend.

## Installation

```go
go get github.com/rrgmc/trcache
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

- Go `map`: [github.com/rrgmc/trcache/cache/map](https://pkg.go.dev/github.com/rrgmc/trcache/cache/map)
- Chain (ordered list of caches): [github.com/rrgmc/trcache/cache/chain](https://pkg.go.dev/github.com/rrgmc/trcache/cache/chain)
- bigcache: [github.com/rrgmc/trcache/cache/bigcache](https://pkg.go.dev/github.com/rrgmc/trcache/cache/bigcache)
- freecache: [github.com/rrgmc/trcache/cache/freecache](https://pkg.go.dev/github.com/rrgmc/trcache/cache/freecache)
- gocache: [github.com/rrgmc/trcache/cache/gocache](https://pkg.go.dev/github.com/rrgmc/trcache/cache/gocache)
- ristretto: [github.com/rrgmc/trcache/cache/ristretto](https://pkg.go.dev/github.com/rrgmc/trcache/cache/ristretto)
- ttlcache: [github.com/rrgmc/trcache/cache/ttlcache](https://pkg.go.dev/github.com/rrgmc/trcache/cache/ttlcache)
- goredis: [github.com/rrgmc/trcache-goredis](https://github.com/rrgmc/trcache-goredis)
- rueidis: [github.com/rrgmc/trcache-rueidis](https://github.com/rrgmc/trcache-rueidis)

## Author

Rangel Reale (rangelreale@gmail.com)
