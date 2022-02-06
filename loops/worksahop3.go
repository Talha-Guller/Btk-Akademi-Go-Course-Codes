package loops

import "fmt"

func Demo5() {
	sayi1 := 17296
	sayi2 := 18416

	sonuc1 := 0
	for i := 1; i < sayi1; i++ {
		if sayi1%i == 0 {
			sonuc1 = sonuc1 + i
		}
	}

	sonuc2 := 0
	for i := 1; i < sayi2; i++ {
		if sayi2%i == 0 {
			sonuc2 = sonuc2 + i
		}
	}
	if sonuc1 == sayi2 {
		if sonuc2 == sayi1 {
			fmt.Printf("%v - %v Arkadaş Sayıdır", sayi1, sayi2)
		}
	} else {
		fmt.Printf("%v - %v Arkadaş Sayı değildir", sayi1, sayi2)
	}
}
