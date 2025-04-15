package exercises

func getSum(a int, b int) int {
	const MaskValue = 0xFFFFFFFF
	const Max32bit = 0x80000000

	// convert to 32-bit int representation via mask
	a = a & MaskValue
	b = b & MaskValue

	// while we have more to add
	for b != 0 {
		// get binary of carry, bit shift to apply it to next highest bit, ensure it's within mask range
		carry := ((a & b) << 1) & MaskValue

		// XOR a and b to get sum without considering the carry
		a = a ^ b

		// continue with carry
		b = carry
	}

	// return a if positive integer in 32bit range
	if a < Max32bit {
		return a
	}

	// a is negative
	return ^(a ^ MaskValue)
}
