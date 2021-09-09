package bai3

/*
Bài 3 Viết function remove những phần tử bị trùng nhau trong mảng
Ví dụ: removeDuplicates([1,2,5,2,6,2,5]) => [1,2,5,6]
*/

/*
remove duplicate element from list
*/
func RemoveDuplicates1(i []int) ([]int, error) {
	nonDuplicateList := []int{}
	for _, v := range i {
		if !contains(nonDuplicateList, v) {
			nonDuplicateList = append(nonDuplicateList, v)
		}
	}
	return nonDuplicateList, nil
}

/*
check element exist in array
*/
func contains(i []int, value int) bool {
	for _, v := range i {
		if v == value {
			return true
		}
	}

	return false
}

/*
remove duplicate element from list
*/
func RemoveDuplicates2(i []int) ([]int, error) {
	keepUniqueElement := map[int]bool{}
	nonDuplicateList := []int{}
	for _, v := range i {
		keepUniqueElement[v] = true
	}

	for key, _ := range keepUniqueElement {
		nonDuplicateList = append(nonDuplicateList, key)
	}

	return nonDuplicateList, nil
}
