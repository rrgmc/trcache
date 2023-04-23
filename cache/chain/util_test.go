package chain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCacheLoopForward(t *testing.T) {
	cl := newCacheLoop(5, StrategyLoopOrderFORWARD)
	var clres []int
	for cl.next() {
		clres = append(clres, cl.current())
	}
	require.Equal(t, []int{0, 1, 2, 3, 4}, clres)
}

func TestCacheLoopForwardEmpty(t *testing.T) {
	cl := newCacheLoop(0, StrategyLoopOrderFORWARD)
	var clres []int
	for cl.next() {
		clres = append(clres, cl.current())
	}
	require.Equal(t, []int(nil), clres)
}

func TestCacheLoopForwardNegative(t *testing.T) {
	cl := newCacheLoop(-5, StrategyLoopOrderFORWARD)
	var clres []int
	for cl.next() {
		clres = append(clres, cl.current())
	}
	require.Equal(t, []int(nil), clres)
}

func TestCacheLoopBackward(t *testing.T) {
	cl := newCacheLoop(5, StrategyLoopOrderBACKWARD)
	var clres []int
	for cl.next() {
		clres = append(clres, cl.current())
	}
	require.Equal(t, []int{4, 3, 2, 1, 0}, clres)
}

func TestCacheLoopBackwardEmpty(t *testing.T) {
	cl := newCacheLoop(0, StrategyLoopOrderBACKWARD)
	var clres []int
	for cl.next() {
		clres = append(clres, cl.current())
	}
	require.Equal(t, []int(nil), clres)
}

func TestCacheLoopBackwardNegative(t *testing.T) {
	cl := newCacheLoop(-5, StrategyLoopOrderBACKWARD)
	var clres []int
	for cl.next() {
		clres = append(clres, cl.current())
	}
	require.Equal(t, []int(nil), clres)
}
