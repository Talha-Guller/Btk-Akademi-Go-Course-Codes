package erorr_handling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {
	aklımdakiSayi := 50
	if tahmin < 0 || tahmin > 100 { // "||"-> yada demek
		return "", errors.New("1-100 arasında bir sayı giriniz.")
	}
	if tahmin > aklımdakiSayi {
		return "Daha Küçük bir Sayi Giriniz", nil
	}
	if tahmin < aklımdakiSayi {
		return "Daha Büyük bir Sayi Giriniz", nil
	}
	return "Bildiniz", nil
}

func Demo2() {
	mesaj, _ := TahminEt(49)
	fmt.Println(mesaj) //bunda sadece mesajı dödürüyor
	//fmt.Println(TahminEt(50)) // bunda ise hem mesajı hem de hatayı döndürüyor
	//fmt.Println(TahminEt(101))
	//fmt.Println(TahminEt(-10))
}
