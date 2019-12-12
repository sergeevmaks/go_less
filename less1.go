package main

import (
	"fmt"
	"math"
)

func main() {
	convertor()
	triangle()
}

func triangle() {
	var (
		cat1 float64
		cat2 float64
	)
	fmt.Println("Введите длиннц катета1")
	fmt.Scanln(&cat1)
	fmt.Println("Введите длиннц катета2")
	fmt.Scanln(&cat2)
	area := (cat1 * cat2) / 2
	hypotenuse := math.Sqrt(math.Pow(cat1, 2) + math.Pow(cat2, 2))
	perimeter := cat1 + cat2 + hypotenuse
	fmt.Printf("площадь треугольника = %.2f \n", area)
	fmt.Printf("гипотенуза треугольника = %.2f \n", hypotenuse)
	fmt.Printf("периметр треугольника = %.2f \n", perimeter)
}

func convertor() {
	const kursDollar float64 = 66.7
	var rub int32
	fmt.Println("Введите сумму в рублях?")
	fmt.Scanln(&rub)
	dollar := float64(rub) / kursDollar
	fmt.Printf("Вы получите = %.2f доллара\n", dollar)
}
