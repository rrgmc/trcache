module github.com/RangelReale/trcache/cache/bigcache

go 1.19

require (
	github.com/RangelReale/trcache v0.3.5
	github.com/RangelReale/trcache/mocks v0.3.5
	github.com/allegro/bigcache/v3 v3.1.0
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
