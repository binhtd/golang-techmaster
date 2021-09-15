package logger

import (
	"day4-sample-codes-demointerface/basic"
)

func DemoLogger() {
	var logger Logger
	logger = OneMountLogger{}

	err := basic.OneMountError{
		Msg:      "Failed to register account",
		Division: "VinID",
		Module:   "ONline Shopping",
		Func:     "Account_Register",
		Line:     120,
	}

	logger.Log(err)
}
