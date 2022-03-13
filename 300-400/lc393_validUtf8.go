package main

const (
	Valid1 = 0b1000_0000

	MaskNext  = 0b1100_0000
	MaskValid = 0b1000_0000

	Mask2  = 0b1110_0000
	Valid2 = 0b1100_0000

	Mask3  = 0b1111_0000
	Valid3 = 0b1110_0000

	Mask4  = 0b1111_1000
	Valid4 = 0b1111_0000
)

func validUtf8(data []int) bool {
	var ln int
	for _, v := range data {
		if ln != 0 {
			if v&MaskNext != MaskValid {
				return false
			}
			ln--
			continue
		}
		if v&Valid1 == 0 {
			continue
		}

		if v&Mask2 == Valid2 {
			ln = 1
			continue
		}
		if v&Mask3 == Valid3 {
			ln = 2
			continue
		}
		if v&Mask4 == Valid4 {
			ln = 3
			continue
		}
		return false
	}
	return ln == 0
}
