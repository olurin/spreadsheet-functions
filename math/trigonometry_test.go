package mathlib

import (
	"log"
	"testing"
)

func TestTrigonometry(t *testing.T) {
	v, err := Acot(0)
	log.Println(v, err)

	v, err = Degree(v)
	log.Println(v, err)

	v, err = AcotH(2)
	log.Println(v, err)

}
