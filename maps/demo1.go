package maps

import (
	"fmt"
)

func Demo1() {
	sozluk := make(map[string]string)
	sozluk["book"] = "kitap"
	sozluk["pencil"] = "Kalem"
	sozluk["table"] = "masa"

	fmt.Println(sozluk["book"])
	fmt.Println("Elaman Sayısı :", len(sozluk))
	fmt.Println("Sözlük :", sozluk)

	delete(sozluk, "book")
	fmt.Println("Sözlük :", sozluk)
	fmt.Println("Elaman Sayısı :", len(sozluk))

	deger, varMi := sozluk["book"]
	fmt.Println(deger)
	fmt.Println("Listede Olma durumu var mı :", varMi)

	sozluk2 := map[string]string{"glass": "bardak", "microphone": "mikrofon"}
	fmt.Println("Sözlük :", sozluk2)

}
