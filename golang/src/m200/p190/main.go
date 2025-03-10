package p190

func reverseBits(num uint32) uint32 {

	r := uint32(0)
	for range 32 {
		r <<= 1
		r |= num & 0b1
		num >>= 1
	}
	return r
}
