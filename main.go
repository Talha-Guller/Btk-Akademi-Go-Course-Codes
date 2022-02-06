package main

import (
	"fmt"
	"golesson/erorr_handling"
)

func main() {
	//variables.Demo1()
	//conditionals.Demo2()

	//loops.Demo5()

	//arrays.Demo4()
	//sileces.Demo2()
	//functions.SelamVer("Talha")
	//var sonuc = functions.Topla(5, 6)
	//fmt.Println(sonuc)

	// toplam, çıkarma, çarpma, bölme := functions.Dortİslem(10, 6)

	// fmt.Println(toplam)
	// fmt.Println(çıkarma)
	// fmt.Println(çarpma)
	// fmt.Println(bölme)

	// fmt.Println(functions.ToplaVariadic(1, 2, 3, 4, 5, 6, 7, 8, 9))
	// fmt.Println(functions.ToplaVariadic(1, 4))
	// fmt.Println(functions.ToplaVariadic())

	// sayilar := []int{5, 6, 9, 7, 8}
	// fmt.Println(functions.ToplaVariadic(sayilar...))

	//maps.Demo1()

	//for_range.Demo3()

	// sayi := 20
	// pointers.Demo1(&sayi)
	// fmt.Println("MAindeki sayi :", sayi)

	// sayilar := []int{1, 2, 3, 40}
	// pointers.Demo2(sayilar)
	// fmt.Println("MAindeki sayi :", sayilar[0])

	//structs.Demo2()

	// go goroutines.CiftSayilar()
	// go goroutines.TekSayilar()
	// time.Sleep(5 * time.Second) //time.Second -> 1 saniye demektir
	// fmt.Println("Main Bitti")

	// ciftSayiCn := make(chan int)
	// tekSayiCn := make(chan int)
	// go channels.CiftSayilar(ciftSayiCn)
	// go channels.TekSayilar(tekSayiCn)

	// ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn
	// carpım := ciftSayiToplam * tekSayiToplam

	// fmt.Println("Carpım :", carpım)

	//interfaces.Demo2()

	//defer_statement.Demo3()

	//erorr_handling.Demo1()

	//interfaces.Demo3()

	fmt.Println(erorr_handling.TahminEt2(102))
}
