module github.com/RangelReale/trcache/example

go 1.19

require (
	github.com/RangelReale/trcache/cache/ttlcache v0.1.1
	github.com/jellydator/ttlcache/v3 v3.0.1
)

require (
	github.com/RangelReale/trcache v0.1.1 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
)

replace github.com/RangelReale/trcache => ../

replace github.com/RangelReale/trcache/cache/ttlcache => ../cache/ttlcache
