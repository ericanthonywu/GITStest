package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	run(7)
}

func run(input int) {
	var result = make([]string, input)
	for i := 0; i < input; i++ {
		var previousInt = 1
		if i != 0 {
			previousInt, _ = strconv.Atoi(result[i-1])
		}

		result[i] = strconv.Itoa(i + previousInt)
	}
	fmt.Println(strings.Join(result, "-"))
}
