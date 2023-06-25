package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//the same

	for i := 0; i < 10; {
		fmt.Println(i)
		i++
	}

	j := 10
	//while
	for j > 0 {
		fmt.Println(j)
		j--
	}

	//never reached
	// sum := 0
	// for {
	// 	sum++
	// }
	// fmt.Println(sum);

	//continue
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
	//break
	for i := 0; true; i++ {
		if i == 5 {
			fmt.Printf("%d is five\n", i)
			break
		}
		fmt.Println(i)
	}
}
