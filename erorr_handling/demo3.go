package erorr_handling

import "fmt"

type borderException struct {
	parametre int
	massage   string
}

func (b *borderException) Error() string {
	return fmt.Sprintf("%d---%s", b.parametre, b.massage)
}

func TahminEt2(tahmin int) (string, error) {
	if tahmin < 0 || tahmin > 100 { // "||"-> yada demek
		return "", &borderException{parametre: tahmin, massage: "Sınırların dışında kaldı"}

	}
	return "Bildiniz", nil
}
