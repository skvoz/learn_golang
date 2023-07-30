package main

import (
	"fmt"
	"strconv"
)

func cube(p float64) float64 {
	return p * p * p
}

func f1(n uint) (int, int) {
	fact := 1
	sum := 0
	for i := 1; i <= int(n); i++ {
		sum += i
		fact *= i
	}
	return fact, sum
}

func myFunc(p string) int {
	p1, err := strconv.Atoi(p)
	p2, err := strconv.Atoi(p + p)
	p3, err := strconv.Atoi(p + p + p)

	if err != nil {
		fmt.Print(err)
	}

	return p1 + p2 + p3
}

func varsum(s ...int) int {
	r := 0
	for _, v := range s {
		r += v
	}
	return r
}
func nakedReturnFunc(a, b int) (x, y int) {
	x = a + 1
	y = b + 1

	return
}

func searchItem(a []string, b string) bool {
	//strings.EqualFold
	//case sensetive insensetive
	for _, v := range a {
		if v == b {
			return true
		}
	}
	return false
}

func incFunc(a int) int {
	return a
}

func main() {
	fmt.Println("Ex1#############################")
	r := cube(4)

	fmt.Printf("cube of 4 %#v\n", r)
	fmt.Println("Ex2#############################")

	r1, r2 := f1(5)
	fmt.Println("fact 5:", r1, "sum 5:", r2)

	fmt.Println("Ex3#############################")
	fmt.Println("fact 5:", myFunc("5"))

	fmt.Println("Ex4#############################")
	fmt.Println("var sum 1 2 3 is :", varsum(1, 2, 3))

	fmt.Println("Ex5#############################")
	fmt.Println(nakedReturnFunc(1, 2))

	fmt.Println("Ex6#############################")
	fmt.Println("return value must be true:", searchItem(
		[]string{"foo", "bar", "baz"},
		"foo",
	))
	fmt.Println("return value must be false:", searchItem(
		[]string{"foo", "bar", "baz"},
		"foo1",
	))

	fmt.Println("Ex7#############################")
	a := incFunc
	fmt.Printf("type func var get func as value: %T\n", a)

	fmt.Printf("get value from var with func as value %#v\n", a(3))

	/**
	Ex1#############################
	cube of 4 64
	Ex2#############################
	fact 5: 120 sum 5: 15
	Ex3#############################
	fact 5: 615
	Ex4#############################
	var sum 1 2 3 is : 6
	Ex5#############################
	2 3
	Ex6#############################
	return value must be true: true
	return value must be false: false
	Ex7#############################
	type func var get func as value: func(int) int
	get value from var with func as value 3
	**/
}
