module github.com/RangelReale/trcache/cache/gocache

go 1.19

require (
	github.com/RangelReale/trcache v0.10.0
	github.com/RangelReale/trcache/mocks v0.10.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/stretchr/testify v1.8.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/RangelReale/trcache => ../..

replace github.com/RangelReale/trcache/mocks => ../../mocks
