package spreadsheet

import "testing"

// TestFutureValue function tests FutureValue function
func TestEffect(t *testing.T) {

	// nominalrate : Nominal interest rate
	// npery       :  Number of compounding periods per year
	nominalrate := 5.25
	npery := 4.0

	if Effect(npery, nominalrate) != 0.05354266737075841 {
		t.Errorf("Effect Function Failed ... ")
	}

	nominalrate = 8.5
	npery = 12.0

	if Effect(npery, nominalrate) != 0.08839090589263554 {
		t.Errorf("Effect Function Failed ... ")
	}
}
