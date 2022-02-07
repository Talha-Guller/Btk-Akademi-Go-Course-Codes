package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "Talha"
	fmt.Println(s.HasPrefix(isim, "Ta"))
	//isim'in Başı Ta diye mi başlıyor onu soruyor.true-false döner.Gerçek hayatta ismin öününde bir ön ek içerip içermediğine bakılır

	fmt.Println(s.HasSuffix(isim, "en"))
	//isim'in sonu en ile bitiyor mu onu soruyor.true-false döner.

	fmt.Println(s.Index(isim, "lh"))
	//başladığı yerin indexini döner.

	harfler := []string{"t", "a", "l", "h", "a"}
	sonuc := s.Join(harfler, "*")
	//string dizileri bir araya getiriyor.İkinci eklediğimiz değer de aralarına ne koyulacaksa onu veriliyor. ve bu artık yeni bi strin oluyor.
	fmt.Println(sonuc)

	fmt.Println(s.Replace(sonuc, "*", "+", -1))
	//sonuc string'inde * lar yerine + koy demek oluyor.-1 yerine 2 filan verilseydi gördüğün 2 tanesini değiştir demekti.
	//-1 ise gördüğün hepsini değiştir demek.Bu gerçek hayatta banklardaki mesela ibanlardaki aralarda - işareti koymak istemiyorlar o zaman bu methodu kullanıyorlar.

	fmt.Println(s.Split(sonuc, ""))
	//sonuc stringini *'a göre ayırmaya yarıyor. ve ayırdıkları her biri dizinin elemanlarına dönüyor.
	//eğer koyulan ayraç bulunmazsa yine tek bir array oluyor

	fmt.Println(s.Repeat(sonuc, 5))
	//sonuc stringini 5 defa yan yana yazdır demek.buda yeni bi string oluyor.

}
