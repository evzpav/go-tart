package tart

func almostZero(v float64) bool {
	return v > -0.00000000000001 && v < 0.00000000000001
}

// math.Max() does some extra we don'tc care (overhead)
func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

type Series interface {
	Size() int64
	NthNewest(int64) float64
}

func Crossover(series1, series2 Series) bool {
	if series1.Size() < 2 || series2.Size() < 2 {
		return false
	}

	s1_2 := series1.NthNewest(1)
	s2_2 := series2.NthNewest(1)
	s1_1 := series1.NthNewest(0)
	s2_1 := series2.NthNewest(0)

	return s1_2 <= s2_2 && s1_1 > s2_1
}

func Crossunder(series1, series2 Series) bool {
	if series1.Size() < 2 || series2.Size() < 2 {
		return false
	}

	s1_2 := series1.NthNewest(1)
	s2_2 := series2.NthNewest(1)
	s1_1 := series1.NthNewest(0)
	s2_1 := series2.NthNewest(0)

	return s1_2 > s2_2 && s1_1 <= s2_1
}
