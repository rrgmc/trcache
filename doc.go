// Package trcache implements strongly-typed generics-based cache library wrappers.
//
// The wrappers are highly customizable and allows accessing the underlying cache functionality if desired.
//
// A loadable (refresh) implementation is available for all caches, which provides a [RefreshCache].GetOrRefresh which
// calls a refresh function if the data was not found in [Cache].Get.
package trcache
