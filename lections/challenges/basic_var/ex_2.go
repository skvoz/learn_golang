package main

func main() {
	//ex #1
	var (
		a int
		b float64
		c bool
		d string
	)

	x, y, z := -20, -15.5, "Gopher!"

	// fmt.Println(a, b, c, d, x, y, z)
	_, _, _, _, _, _, _ = a, b, c, d, x, y, z
}
