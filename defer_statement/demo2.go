package defer_statement

import "fmt"

func Demo2(sayi int) string {
	DeferFunc()
	if sayi%2 == 0 {
		return "Çift sayidir"
	}

	if sayi%2 != 0 {
		return "Tek sayidir"
	}
	return "Belli değil"
}

func Test() {
	sonuc := Demo2(10)
	fmt.Println(sonuc)
}

func DeferFunc() {
	fmt.Println("Bitti")
}
