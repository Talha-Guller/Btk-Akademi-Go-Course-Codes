package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Println("Paranız Yeterli değil")
	} else if cekilmekIstenen == hesap {
		fmt.Println("Paranız hazırlanıyor")
		fmt.Println("paranız kalmadı")
	} else {
		fmt.Println("Paranız hazırlanıyor")
	}
}
