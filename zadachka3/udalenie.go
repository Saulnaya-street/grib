package zadachka3

import "fmt"

//3. Удаление элемента из слайса
//Создай функцию, которая удаляет элемент по индексу из слайса целых чисел
//и возвращает обновленный слайс. Не забудь учесть случаи, когда индекс выходит за границы слайса.

func Udalenie(slice []int, d int) []int {
	if d < 0 || d >= len(slice) {
		fmt.Println("вышел за пределы")
		return slice
	}
	slice[d] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	z := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(z)
	newZ := Udalenie(z, 3)
	fmt.Println(newZ)
}
