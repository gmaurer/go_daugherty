package offense

import "testing"

func FuzzConvertStringToStat(f *testing.F) {
	f.Add("0")
	f.Add("NA")
	f.Add("123")

	f.Fuzz(func(t *testing.T, value string) {
		convertStringToStat(value)
	})
}

func FuzzToFixed(f *testing.F) {
	f.Add(4.452, 3)
	f.Add(1.110, 2)
	f.Add(0.001, 5)

	f.Fuzz(func(t *testing.T, value float64, prec int) {
		toFixed(value, prec)
	})
}
