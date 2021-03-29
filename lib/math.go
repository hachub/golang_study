package math

func Stat(ix []float64, iy []float64) float64 {
	var stat float64 = 0.0
	for _, x := range ix {
		for _, y := range iy {
			if y != -1 {
				stat += (1 + x) / (1 + y)
			} else {
				stat += (1 + x) / (1 - y)
			}
		}
	}
	return stat
}
