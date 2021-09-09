/*
Bài 1 Viết function tìm ra số lớn thứ nhì trong mảng các số.
Ví dụ: max2Numbers([2, 1, 3, 4]) => 3
*/
package bai1

import (
	"errors"
	"sort"
)

/*
get the second largest number in array
ex:
1, 3, 1, 3 -> 3
1, 3, 2, 4 -> 3
*/
func SecondLargestNumber1(i []int) (int, error) {
	if len(i) < 2 {
		return 0, errors.New("array length is less than 2")
	}

	sort.Ints(i)
	return i[len(i)-2], nil
}

/*
get the second largest number in array
ex:
1, 3, 1, 3 -> 3
1, 3, 2, 4 -> 3
*/

func SecondLargestNumber2(i []int) (int, error) {
	if len(i) < 2 {
		return 0, errors.New("array length is less than 2")
	}

	_, maxIndex, _ := findMaxNumber(i)
	max2nd := i[0]

	for index, value := range i {
		if max2nd < value && index != maxIndex {
			max2nd = value
		}
	}

	return max2nd, nil
}

/*
find max number and index of max number
*/
func findMaxNumber(i []int) (int, int, error) {
	if len(i) < 1 {
		return 0, 0, errors.New("array length is less than 1")
	}

	maxValue, maxIndex := i[0], 0
	for index, value := range i {
		if maxValue < value {
			maxValue = value
			maxIndex = index
		}
	}

	return maxValue, maxIndex, nil
}
