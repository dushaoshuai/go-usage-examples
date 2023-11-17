package nnkosmos

// see slices.BinarySearch for details
func binarySearch(x []int, target int) int {
	i, j := 0, len(x)
	for i < j {
		h := int(uint(i+j) >> 1)
		if x[h] < target {
			i = h + 1
		} else {
			j = h
		}
	}
	if i < len(x) && x[i] == target {
		return i
	}
	return -1
}
