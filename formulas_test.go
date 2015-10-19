package formulas

import "testing"

func TestAbs(t *testing.T) {

	x := -19.0
	y := 0.0
	z := -10.34

	t1, err := Abs(x)
	if err != nil {
		t.Errorf("Abs Function returned an error")
	}

	if t1 != 19.000 {
		t.Errorf("Abs Function failed")
	}

	t1, err = Abs(y)
	if err != nil {
		t.Errorf("Abs Function returned an error")
	}

	if t1 != 0.0 {
		t.Errorf("Abs Function failed")
	}

	t1, err = Abs(z)
	if err != nil {
		t.Errorf("Abs Function returned an error")
	}

	if t1 != 10.34 {
		t.Errorf("Abs Function failed")
	}
}

func BenchmarkAbs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Abs(-110.0)
	}
}

func TestEXP(t *testing.T) {

	x := 100.0
	y := 0.1
	z := -5.0

	t1, err := Exp(x)

	if err != nil {
		t.Errorf("Exp Function returned an error")
	}

	if t1 != 2.6881171418161356e+43 {
		t.Errorf("Exp Function Failed ... ")
	}

	t1, err = Exp(y)

	if err != nil {
		t.Errorf("Exp Function returned an error")
	}

	if t1 != 1.1051709180756477 {
		t.Errorf("Exp Function Failed ... ")
	}

	t1, err = Exp(z)

	if err != nil {
		t.Errorf("Exp Function returned an error")
	}

	if t1 != 0.006737946999085467 {
		t.Errorf("Exp Function Failed ... ")
	}
}

func BenchmarkExp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Exp(1.0)
	}
}

func TestGCD(t *testing.T) {

	a := []int{203, 91, 77}
	b := []int{77, 91, 203}
	c := []int{5, 2}
	x := []int{24, 36}
	y := []int{5, 0}
	z := []int{15, 10, 25}

	result, err := GCD(a)

	if err != nil {
		t.Errorf("GCD Function returned an error")
	}

	if result != 7 {
		t.Errorf("GCD Function Failed ")
	}

	result, err = GCD(b)

	if err != nil {
		t.Errorf("GCD Function returned an error")
	}

	if result != 7 {
		t.Errorf("GCD Function Failed ")
	}

	result, err = GCD(c)

	if err != nil {
		t.Errorf("GCD Function returned an error")
	}

	if result != 1 {
		t.Errorf("GCD Function Failed ")
	}

	result, err = GCD(x)

	if err != nil {
		t.Errorf("GCD Function returned an error")
	}

	if result != 12 {
		t.Errorf("GCD Function Failed ")
	}

	result, err = GCD(y)

	if err != nil {
		t.Errorf("GCD Function returned an error")
	}

	if result != 5 {
		t.Errorf("GCD Function Failed ")
	}

	result, err = GCD(z)

	if err != nil {
		t.Errorf("GCD Function returned an error")
	}

	if result != 5 {
		t.Errorf("GCD Function Failed ")
	}

}

func BenchmarkGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GCD([]int{203, 91, 77})
	}
}

func TestLCM(t *testing.T) {

	a := []int{1, 5}
	b := []int{15, 10, 25}
	c := []int{1, 8, 12}
	x := []int{7, 2}

	result, err := LCM(a)

	if err != nil {
		t.Errorf("LCM Function returned an error")
	}

	if result != 5 {
		t.Errorf("LCM Function Failed ")
	}

	result, err = LCM(b)

	if err != nil {
		t.Errorf("Abs Function returned an error")
	}

	if result != 150 {
		t.Errorf("LCM Function Failed ")
	}

	result, err = LCM(c)

	if err != nil {
		t.Errorf("LCM Function returned an error")
	}

	if result != 24 {
		t.Errorf("LCM Function Failed ")
	}

	result, err = LCM(x)

	if err != nil {
		t.Errorf("LCM Function returned an error")
	}

	if result != 14 {
		t.Errorf("LCM Function Failed ")
	}

}

func BenchmarkLCM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LCM([]int{203, 91, 77})
	}
}

func TestLN(t *testing.T) {

	x := 0.5
	y := 100.00

	result, err := LN(x)

	if err != nil {
		t.Errorf("LN Function returned an error")
	}

	if result != -0.6931471805599453 {
		t.Errorf("LN Function Failed ... ")
	}

	result, err = LN(y)

	if err != nil {
		t.Errorf("LN Function returned an error")
	}

	if result != 4.605170185988092 {
		t.Errorf("LN Function Failed ... ")
	}

}

func BenchmarkLN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LN(23.98)
	}
}
