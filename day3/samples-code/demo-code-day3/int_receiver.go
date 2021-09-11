package main

import (
	"fmt"
)

type MyInt int

func (p *MyInt) Double() {
	*p = *p * 2
}

func DemoIntReceiver() {
	muoi := new(MyInt)
	*muoi = 10
	muoi.Double()
	fmt.Println(*muoi)
}
