/*
------------------------------------
ixakblt - ibrahim AKBULUT
------------------------------------
Web Site :ixakblt
------------------------------------
All Sites : @ixakblt
------------------------------------
*/
package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
) //kütüphaneleri dahil ettik

func main() {
	//Sonsuz döngüye soktuk programı ki her seferinde yeniden başlatmayalım for un while yerine kullanımına örnek
	for true {
		//degişkenler oluşturuldu
		var islem int
		var deger1 int
		var deger2 int
		fmt.Println("\n-------------------------")
		fmt.Println("---------ixHesap---------")
		fmt.Println("-------------------------")
		fmt.Print("1. Degeri giriniz : ")
		fmt.Scan(&deger1) //fmt kütüphanesi içinde scan fonksiyonu ile girdiyi aldık & ile basit bir pointer kullanımı ile alıyoruz veriyi C dilinde scanf  gibi bir girdi mevcut
		fmt.Print("2. degeri girin : ")
		fmt.Scan(&deger2)
		temizle() // main fonksiyonu dışında tanımladığım fonksiyonu çağırdım ekranı temizlemeye yarıyor
		fmt.Println("islem seciniz \n\n1)toplama\n2)cikarma\n3)carpma\n4)bolme\n")
		fmt.Print("islem : ")
		fmt.Scan(&islem)
		if islem > 4 || islem < 0 {
			// basit bir if kullanımı ile işlem 4 den büyük veya 0 dan küçük ise parantez içini gerçekleştir demiş olduk
			fmt.Print("gecersiz islem")
			time.Sleep(3 * time.Second) //3 saniye bekleme ekledik
			temizle()
		} else if islem == 1 {
			//else if kullanımı
			fmt.Printf("islem sonucunuz : %d ", topla(deger1, deger2)) //topla fonksiyonununun return değerini %d ile format velirleyici kullandık
		} else if islem == 2 {
			fmt.Printf("islem sonucu : %d ", eksilt(deger1, deger2))
		} else if islem == 3 {
			fmt.Printf("islem sonucu : %d : ", carp(deger1, deger2))
		} else if islem == 4 {
			fmt.Printf("islem sonucu : %d  ", bol(deger1, deger2))
		}
	}

}

func topla(x int, y int) int {
	//topla fonksiyonunu yazdık int değer return ediyor diğer fonksiyonlawr da aynı şekilde
	return x + y

}

func eksilt(x int, y int) int {
	return x - y
}

func carp(x int, y int) int {
	return x * y
}

func bol(x int, y int) int {
	return x / y
}

func temizle() {
	//temizleme fonksiyonu windos için cls komutu vermekte python da ki os.system ile aynı görev
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
