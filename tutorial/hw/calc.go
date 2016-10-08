package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type calculator func(int, int) int

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(": ")
	text, _ := reader.ReadString('\n')

	// do parse content
	op := checkOperator(text)

	reg := regexp.MustCompile("[\\+\\-\\*\\/\\%(\\r\\n)]")
	splitted := reg.Split(text, 3)

	a, _ := strconv.Atoi(splitted[0])
	b, _ := strconv.Atoi(splitted[1])

	fmt.Println(calc(op, a, b))
}

func calc(method calculator, a int, b int) int {
	return method(a, b)
}

func checkOperator(str string) calculator {
	switch {
	case strings.Index(str, "+") != -1:
		return func(a int, b int) int { return a + b }
	case strings.Index(str, "-") != -1:
		return func(a int, b int) int { return a - b }
	case strings.Index(str, "*") != -1:
		return func(a int, b int) int { return a * b }
	case strings.Index(str, "/") != -1:
		return func(a int, b int) int { return a / b }
	case strings.Index(str, "%") != -1:
		return func(a int, b int) int { return a % b }
	}
	return func(a int, b int) int { return a + b }
}
