package logger

import (
	"day4-sample-codes-demointerface/basic"
	"fmt"
)

type Logger interface {
	Log(err error)
}

type OneMountLogger struct{}
type TechMasterLogger struct{}

func (logger OneMountLogger) Log(err error) {
	switch err.(type) {
	case basic.OneMountError:
		fmt.Println("Send to One Mount QA Team")
		fmt.Println(err)
	case basic.NetworkError:
		fmt.Println("Send to Devops Team")
		fmt.Println(err)
	default:
		fmt.Println(err)
	}
}

func (logger TechMasterLogger) Log(err error) {
	fmt.Println(err)
}
