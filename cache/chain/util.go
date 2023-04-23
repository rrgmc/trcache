package chain

type cacheLoop struct {
	cacheLen  int
	loopOrder StrategyLoopOrder
	index     int
}

func newCacheLoop(cacheLen int, loopOrder StrategyLoopOrder) *cacheLoop {
	ret := &cacheLoop{
		cacheLen:  cacheLen,
		loopOrder: loopOrder,
	}
	ret.init()
	return ret
}

func (c *cacheLoop) init() {
	c.index = -1
}

func (c *cacheLoop) current() int {
	return c.index
}

func (c *cacheLoop) next() bool {
	if c.index < 0 {
		if c.cacheLen <= 0 {
			return false
		}
		if c.loopOrder == StrategyLoopOrderBACKWARD {
			c.index = c.cacheLen - 1
		} else {
			c.index = 0
		}
	} else {
		if c.loopOrder == StrategyLoopOrderBACKWARD {
			if c.index == 0 {
				return false
			}
			c.index--
		} else {
			if c.index == c.cacheLen-1 {
				return false
			}
			c.index++
		}
	}
	return true
}
