package tart

type Cmo struct {
	n     int64
	su    *Ema
	sd    *Ema
	prevC float64
	sz    int64
}

func NewCmo(n int64) *Cmo {
	k := 1.0 / float64(n)
	return &Cmo{
		n:     n,
		su:    NewEma(n, k),
		sd:    NewEma(n, k),
		prevC: 0,
		sz:    0,
	}
}

func (c *Cmo) Update(v float64) float64 {
	d := v - c.prevC
	c.prevC = v
	c.sz++

	if c.sz == 1 {
		return 0
	}

	var asu, asd float64
	if d > 0 {
		asu = c.su.Update(d)
		asd = c.sd.Update(0)
	} else {
		asu = c.su.Update(0)
		asd = c.sd.Update(-d)
	}

	sum := asu + asd
	if almostZero(sum) {
		return 0
	}

	return (asu - asd) / sum * 100.0
}

func CmoArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	c := NewCmo(n)
	for i, v := range in {
		out[i] = c.Update(v)
	}

	return out
}