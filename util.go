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

func Crossover(series1, series2 *CBuf) bool {
	if series1.Size() < 3 || series2.Size() < 3 {
		return false
	}

	return series1.NthOldest(2) <= series2.NthOldest(2) && series1.NthOldest(1) > series2.NthOldest(1)
}

func Crossunder(series1, series2 *CBuf) bool {
	if series1.Size() < 3 || series2.Size() < 3 {
		return false
	}

	return series1.NthOldest(2) > series2.NthOldest(2) && series1.NthOldest(1) <= series2.NthOldest(1)
}
