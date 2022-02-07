package string_functions

import (
	"fmt"
	s "strings" // böyle yazılması strings yerine artık s yazacam demektir. Bunlara (alias) denir
)

func Demo1() {
	isim := "Talha"
	fmt.Println(s.Count(isim, "a"))
	//isim'in içinde ne Kadar a varsa onu döndürüyor. int dödürür.(case sensitive)Küçük-Büyük harf duyarlı

	fmt.Println(s.Contains(isim, "A"))
	//isim'in içinde A var mı yok mu onu dödürür.Yani true-false

	fmt.Println(s.Index(isim, "k"))
	//Aranan kelimenin string içade de kaçıncı sırada olduğunu döndürür.Ama ilk göründüğü yerden başka bir yerde tekrardan var ise onu değil en baştakini döndürür.Eğer bulamazsa -1 döndürür.

	fmt.Println(s.ToLower(isim))
	//metni Küçük harfe çevir

	fmt.Println(s.ToUpper(isim))
	//metni Büyük harfe çevir

}
