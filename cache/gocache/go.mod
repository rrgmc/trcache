module github.com/rrgmc/trcache/cache/gocache

go 1.19

require (
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/rrgmc/trcache v0.15.0
	github.com/stretchr/testify v1.8.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/rrgmc/trcache => ../..
