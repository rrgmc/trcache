package trmap

import (
	"context"
	"fmt"
)

func ExampleCache() {
	ctx := context.Background()

	cache, err := NewDefault[string, string]()

	err = cache.Set(ctx, "a", "12")
	if err != nil {
		panic(err)
	}

	v, err := cache.Get(ctx, "a")
	if err != nil {
		panic(err)
	}
	fmt.Println(v)

	err = cache.Delete(ctx, "a")
	if err != nil {
		panic(err)
	}

	_, err = cache.Get(ctx, "a")
	fmt.Println(err)
	// Output:
	// 12
	// not found
}
