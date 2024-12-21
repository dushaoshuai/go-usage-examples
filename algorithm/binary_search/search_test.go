package binary_search

import (
	"cmp"
	"testing"
)

func TestFind(t *testing.T) {
	str1 := []string{"foo"}
	str2 := []string{"ab", "ca"}
	str3 := []string{"mo", "qo", "vo"}
	str4 := []string{"ab", "ad", "ca", "xy"}

	// slice with repeating elements
	strRepeats := []string{"ba", "ca", "da", "da", "da", "ka", "ma", "ma", "ta"}

	// slice with all element equal
	strSame := []string{"xx", "xx", "xx"}

	tests := []struct {
		data      []string
		target    string
		wantPos   int
		wantFound bool
	}{
		{[]string{}, "foo", -1, false},
		{[]string{}, "", -1, false},

		{str1, "foo", 0, true},
		{str1, "bar", -1, false},
		{str1, "zx", 0, false},

		{str2, "aa", -1, false},
		{str2, "ab", 0, true},
		{str2, "ad", 0, false},
		{str2, "ca", 1, true},
		{str2, "ra", 1, false},

		{str3, "bb", -1, false},
		{str3, "mo", 0, true},
		{str3, "nb", 0, false},
		{str3, "qo", 1, true},
		{str3, "tr", 1, false},
		{str3, "vo", 2, true},
		{str3, "xr", 2, false},

		{str4, "aa", -1, false},
		{str4, "ab", 0, true},
		{str4, "ac", 0, false},
		{str4, "ad", 1, true},
		{str4, "ax", 1, false},
		{str4, "ca", 2, true},
		{str4, "cc", 2, false},
		{str4, "dd", 2, false},
		{str4, "xy", 3, true},
		{str4, "zz", 3, false},

		{strRepeats, "da", 4, true},
		{strRepeats, "db", 4, false},
		{strRepeats, "ma", 7, true},
		{strRepeats, "mb", 7, false},

		{strSame, "xx", 2, true},
		{strSame, "ab", -1, false},
		{strSame, "zz", 2, false},
	}

	for _, tt := range tests {
		t.Run(tt.target, func(t *testing.T) {
			cmpFunc := func(i int) int {
				return cmp.Compare(tt.target, tt.data[i])
			}

			pos, found := FindMaxLte(len(tt.data), cmpFunc)
			if pos != tt.wantPos || found != tt.wantFound {
				t.Errorf("FindMaxLte got (%v, %v), want (%v, %v)", pos, found, tt.wantPos, tt.wantFound)
			}
		})
	}
}
