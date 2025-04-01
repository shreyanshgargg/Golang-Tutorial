package main

import "fmt"

func main() {
	fmt.Println("loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	//		for d:=0; d<len(days); d++ {
	//			fmt.Println(days[d])
	//		}
	//	}
	// for i := range days {
	// 	fmt.Println(i, days[i])

	// }
	// for index, days := range days {
	// 	fmt.Printf("index is %d and day is %s\n", index, days)
	// }

	rougevalue := 1

	for rougevalue < 10 {
		if rougevalue == 2 {
			goto lco
		}
		if rougevalue == 5 {
			rougevalue++
			continue
		}
		fmt.Println(rougevalue)
		rougevalue++
	}

lco:
	fmt.Println("jumping at lco")

}
