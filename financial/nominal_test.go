package financial

import "testing"

// TestFutureValue function tests FutureValue function
func TestNominal(t *testing.T) {

	// effectrate : Nominal interest rate
	// npery       :  Number of compounding periods per year
	effectrate := 5.3543
	npery := 4.0

	if Nominal(npery, effectrate) != 0.052500319868356016 {
		t.Errorf("Effect Function Failed ... ")
	}
}
