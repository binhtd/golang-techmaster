package main

import (
	"demo-module/utils"
	"fmt"

	"github.com/binhtd/mygomodule/mathutil"
	mygomoduleV2 "github.com/binhtd/mygomodule/v2"
	"rsc.io/quote"
)

func main() {
	fmt.Println(utils.IsPrime(7))
	fmt.Println(utils.ProperCase("welcome to the our world!"))

	fmt.Println(mathutil.Add(10, 11))
	fmt.Println(mathutil.Minus(10, 11))
	fmt.Println(mygomoduleV2.Lower("Toi yeu Viet Nam"))
	fmt.Println(quote.Go())
}
