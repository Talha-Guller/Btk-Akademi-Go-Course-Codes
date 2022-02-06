package arrays

import "fmt"

func Demo3() {
	sayilar := [5]int{20, 30, 50, 10, 2}
	fmt.Println(sayilar)

	enBuyuk := sayilar[0]
	for i := 0; i < 5; i++ {
		if sayilar[i] > enBuyuk {
			enBuyuk = sayilar[i]
		}
	}
	fmt.Println("En büyük : ", enBuyuk)
}
