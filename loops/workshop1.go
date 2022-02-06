package loops

import "fmt"

func Demo2() {
	var bulunmasıGereknSayi, tahminEdilenSayi, sayac int = 80, 0, 1
	var deger string = ""

	fmt.Print("Bir sayı Giriniz :")
	fmt.Scanln(&tahminEdilenSayi)

	for bulunmasıGereknSayi != tahminEdilenSayi {

		if bulunmasıGereknSayi > tahminEdilenSayi {
			fmt.Println("Daha büyük bir sayı giriniz ")
			fmt.Scanln(&tahminEdilenSayi)
			sayac = sayac + 1
		} else if bulunmasıGereknSayi < tahminEdilenSayi {
			fmt.Println("Daha küçük bir sayı giriniz ")
			fmt.Scanln(&tahminEdilenSayi)
			sayac = sayac + 1
		}

		if sayac > 0 || sayac <= 3 {
			deger = "Süper"
		} else if sayac >= 4 && sayac <= 6 {
			deger = "güzel"
		} else if sayac > 6 {
			deger = "fena değil"
		}

	}
	fmt.Printf("brova bildiniz. %v Tahmin : %v", sayac, deger)

}
