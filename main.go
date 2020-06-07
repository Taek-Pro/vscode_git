package main

import "fmt"

func main() {
	iResult := sumFunc(1, 2, 3, 4) //:= 은 자료형을 선언하지 않아도 된다.

	fmt.Println("Result : ")
	fmt.Println(iResult)
}

func sumFunc(iValue ...int) int {
	var iResult int = 0
	iResult = 0

	for _, i := range iValue {
		iResult += i
	}

	return iResult
}

func Result(iValue ...int) string {

}
