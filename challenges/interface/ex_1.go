package main

import "fmt"

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	return c.licenceNo
}

func (c car) Name() string {
	return c.brand
}

func main() {
	var v vehicle = car{licenceNo: "111", brand: "222"}
	fmt.Println(v.License(), v.Name())
}
