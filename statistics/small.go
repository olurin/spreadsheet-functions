package statistics

import (
	"errors"
	"sort"
)

//SmallNums implements sort.Interface for numbers passed into Small function
type SmallNums []float64

func (a SmallNums) Len() int           { return len(a) }
func (a SmallNums) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SmallNums) Less(i, j int) bool { return a[i] < a[j] }

//Small returns the Kth SMALLEST value from a list of supplied numbers, for a given value K
func Small(index int, numbers ...float64) (float64, error) {

	if index < 1 || index > len(numbers) {
		return 0.0, errors.New("#NUM!, Occurred because supplied index is less than 1 or greater than the number of values in the supplied array")
	}
	sort.Sort(SmallNums(numbers))

	return numbers[index-1], nil
}
