package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("welcome to a class on slices")

	var fruitlist = []string{"apple", "banana", "grape", "orange", "kiwi"}

	fmt.Println("type of fruitlist is %T\n", fruitlist)

	fruitlist = append(fruitlist, "mango", "peach", "plum")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:3])
	fmt.Println(fruitlist)

	highscores := make([]int, 4)

	highscores[0] = 234
	highscores[1] = 456
	highscores[2] = 789
	highscores[3] = 123

	highscores = append(highscores, 555, 666, 777)

	fmt.Println(highscores)

	sort.Ints(highscores)
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

}
