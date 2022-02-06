package interfaces

import "fmt"

func dogrula(i interface{}) {
	sayi, ok := i.(int) // burda tip dönüşümü yap yapa biliyorsan tru dön yapamıyorsa false dödürüyor "ok" -- sayi da değere at
	fmt.Println(sayi, ok)
}

func Demo3() {
	var sayi1 interface{} = "Engin"
	dogrula(sayi1)

	var sayi2 interface{} = 5
	dogrula(sayi2)
}
