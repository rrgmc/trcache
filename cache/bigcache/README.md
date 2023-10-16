# trcache BigCache [![GoDoc](https://godoc.org/github.com/rrgmc/trcache/cache/bigcache?status.png)](https://godoc.org/github.com/rrgmc/trcache/cache/bigcache)

This is a [trcache](https://github.com/rrgmc/trcache) wrapper for [BigCache](https://github.com/allegro/bigcache). 

## Info

### BigCache library

| info        |          |
|-------------|----------|
| Generics    | No       |
| Key types   | `string` |
| Value types | `[]byte` |
| TTL         | Yes      |

### wrapper

| info              |                  |
|-------------------|------------------|
| Default codec     | `GOBCodec`       |
| Default key codec | `StringKeyCodec` |

## Installation

```shell
go get github.com/rrgmc/trcache/cache/bigcache
```
