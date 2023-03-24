package trcache

type ChainCache[K comparable, V any] struct {
}

func NewChainCache[K comparable, V any]() *ChainCache[K, V] {
	return &ChainCache[K, V]{}
}
