package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"Bİlgisayar", 500, "XYZ"})
	fmt.Println(product{"Bİlgisayar", 500, "abc"})
	fmt.Println(product{name: "Bİlgisayar"})
}

type product struct {
	name      string
	unitPrice float64
	brand     string
}
