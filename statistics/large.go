package statistics

import (
	"errors"
	"sort"
)

//Nums implements sort.Interface for numbers passed into Large function
type Nums []float64

func (a Nums) Len() int           { return len(a) }
func (a Nums) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Nums) Less(i, j int) bool { return a[i] > a[j] }

//Large returns the Kth LARGEST value from a list of supplied numbers, for a given value K
func Large(index int, numbers ...float64) (float64, error) {

	if index < 1 || index > len(numbers) {
		return 0.0, errors.New("#NUM!, Occurred because supplied index is less than 1 or greater than the number of values in the supplied array")
	}
	sort.Sort(Nums(numbers))

	return numbers[index-1], nil
}
