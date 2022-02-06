package conditionals

import "fmt"

func Demo2() {
	var sayi1 int = 20
	var sayi2 int = 10
	var sayi3 int = 15

	if sayi1 > sayi2 {
		if sayi1 > sayi3 {
			fmt.Println("en büyük sayi1 : " + fmt.Sprintf("%v", sayi1))
		} else {
			fmt.Println("en büyük sayi3 : " + fmt.Sprintf("%v", sayi3))
		}
	} else if sayi2 > sayi3 {
		fmt.Println("en büyük sayi2 : " + fmt.Sprintf("%v", sayi2))
	} else {
		fmt.Println("en büyük sayi3 : " + fmt.Sprintf("%v", sayi3))

	}
}
