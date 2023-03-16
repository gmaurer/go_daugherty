package offense

import "testing"

//func calculateOnBasePlusSlugging(batterPtr *Batter) float64 {
//	return toFixed(batterPtr.OnBasePercentage+batterPtr.SluggingPercentage, 3)
//}

func TestCalculateOnBasePlusSlugging(t *testing.T) {
	testBatter := Batter{"troutmi01", 2014, 1, "LAA", "AL", 157, 602, 115, 173, 39, 9, 36, 111, 16, 2, 83, 184, 6, 10, 0, 10, 6, 0.250, .364, .536, 0}
	result := calculateOnBasePlusSlugging(&testBatter)
	expected := 0.900

	if result != expected {
		t.Errorf("Returned %f, but expected %f", result, expected)
	}
}

func TestCalculateBattingAverage(t *testing.T) {
	testBatter := Batter{"troutmi01", 2014, 1, "LAA", "AL", 157, 602, 115, 173, 39, 9, 36, 111, 16, 2, 83, 184, 6, 10, 0, 10, 6, 0, 0, 0, 0}
	testBatter2 := Batter{"cabremi01", 2013, 1, "DET", "AL", 148, 555, 103, 193, 26, 1, 44, 137, 3, 0, 90, 94, 19, 5, 0, 2, 19, 0, 0, 0, 0}

	var cases = []struct {
		name     string
		input    *Batter
		expected float64
	}{
		{"Mike Trout", &testBatter, .287},
		{"Miguel Cabrera", &testBatter2, .347},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := calculateBattingAverage(tc.input)
			if result != tc.expected {
				t.Errorf("Return %f, but wanted %f", result, tc.expected)
			}
		})
	}
}
