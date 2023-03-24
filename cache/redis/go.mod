module github.com/RangelReale/trcache/cache/redis

go 1.19

require (
	github.com/RangelReale/trcache v0.0.0
	github.com/redis/go-redis/v9 v9.0.2
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
)

replace github.com/RangelReale/trcache => ../..