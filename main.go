package main

import "fmt"

func main() {
	fmt.Println("Добро пожаловать в кофе машину")
	var amount int
	fmt.Scan(&amount)
	drinkPricec := []int{100, 200, 300}
	drinkNames := []string{"Кофе", "Лате", "Капучино"}
	//fmt.Scan(&drinkPricec, drinkNames)

	for i := 0; i < 3; i++ {
		if amount >= drinkPricec[i] { // если  тебе скажет mistapes []int
			fmt.Println("Вам доступно", drinkNames[i]) // если укажешь [i]то он тебе по отдельности будет показывать
		}

	}
	var smt string
	fmt.Scan(&smt)

}


