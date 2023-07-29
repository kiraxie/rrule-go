package rrule

// Every mask is 7 days longer to handle cross-year weekly periods.
var (
	mask366         []int
	mask365         []int
	mask366day      []int
	mask365day      []int
	mask366monthDay []int
	mask365monthDay []int
	maskDay         []int
	range366        = []int{0, 31, 60, 91, 121, 152, 182, 213, 244, 274, 305, 335, 366}
	range365        = []int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334, 365}
)

func init() {
	mask366 = concat(repeat(1, 31), repeat(2, 29), repeat(3, 31),
		repeat(4, 30), repeat(5, 31), repeat(6, 30), repeat(7, 31),
		repeat(8, 31), repeat(9, 30), repeat(10, 31), repeat(11, 30),
		repeat(12, 31), repeat(1, 7))
	mask365 = concat(mask366[:59], mask366[60:])
	M29, M30, M31 := rang(1, 30), rang(1, 31), rang(1, 32)
	mask366day = concat(M31, M29, M31, M30, M31, M30, M31, M31, M30, M31, M30, M31, M31[:7])
	mask365day = concat(mask366day[:59], mask366day[60:])
	M29, M30, M31 = rang(-29, 0), rang(-30, 0), rang(-31, 0)
	mask366monthDay = concat(M31, M29, M31, M30, M31, M30, M31, M31, M30, M31, M30, M31, M31[:7])
	mask365monthDay = concat(mask366monthDay[:31], mask366monthDay[32:])
	for i := 0; i < 55; i++ {
		maskDay = append(maskDay, []int{0, 1, 2, 3, 4, 5, 6}...)
	}
}
