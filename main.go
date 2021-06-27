/**
main.go
Brief:  This is the main program flow, it reverses an integer array.

@author Mauricio Portugués Castellanos.
@version    1.0, 2021-06-27.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Pow(base int, pow int) int {
	var index int = 1
	var result int = base

	for ; index < pow; index++ {
		result *= base
	}

	return result
}

func read_input(array *[]int, input_array *[]string) bool {
	var array_size int = 0
	var raw_input string
	var reader = bufio.NewReader(os.Stdin)

	fmt.Scanln(&array_size)

	if array_size < 1 || array_size > Pow(10, 3) {
		return false
	}

	raw_input, _ = reader.ReadString('\n')
	raw_input = strings.Trim(raw_input, "\n")

	*array = make([]int, array_size)
	*input_array = strings.Split(raw_input, " ")

	return true
}

func reverse_array(input []string, array *[]int) {
	var array_size int = len(*array) - 1

	for index := 0; index <= array_size; index++ {
		(*array)[index], _ = strconv.Atoi(input[Abs(index-array_size)])
	}
}

func main() {
	var array []int
	var input_array []string

	if read_input(&array, &input_array) {
		fmt.Printf("\n%v\n", input_array)
		reverse_array(input_array, &array)
		fmt.Printf("\n%v\n", array)
	}
}
