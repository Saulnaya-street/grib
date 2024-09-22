package zadachka2

//2. Поиск дубликатов
//Напиши программу, которая принимает слайс строк и выводит все строки, которые встречаются более одного раза.

import "fmt"

func Dublicat(slova []string) {
	slovaKol := make(map[string]int)

	for _, w := range slova {
		slovaKol[w]++

	}
	fmt.Println("Дубликат:")

	for w, count := range slovaKol {
		if count > 1 {
			fmt.Println(w)
		}
	}
}

func main() {
	z := []string{"саня", "саня", "саня", "ваня", "СТАС", "ДЭН", "СТАС"}
	Dublicat(z)
}
