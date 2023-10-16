# trcache gocache [![GoDoc](https://godoc.org/github.com/rrgmc/trcache/cache/gocache?status.png)](https://godoc.org/github.com/rrgmc/trcache/cache/gocache)

This is a [trcache](https://github.com/rrgmc/trcache) wrapper for [go-cache](https://github.com/patrickmn/go-cache).

## Info

### go-cache library

| info        |          |
|-------------|----------|
| Generics    | No       |
| Key types   | `string` |
| Value types | `any`    |
| TTL         | Yes      |

### wrapper

| info              |                  |
|-------------------|------------------|
| Default codec     | `ForwardCodec`   |
| Default key codec | `StringKeyCodec` |

## Installation

```shell
go get github.com/rrgmc/trcache/cache/gocache
```
