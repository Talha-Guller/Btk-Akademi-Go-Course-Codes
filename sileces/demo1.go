package sileces

import "fmt"

func Demo1() {
	isimler := make([]string, 3)

	fmt.Println(isimler)

	isimler[0] = "engin"
	isimler[1] = "derin"
	isimler[2] = "zeynep"
	isimler = append(isimler, "talha")

	fmt.Println(isimler)
	fmt.Println(len(isimler))

}
