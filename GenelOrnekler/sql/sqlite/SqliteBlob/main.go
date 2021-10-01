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
	"database/sql"
	"io/ioutil"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	rdresim, _ := ioutil.ReadFile("a.png") //okuduk
	db, _ := MakeDb("Kedim.db")            //db nesnesi üretik
	defer db.Close()                       //iş bitince kapan knk
	CreateTable(db)                        // tablo oluşturduk
	InsertBlob(rdresim, db)                //tabloya resim dosyası ekle
	Toimg(db, 1, "kbr.png")                // resim dosyasını resime geri çeviri
}

//MakeDB DB oluştrduk
func MakeDb(isim string) (*sql.DB, error) {
	return sql.Open("sqlite3", isim)
}

//CreateTable tablo oluşturduk
func CreateTable(db *sql.DB) {
	db.Exec("CREATE TABLE IF NOT EXISTS `Res`(`id`INTEGER PRIMARY KEY AUTOINCREMENT,`veri`BLOB)")

}

//InsertBlob gelen resim dosyasını ekledik
func InsertBlob(res []byte, db *sql.DB) {
	a, _ := db.Prepare("INSERT INTO Res(veri) VALUES(?)")
	a.Exec(res)
}

//Toimg verdiğimiz iddeki veriyi veridğğimiz isimde dönüştürdü
func Toimg(db *sql.DB, id int, filename string) {
	var veri []byte
	row := db.QueryRow("SELECT `veri` FROM `Res` WHERE id=$1", id)
	row.Scan(&veri)
	img, _ := os.Create(filename)
	img.Write(veri)
}
