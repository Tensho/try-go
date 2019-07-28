package popcountloop

func PopCount(x uint64) int {
	var count int
	for i := uint(0); i < 64; i++ {
		if x & (1 << i) != 0 {
			count++
		}
	}
	return count
}
