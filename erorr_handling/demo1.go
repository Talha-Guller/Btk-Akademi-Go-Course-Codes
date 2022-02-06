package erorr_handling

import (
	"fmt"
	"os"
)

// type assertion --> tip doğrulama
func Demo1() {
	f, err := os.Open("demo11.txt")
	// eğer dosya mevcutsa err=nil döner hata varsa nil den farklı bir şey döner - f de okunan değeri alır
	if err != nil { //yani hata ar demek
		if pErr, ok := err.(*os.PathError); ok { // ilk yapılan işlem atama işlemi pErr hatayı o değişkene atıyor - ok ise tru mu false mu değerini döndürüyor
			fmt.Println("Dosya Bulunamadı :", pErr.Path)
			return
		} else {
			fmt.Println("Dosya Bulunamadı")
			return
		}
	} else {
		fmt.Println(f.Name())
	}
}
