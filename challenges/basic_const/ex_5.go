package main

func main() {
	const a int = 7
	const b float64 = 5.6
	const c = float64(a) * b //more simple way do a untyped cost

	x := 8
	_ = x
	// const xc int = x

	// const noIPv6 = math.Pow(2, 128)
	//error ->
	//const rules
	//1. can't change
	//2. can't initiate at runtime (math.pow - runtime function , len - not runtime function )
	//3. can't initiate with variables
}
