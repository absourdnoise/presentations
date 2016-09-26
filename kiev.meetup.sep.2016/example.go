package slides

import (
	"math"
	"sync/atomic"
)

func sliceSum(s []int) (ret int) {
	for _, v := range s {
		ret += v
	}
	return
}

func addF64(value *uint64, delta float64) {
	for {
		cur := atomic.LoadUint64(value)
		curVal := math.Float64frombits(cur)
		nxtVal := curVal + delta
		nxt := math.Float64bits(nxtVal)
		if atomic.CompareAndSwapUint64(value, cur, nxt) {
			return
		}
	}
}

func bytesEq(s1, s2 []byte) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
