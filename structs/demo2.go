package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

//struct bir metotdur
func (c customer) save() {
	fmt.Println("Eklendi :", c.firstName)
}

func Demo2() {
	c := customer{firstName: "Talha", lastName: "Güller", age: 20}
	c.save()
}
