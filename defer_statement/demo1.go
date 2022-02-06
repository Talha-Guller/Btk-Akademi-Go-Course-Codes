package defer_statement

import "fmt"

func A() {
	fmt.Println("A fonk Çalıştı")
}

func C() {
	fmt.Println("C fonk Çalıştı")
}

func D() {
	fmt.Println("D fonk Çalıştı")
}

func B() {
	defer A()
	defer C()
	defer D()
	fmt.Println("B fonk Çalıştı")
}
