package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, args := range os.Args[1:] { // value pair {index, element}
		s += sep + args
		sep = " "
	}

	// 1. blank identifier "_" : golang 에서는 프로그램에서 쓰이지 않는 변수를 에러 처리한다.
	// 이를 피하기 위해서는 "_"를 변수명으로 사용해야 한다.

	// 2. 빈 문자열 변수를 생성하는 몇 가지 방법
	// s := ""
	// var s string
	// var s = ""
	// var s string = ""

	fmt.Println(s)
}
