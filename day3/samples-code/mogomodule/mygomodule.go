package mygomodule

func Reverse(input string) (result string) {
	for _, v := range input {
		result = string(v) + result
	}
	return
}
