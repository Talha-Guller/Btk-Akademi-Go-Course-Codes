package loops

import "fmt"

//1 kullanıcadan bir sayi girmesini ve bu sayının asal olup olmadığını kontrol et
func Demo4() {
	sayi := 1

	fmt.Print("Bir sayı Giriniz : ")
	fmt.Scanln(&sayi)

	asalMi := true
	for i := 2; i < sayi; i++ {
		if sayi%i == 0 {
			asalMi = false
		}

	}

	if asalMi == true {
		fmt.Printf("%v Asaldır", sayi)
	} else {
		fmt.Printf("%v Asal değildir", sayi)
	}

}
