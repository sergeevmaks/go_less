package main

import (
	"fmt"
	"math/big"
)

func main() {
	numbers := make([]*big.Int, 0, 100)
	num := new(big.Int)
	numbers = append(numbers, num.SetInt64(1))
	numbers = append(numbers, num.SetInt64(1))

	for n := 2; n < 100; n++ {
		fibonacci(&numbers, n)
	}
	fmt.Println(numbers)

}

func fibonacci(slice *[]*big.Int, i int) {
	var n = new(big.Int)
	fmt.Print((*slice))
	a := (*slice)[i-1]
	b := (*slice)[i-2]
	n.Add(a, b)
	*slice = append(*slice, n)

}

func division(num, div int) bool {
	if num%div == 0 {
		return true
	}
	return false
}
