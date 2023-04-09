# trcache FreeCache  [![GoDoc](https://godoc.org/github.com/RangelReale/trcache/cache/freecache?status.png)](https://godoc.org/github.com/RangelReale/trcache/cache/freecache)

This is a [trcache](https://github.com/RangelReale/trcache) wrapper for [FreeCache](https://github.com/coocood/freecache).

## Info

### FreeCache library

| info        |            |
|-------------|------------|
| Generics    | No         |
| Key types   | `[]byte`   |
| Value types | `[]byte`   |
| TTL         | Yes        |

### wrapper

| info              |                  |
|-------------------|------------------|
| Default codec     | `GOBCodec`       |
| Default key codec | `StringKeyCodec` |

## Installation

```shell
go get github.com/RangelReale/trcache/cache/freecache
```
