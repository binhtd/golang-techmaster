package main

import (
	"fmt"
)

func demoEnglishTranslation() {
	dict := map[string]string{
		"hello":   "xin chao",
		"goodbye": "tam biet",
	}

	for key, value := range dict {
		fmt.Println(key, " ", value)
	}
}

func demoMapInterface() {
	configs := map[string]interface{}{
		"host":     "v1.int.dantri.vn",
		"tls":      true,
		"user":     "user1",
		"password": "pass1",
		"port":     8080,
	}

	for key, value := range configs {
		fmt.Println(key, " ", value)
	}
}
