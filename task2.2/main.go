package main

import (
	"fmt"

	TimeGetter "github.com/KoljanTestwamore/go-time-getter"
)

func main() {
	time, err := TimeGetter.GetTime()

	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Printf("Current time %v\n", time)
}