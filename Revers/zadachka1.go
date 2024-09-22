package Revers

//1. Реверс слайса
//Напиши функцию, которая принимает слайс целых чисел
//и возвращает новый слайс с элементами в обратном порядке.

import "fmt"

func Revers(a []int) {
	x := []int{}
	for i := len(a) - 1; i >= 0; i-- {
		x = append(x, a[i])
	}
	fmt.Println(x)
}

func main() {
	slice := []int{1, 2, 3, 4}
	Revers(slice)
}
