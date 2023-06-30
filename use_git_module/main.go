package main

//import package, not module
import (
	"fmt"

	"github.com/skvoz/go_test_mod/geometry"
)

func main()  {
	fmt.Println(geometry.CubeVolume(5))
}