package bai2

/*
Bài 2 Cho 1 mảng các chuỗi. Viết function lọc ra các phần tử có độ dài lớn nhất.
Ví dụ: findMaxLengthElement["aba", "aa", "ad", "c", "vcd"] => ["aba", "vcd"]
*/

/*
find list element has length that is equal max length
*/
func FindListElementHasMaxLength1(s []string) ([]string, error) {
	elements := map[int][]string{}

	for index := range s {
		elements[len(s[index])] = append(elements[len(s[index])], s[index])
	}

	keyHasMaxLength := 0
	listElementHasMaxLength := []string{}
	for key, value := range elements {
		if keyHasMaxLength < key {
			keyHasMaxLength = key
			listElementHasMaxLength = value
		}
	}
	return listElementHasMaxLength, nil
}

/*
find list element has length that is equal max length
*/
func FindListElementHasMaxLength2(s []string) ([]string, error) {
	maxLength := 0
	listElementHasMaxLength := []string{}
	for i := range s {
		if maxLength < len(s[i]) {
			maxLength = len(s[i])
		}
	}

	for i := range s {
		if maxLength == len(s[i]) {
			listElementHasMaxLength = append(listElementHasMaxLength, s[i])
		}
	}

	return listElementHasMaxLength, nil
}
