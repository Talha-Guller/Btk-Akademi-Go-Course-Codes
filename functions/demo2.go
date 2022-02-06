package functions

func Dortİslem(sayi1, sayi2 int) (int, int, int, float32) {
	toplam := sayi1 + sayi2
	çıkarım := sayi1 - sayi2
	çarpma := sayi1 * sayi2
	bölme := float32(sayi1 / sayi2)
	return toplam, çıkarım, çarpma, float32(bölme)
}
