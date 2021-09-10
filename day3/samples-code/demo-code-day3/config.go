package main

import (
	"fmt"
	"time"
)

type Config struct {
	Host string
	Port int
}

var config = Config{
	Host: "localhost",
	Port: 8080,
}
var pConfig = new(Config)

func init() {
	pConfig.Host = "127.0.0.1"
	pConfig.Port = 9000
}

func useConfigNormalStyle() {
	fmt.Println("Config #1")
	fmt.Println(config.Host, config.Port)

	fmt.Println("Config #2")
	fmt.Println(pConfig.Host, pConfig.Port)

}

func useConfigGoStyle() {
	go func() {
		fmt.Println("Config #1")
		fmt.Println(config.Host, config.Port)
	}()
	go func() {
		fmt.Println("Config #2")
		fmt.Println(pConfig.Host, pConfig.Port)
	}()
	time.Sleep(1 * time.Second)
}
