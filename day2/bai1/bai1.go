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

	max2ndIndex := len(i) - 2
	for {
		if max2ndIndex < 0 {
			return 0, errors.New("not found max2d")
		}

		if i[max2ndIndex] == i[max2ndIndex+1] && max2ndIndex >= 0 {
			max2ndIndex--
		} else {
			return i[max2ndIndex], nil
		}
	}
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

	maxValue, _, _ := findMaxNumber(i)
	max2nd := i[0]

	for _, value := range i {
		if max2nd < value && value != maxValue {
			max2nd = value
		}
	}

	if max2nd == maxValue {
		return 0, errors.New("not found max2d")
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
