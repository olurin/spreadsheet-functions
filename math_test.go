package formulas

import (
	"log"
	"testing"
)

func TestABS(t *testing.T) {
	x := 9.0
	y := -10.0
	z := -10.34

	t1, err := ABS(x)
	if err != nil {
		t.Errorf("Abs Function returned an error")
	}

	if t1 != 9 {
		t.Errorf("Abs Function failed")
	}

	t1, err = ABS(y)
	if err != nil {
		t.Errorf("Abs Function returned an error")
	}

	if t1 != 10.0 {
		t.Errorf("Abs Function failed")
	}

	t1, err = ABS(z)
	if err != nil {
		t.Errorf("Abs Function returned an error")
	}

	if t1 != 10.34 {
		t.Errorf("Abs Function failed")
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

func TestGCD(t *testing.T) {

	a := []int{203, 91, 77}
	// b := []int{77, 91, 203}
	c := []int{5, 2}
	x := []int{24, 36}
	y := []int{5, 0}
	z := []int{15, 10, 25}

	result, err := GCD(a...)

	if err != nil {
		t.Errorf("GCD Function returned an error")
	}

	if result != 7 {
		t.Errorf("GCD Function Failed ")
	}

	result, err = GCD(77, 91, 203)

	if err != nil {
		t.Errorf("GCD Function returned an error")
	}

	if result != 7 {
		t.Errorf("GCD Function Failed ")
	}

	result, err = GCD(c...)

	if err != nil {
		t.Errorf("GCD Function returned an error")
	}

	if result != 1 {
		t.Errorf("GCD Function Failed ")
	}

	result, err = GCD(x...)

	if err != nil {
		t.Errorf("GCD Function returned an error")
	}

	if result != 12 {
		t.Errorf("GCD Function Failed ")
	}

	result, err = GCD(y...)

	if err != nil {
		t.Errorf("GCD Function returned an error")
	}

	if result != 5 {
		t.Errorf("GCD Function Failed ")
	}

	result, err = GCD(z...)

	if err != nil {
		t.Errorf("GCD Function returned an error")
	}

	if result != 5 {
		t.Errorf("GCD Function Failed ")
	}

}
func TestLCM(t *testing.T) {

	a := []int{1, 5}
	b := []int{15, 10, 25}
	c := []int{1, 8, 12}
	x := []int{7, 2}

	result, err := LCM(a...)

	if err != nil {
		t.Errorf("LCM Function returned an error")
	}

	if result != 5 {
		t.Errorf("LCM Function Failed ")
	}

	result, err = LCM(b...)

	if err != nil {
		t.Errorf("Abs Function returned an error")
	}

	if result != 150 {
		t.Errorf("LCM Function Failed ")
	}

	result, err = LCM(c...)

	if err != nil {
		t.Errorf("LCM Function returned an error")
	}

	if result != 24 {
		t.Errorf("LCM Function Failed ")
	}

	result, err = LCM(x...)

	if err != nil {
		t.Errorf("LCM Function returned an error")
	}

	if result != 14 {
		t.Errorf("LCM Function Failed ")
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

func TestLog10(t *testing.T) {

	x := 86.0
	y := 100.0
	z := 10.0

	result, err := Log10(x)
	if err != nil {
		t.Errorf("Log10 Function returned an error")
	}

	if result != 1.9344984512435675 {
		t.Errorf("Log10 Function Failed ... ")
	}

	result, err = Log10(y)

	if err != nil {
		t.Errorf("Log10 Function returned an error")
	}

	if result != 2 {
		t.Errorf("Log10 Function Failed ... ")
	}

	result, err = Log10(z)

	if err != nil {
		t.Errorf("Log10 Function returned an error")
	}

	if result != 1 {
		t.Errorf("Log10 Function Failed ... ")
	}
}

func TestMod(t *testing.T) {

	result, err := Mod(3, 2)

	if err != nil {
		t.Errorf("Mod Function returned an error")
	}

	if result != 1 {
		t.Errorf("Mod Function Failed Mod(3, 2) ... ")
	}

	result, err = Mod(-3, 2)

	if err != nil {
		t.Errorf("Mod Function returned an error")
	}

	if result != 1 {
		t.Errorf("Mod Function Failed  Mod(-3, 2) ... ")
	}

	result, err = Mod(3, -2)

	if err != nil {
		t.Errorf("Mod Function returned an error")
	}

	if result != -1 {
		t.Errorf("Mod Function Failed  Mod(3, -2) ... ")
	}

	result, err = Mod(-3, -2)

	if err != nil {
		t.Errorf("Mod Function returned an error")
	}

	if result != -1 {
		t.Errorf("Mod Function Failed Mod(-3, -2) ... ")
	}
}

func TestPower(t *testing.T) {

	p, err := Power(10, 4)
	if p != 10000 || err != nil {
		t.Errorf("Product Function Failed ... ")
	}

	p, err = Power(2, 4)
	if p != 16 || err != nil {
		log.Println(p)
		t.Errorf("Product Function Failed ... ")
	}

	p, err = Power(5, 2)
	if p != 25 || err != nil {
		t.Errorf("Product Function Failed ... ")
	}
}

func TestProduct(t *testing.T) {

	y := []float64{100, 10, 45}

	p, err := Product(45, 2, 23)
	if p != 2070 || err != nil {
		t.Errorf("Product Function Failed ... ")
	}

	p, err = Product(y...)
	if p != 45000 || err != nil {
		log.Println(p)
		t.Errorf("Product Function Failed ... ")
	}
}

func TestQuotient(t *testing.T) {

	result, err := Quotient(5, 2)

	if result != 2 || err != nil {
		t.Errorf("Quotient Function Failed ... ")
	}

	result, err = Quotient(4.5, 3.1)

	if result != 1 || err != nil {
		t.Errorf("Quotient Function Failed ... ")
	}

	result, err = Quotient(-10, 3)

	if result != -3 || err != nil {
		t.Errorf("Quotient Function Failed ... ")
	}
}

func TestRound(t *testing.T) {

	if Round(4.05, 0) != 4 {
		t.Errorf("Round function in Math library failed...")
	}
	if Round(4.05, 2) != 4.05 {
		t.Errorf("Round function in Math library failed...")
	}
	if Round(4.05, 1) != 4.1 {
		t.Errorf("Round function in Math library failed...")
	}

	if Round(4.05, 10) != 4.05 {
		t.Errorf("Round function in Math library failed...")
	}
}

func TestSqrt(t *testing.T) {
	x, err := Sqrt(81)

	if x != 9 && err != nil {
		t.Errorf("Sqrt Function Failed ... ")
	}

	x1, err1 := Sqrt(-81)

	if x1 != 0.0 && err1 == nil {
		t.Errorf("Sqrt Function Failed ... ")
	}
}

func TestSum(t *testing.T) {

	x, err := Sum(12, 34, 65, 34, 23)

	if x != 168 && err != nil {
		t.Errorf("Sum Function Failed ... ")
	}

	x, err = Sum(12, 23)

	if x != 35 && err != nil {
		t.Errorf("Sum Function Failed ... ")
	}
	x, err = Sum(-12, 23)
	if x != 11 && err != nil {
		t.Errorf("Sum Function Failed ... ")
	}
}

func TestSqrtPI(t *testing.T) {
	x, err := SqrtPI(5)
	if x != 3.9633272976060088 || err != nil {
		t.Errorf("SqrtPI Function Failed ... ")
	}

	x, err = SqrtPI(0.2)

	if x != 0.7926654595212018 || err != nil {
		t.Errorf("SqrtPI Function Failed ... ")
	}

	x, err = SqrtPI(100)

	if x != 17.72453850905515 || err != nil {
		t.Errorf("SqrtPI Function Failed ... ")
	}

	x, err = SqrtPI(0)

	if x != 0 || err != nil {
		t.Errorf("SqrtPI Function Failed ... ")
	}
}

func TestDegree(t *testing.T) {
	x, err := Degrees(1)
	if x != 57.29577951308238 || err != nil {
		t.Errorf("Degrees Function Failed ... ")
	}

	x, err = Degrees(-2.5)
	if x != -143.23944878270595 || err != nil {
		t.Errorf("Degrees Function Failed ... ")
	}
	x, err = Degrees(PI())
	if x != 180 || err != nil {
		t.Errorf("Degrees Function Failed ... ")
	}

	x, err = Degrees(2 * PI())
	if x != 360 || err != nil {
		t.Errorf("Degrees Function Failed ... ")
	}
}
