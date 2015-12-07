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

func TestDegreeS(t *testing.T) {
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

func TestRadians(t *testing.T) {
	x, err := Radians(50)

	if x != 0.8726646259971639 || err != nil {
		t.Errorf("Radians Function Failed ... ")
	}

	x, err = Radians(-180)

	if x != -3.14159265358979 || err != nil {
		t.Errorf("Radians Function Failed ... ")
	}
	x, err = Radians(180)

	if x != 3.14159265358979 || err != nil {
		t.Errorf("Radians Function Failed ... ")
	}

	x, err = Radians(360)

	if x != 6.28318530717958 || err != nil {
		t.Errorf("Radians Function Failed ... ")
	}
}

func TestCos(t *testing.T) {
	x, err := Cos(0.785398163)
	if x != 0.707106781467586 || err != nil {
		t.Errorf("Cos Function Failed ... ")
	}

	x, err = Cos(0)
	if x != 1 || err != nil {
		t.Errorf("Cos Function Failed ... ")
	}

	x, err = Cos(-PI() / 3)
	if x != 0.5000000000000009 || err != nil {
		t.Errorf("Cos Function Failed ... ")
	}

	y, _ := Radians(-30)

	x, err = Cos(y)
	if x != 0.8660254037844389 || err != nil {
		t.Errorf("Cos Function Failed ... ")
	}
}

func TestAcos(t *testing.T) {
	x, err := Acos(-1)

	if x != 3.141592653589793 || err != nil {
		t.Errorf("Acos Function Failed ... ")
	}

	x, err = Acos(0)
	if x != 1.5707963267948966 || err != nil {
		t.Errorf("Acos Function Failed ... ")
	}

	sqrt, _ := Sqrt(2)

	x, err = Acos(1 / sqrt)
	if x != 0.7853981633974484 || err != nil {
		t.Errorf("Acos Function Failed ... ")
	}

	y, _ := Degrees(x)

	if Round(y, 0) != 45 {
		t.Errorf("Acos Function Failed ... ")
	}
}

func TestCosh(t *testing.T) {
	x, err := CosH(0)
	if x != 1 || err != nil {
		t.Errorf("CosH Function Failed ... ")
	}

	x, err = CosH(0.5)
	if Round(x, 9) != 1.127625965 || err != nil {
		t.Errorf("CosH Function Failed ... ")
	}

	x, err = CosH(-2)
	if Round(x, 9) != 3.762195691 || err != nil {
		t.Errorf("CosH Function Failed ... ")
	}
}

func TestAcosH(t *testing.T) {
	x, err := AcosH(1)
	if x != 0 || err != nil {
		t.Errorf("AcosH Function Failed ... ")
	}

	x, err = AcosH(2.5)
	if Round(x, 9) != 1.566799237 || err != nil {
		t.Errorf("AcosH Function Failed ... ")
	}

	x, err = AcosH(5)
	if Round(x, 9) != 2.29243167 || err != nil {
		t.Errorf("AcosH Function Failed ... ")
	}
}

func TestSec(t *testing.T) {
	x, err := Sec(-3.14159265358979)
	if x != -1 || err != nil {
		t.Errorf("Sec Function Failed ... ")
	}

	x, err = Sec(0)
	if x != 1 || err != nil {
		t.Errorf("Sec Function Failed ... ")
	}

	x, err = Sec(PI() / 4)
	if Round(x, 9) != 1.414213562 || err != nil {
		t.Errorf("Sec Function Failed ... ")
	}
}

func TestSecH(t *testing.T) {
	x, err := SecH(-3.14159265358979)

	if x != 0.0862667383340547 || err != nil {
		t.Errorf("SecH Function Failed ... ")
	}

	x, err = SecH(0)
	if x != 1 || err != nil {
		t.Errorf("SecH Function Failed ... ")
	}

	x, err = SecH(PI() / 4)

	if x != 0.7549397087141317 || err != nil {
		t.Errorf("SecH Function Failed ... ")
	}
}

func TestSin(t *testing.T) {
	x, err := Sin(0.785398163)
	if x != 0.7071067809055092 || err != nil {
		t.Errorf("Sin Function Failed ... ")
	}

	x, err = Sin(PI() / 6)
	if Round(x, 1) != 0.5 || err != nil {
		t.Errorf("Sin Function Failed ... ")
	}

	x, err = Sin(-PI() / 3)
	if x != -0.866025403784438 || err != nil {
		t.Errorf("Sin Function Failed ... ")
	}
}

func TestAsin(t *testing.T) {
	x, err := Asin(-1)

	if x != -1.5707963267948966 || err != nil {
		t.Errorf("Asin Function Failed ... ")
	}

	x, err = Asin(0)
	if x != 0 || err != nil {
		t.Errorf("Asin Function Failed ... ")
	}

	sqrt, _ := Sqrt(2)

	x, err = Asin(1 / sqrt)
	if x != 0.7853981633974482 || err != nil {
		t.Errorf("Asin Function Failed ... ")
	}

	y, _ := Degrees(x)

	if Round(y, 0) != 45 {
		t.Errorf("Asin Function Failed ... ")
	}
}

func TestSinH(t *testing.T) {
	x, err := SinH(0)
	if x != 0 || err != nil {
		t.Errorf("SinH Function Failed ... ")
	}

	x, err = SinH(0.5)
	if Round(x, 9) != 0.521095305 || err != nil {
		t.Errorf("SinH Function Failed ... ")
	}

	x, err = SinH(-2)
	if Round(x, 8) != -3.62686041 || err != nil {
		t.Errorf("SinH Function Failed ... ")
	}
}

func TestAsinH(t *testing.T) {
	x, err := AsinH(0)
	if x != 0 || err != nil {
		t.Errorf("AsinH Function Failed ... ")
	}

	x, err = AsinH(-0.5)
	if Round(x, 8) != -0.48121183 || err != nil {
		t.Errorf("AsinH Function Failed ... ")
	}

	x, err = AsinH(2)
	if Round(x, 9) != 1.443635475 || err != nil {
		t.Errorf("AsinH Function Failed ... ")
	}
}

func TestCsc(t *testing.T) {
	x, err := Csc(-6)
	if Round(x, 9) != 3.578899547 || err != nil {
		t.Errorf("Csc Function Failed ... ")
	}

	x, err = Csc(1.5707963267949)
	if x != 1 || err != nil {
		t.Errorf("Csc Function Failed ... ")
	}

	x, err = Csc(PI() / 4)
	if Round(x, 9) != 1.414213562 || err != nil {
		t.Errorf("Csc Function Failed ... ")
	}
}

func TestCscH(t *testing.T) {
	x, err := CscH(-3.14159265358979)

	if Round(x, 9) != -0.086589538 || err != nil {
		t.Errorf("CscH Function Failed ... ")
	}

	x, err = CscH(PI() / 3)
	if Round(x, 9) != 0.800405293 || err != nil {
		t.Errorf("CscH Function Failed ... ")
	}

	x, err = CscH(PI() / 4)

	if Round(x, 9) != 1.151183871 || err != nil {
		t.Errorf("CscH Function Failed ... ")
	}
}
