// 100% beat
func reverseBits(num uint32) uint32 {
	// i, 32-i-1
	var ret uint32
	for i := 0; i < 16; i++ {
		ret |= num >> i & 0x1 << (32 - i - 1)
		ret |= num >> (32 - i - 1) & 0x1 << i
	}
	return ret
}
