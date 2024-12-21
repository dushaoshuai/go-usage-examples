package binary_search

// FindMaxLte is a complement to sort.Find.
//
// FindMaxLte uses binary search to find and return the biggest index i in [0, n)
// at which cmp(i) >= 0. If there is no such index i, FindMaxLte returns i = -1.
// The found result is true if i < n and cmp(i) == 0.
// FindMaxLte calls cmp(i) only for i in the range [0, n).
//
// To permit binary search, FindMaxLte requires that cmp(i) > 0 for a leading
// prefix of the range, cmp(i) == 0 in the middle, and cmp(i) < 0 for
// the final suffix of the range. (Each subrange could be empty.)
// The usual way to establish this condition is to interpret cmp(i)
// as a comparison of a desired target value t against entry i in an
// underlying indexed data structure x, returning <0, 0, and >0
// when t < x[i], t == x[i], and t > x[i], respectively.
//
// For example, to look for a particular string in a sorted, random-access
// list of strings:
//
//	i, found := sort.FindMaxLte(x.Len(), func(i int) int {
//	    return strings.Compare(target, x.At(i))
//	})
//	if found {
//	    fmt.Printf("found %s at entry %d\n", target, i)
//	} else {
//	    fmt.Printf("%s not found, would insert at %d", target, i)
//	}
func FindMaxLte(n int, cmp func(int) int) (i int, found bool) {
	// Define cmp(-1) >= 0 and cmp(n) < 0
	// Invariant: cmp(i) >= 0, cmp(j+1) < 0

	if n == 0 {
		return -1, false
	}

	i, j := 0, n-1
	for i+1 < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if cmp(h) < 0 {
			j = h - 1 // preserves cmp(j+1) < 0
		} else {
			i = h // preserves cmp(i) >= 0
		}
	}

	// i <= j, cmp(i) >= 0 and cmp(j+1) < 0
	if cmp(j) >= 0 {
		return j, cmp(j) == 0
	} else if cmp(i) >= 0 {
		return i, cmp(i) == 0
	} else {
		return -1, false
	}
}
