package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()

	fmt.Println(strings.Join(os.Args[1:], " "))

	elapsedTime := time.Since(startTime)
	fmt.Printf("Join 사용 : %d\n\n", elapsedTime)

	// -----------------------------------------------

	startTime = time.Now()

	fmt.Println(os.Args[1:])

	elapsedTime = time.Since(startTime)
	fmt.Printf("Join 미사용 : %d\n", elapsedTime)

	// Join을 사용해서 배열을 출력한 코드 : 21125, 미사용한 코드 : 217500
	// Join을 사용한 코드가 10배 이상 빠르다.
	// Join 함수는 Buffer의 크기를 늘리고 이후 문자열을 이어붙이는 방법을 사용해서 속도가 매우 빠르다.
}
