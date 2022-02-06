package defer_statement

import "fmt"

type product struct {
	productName string
	unitPrace   int
}

func (p product) save() {
	fmt.Println("Ürün Kaydedildi :", p.productName)
	defer Log()
}

func Log() {
	fmt.Println("Loglandı")
}

func Demo3() {
	p := product{productName: "Laptop", unitPrace: 5000}
	defer p.save()
	p = product{productName: "Mause", unitPrace: 45}

	fmt.Println("İşlem Başarılı")
	fmt.Println(p.productName)
}
