package main

import (
	"fmt"
	"math/big"
)

func main() {

	//подключаем библиотеку big для работы с особыми числами и создаем две исходных переменных > 2^20
	a := big.NewInt(8388608)
	b := big.NewInt(1230123)

	//сложение
	add := big.NewInt(0)
	add = add.Add(a, b)
	fmt.Println(add)

	//вычитание
	sub := big.NewInt(0)
	sub = sub.Sub(a, b)
	fmt.Println(sub)

	//умножение
	mul := big.NewInt(0)
	mul = mul.Mul(a, b)
	fmt.Println(mul)

	//деление
	div := big.NewInt(0)
	div = div.Div(a, b)
	fmt.Println(div)

}
