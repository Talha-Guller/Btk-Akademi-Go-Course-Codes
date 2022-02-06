package functions

import "fmt"

func Topla(sayi1, sayi2 int) int {
	sonuc := sayi1 * sayi2
	return sonuc
}

func SelamVer(kullanıcıAdı string) {
	fmt.Println("Merhaba", kullanıcıAdı)
}
