package main

import (
	"fmt"
	"strings"
)

func f1(a ...int) {
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)
}

func f2(a ...int) {
	a[0] = 50
}

func SumAndProduct(a ...float64) (float64, float64) {
	sum := 0.
	product := 1.
	for _, v := range a {
		sum += v
		product *= v
	}

	return sum, product
}

func personInformation(age int, names ...string) string {
	fullName := strings.Join(names, " ")
	returnString := fmt.Sprintf("Age: %d, Full Name : %s", age, fullName)

	return returnString
}
func main() {
	f1(1, 2, 3, 4)
	//[]int
	//[]int{1, 2, 3, 4}
	nums := []int{1, 2}
	nums = append(nums, 3, 4, 5)
	f1(nums...)
	//[]int
	//[]int{1, 2, 3, 4, 5}

	f2(nums...)
	fmt.Println("Nums:", nums) //Nums: [50 2 3 4 5]

	s, p := SumAndProduct(2.0, 5., 10.)
	fmt.Println(s, p) //17 100

	info := personInformation(10, "Petr", "Korbel")
	fmt.Println(info) //17 100

}
