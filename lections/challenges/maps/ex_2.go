package main

func main() {
	var m1 map[int]bool
	m1[5] = true

	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}

	// tried this but not :)
	// s1 := fmt.Printf("%#v", m2)
	// s2 := fmt.Printf("%#v", m3)
	// // fmt.Println(m2 == m3)
	// fmt.Println(s1 == s2)

	//correct solution
	_, _, _ = m1, m2, m3
}
