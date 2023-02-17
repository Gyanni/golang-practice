package main

import (
	"fmt"
	"os"
	"strings"
)

// Input : hi hello good
func main() {
	fmt.Println(strings.Join(os.Args[1:], " ")) // output : hi hello good
	fmt.Println(os.Args[1:])                    // output : [hi hello good]
	fmt.Println(os.Args[0:])                    // Exercise 1.1
	// output : [/private/var/folders/y3/gzcc6kg97n7grfc26115w6dh0000gn/T/GoLand/___go_build_echo3_go hi hello good]
	// os.Args[0]에는 build 명령어가 매개변수로 입력된다.
}
