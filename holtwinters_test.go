package holtwinters

import (
	"testing"
)

func TestHoltwinters(t *testing.T) {
	y := [...]float64{362, 385, 432, 341, 382, 409, 498, 387, 473, 513,
		582, 474, 544, 582, 681, 557, 628, 707, 773, 592, 627, 725,
		854, 661}
	period := 4
	m := 4

	alpha := 0.5
	beta := 0.4
	gamma := 0.6

	prediction, _ := Forecast(y[:], alpha, beta,
		gamma, period, m)
	expected := [...]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0,
		594.8043646513713, 357.12171044215734, 410.9203094983815,
		444.67743912921156, 550.9296957593741, 421.1681718160631,
		565.905732450577, 639.2910221068818, 688.8541669002238,
		532.7122406111591, 620.5492369959037, 668.5662327429854,
		773.5946568453546, 629.0602103529998, 717.0290609530134,
		836.4643466657625, 884.1797655866865, 617.6686414831381,
		599.1184450128665, 733.227872348479, 949.0708357438998,
		748.6618488792186}
	if Compare(expected[:], prediction) != 0 {
		t.Fatal("failed")
	}
}

func Compare(a, b []float64) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}
	return 0
}
