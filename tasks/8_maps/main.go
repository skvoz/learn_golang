package main

import "fmt"

func main() {
	fmt.Println("Ex1#############################")
	var m1 map[float64]bool
	m2 := map[int]string{1: "foo", 2: "bar"}
	m2[10] = "Abba"
	fmt.Printf("m1 is %#v\nm2 is %#v\n", m1, m2)
	fmt.Printf("m2[1] is %#v\n", m2[1])
	fmt.Printf("unexists key m2[14] is %#v\n", m2[14])

	fmt.Println("Ex2#############################")
	m := map[int]bool{1: true, 2: false, 3: false}
	fmt.Printf("output m %#v\n", m)
	delete(m, 1)
	fmt.Printf("output m after delete 1 el %#v\n", m)
	for i, v := range m {
		fmt.Printf("index=%d, value=%v\n", i, v)
	}
	/**
	  	Ex1#############################
	  m1 is map[float64]bool(nil)
	  m2 is map[int]string{1:"foo", 2:"bar", 10:"Abba"}
	  m2[1] is "foo"
	  unexists key m2[14] is ""
	  Ex2#############################
	  output m map[int]bool{1:true, 2:false, 3:false}
	  output m after delete 1 el map[int]bool{2:false, 3:false}
	  index=2, value=false
	  index=3, value=false
	  **/
}
