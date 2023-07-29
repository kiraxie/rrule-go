package rrule

// Every mask is 7 days longer to handle cross-year weekly periods.
var (
	m366MASK     []int
	m365MASK     []int
	mDAY366MASK  []int
	mDAY365MASK  []int
	nMDAY366MASK []int
	nMDAY365MASK []int
	wDAYMASK     []int
	m366RANGE    = []int{0, 31, 60, 91, 121, 152, 182, 213, 244, 274, 305, 335, 366}
	m365RANGE    = []int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334, 365}
)

func init() {
	m366MASK = concat(repeat(1, 31), repeat(2, 29), repeat(3, 31),
		repeat(4, 30), repeat(5, 31), repeat(6, 30), repeat(7, 31),
		repeat(8, 31), repeat(9, 30), repeat(10, 31), repeat(11, 30),
		repeat(12, 31), repeat(1, 7))
	m365MASK = concat(m366MASK[:59], m366MASK[60:])
	M29, M30, M31 := rang(1, 30), rang(1, 31), rang(1, 32)
	mDAY366MASK = concat(M31, M29, M31, M30, M31, M30, M31, M31, M30, M31, M30, M31, M31[:7])
	mDAY365MASK = concat(mDAY366MASK[:59], mDAY366MASK[60:])
	M29, M30, M31 = rang(-29, 0), rang(-30, 0), rang(-31, 0)
	nMDAY366MASK = concat(M31, M29, M31, M30, M31, M30, M31, M31, M30, M31, M30, M31, M31[:7])
	nMDAY365MASK = concat(nMDAY366MASK[:31], nMDAY366MASK[32:])
	for i := 0; i < 55; i++ {
		wDAYMASK = append(wDAYMASK, []int{0, 1, 2, 3, 4, 5, 6}...)
	}
}
