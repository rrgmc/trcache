# trcache gocache [![GoDoc](https://godoc.org/github.com/RangelReale/trcache/cache/gocache?status.png)](https://godoc.org/github.com/RangelReale/trcache/cache/gocache)

This is a [trcache](https://github.com/RangelReale/trcache) wrapper for [go-cache](https://github.com/patrickmn/go-cache).

## Info

### go-cache library

| info        |          |
|-------------|----------|
| Generics    | No       |
| Key types   | `string` |
| Value types | `any`    |
| TTL         | Yes      |

### wrapper

| info              |                   |
|-------------------|-------------------|
| Default codec     | `NewForwardCodec` |
| Default key codec | `StringKeyCodec`  |

## Installation

```shell
go get github.com/RangelReale/trcache/cache/gocache
```
