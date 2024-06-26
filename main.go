package main

import (
	"fmt"
	"time"
)

const NMAX = 1000

type dataBuku struct {
	judul, pengarang, nomorISBN  string
	jumlahEksemplar, tahunTerbit int
}

type tabBuku [NMAX]dataBuku

var buku tabBuku
var i, index int

type dataPeminjaman struct {
	judul, namaPeminjam string
	tanggalPeminjaman   time.Time
}

type tabPeminjaman [NMAX]dataPeminjaman

var riwayatPeminjaman tabPeminjaman
var jumlahPeminjaman int

func main() {
	loading()
	menu()
}

func clearscreen() {
	// "\033[2J" adalah string yang berisi urutan escape ANSI. Urutan ini digunakan untuk membersihkan layar terminal. Penjelasan lebih lanjut:
	// 1) \033 adalah escape character (ESC) dalam notasi oktal (ASCII code 27).
	// 2) [2J adalah bagian dari urutan escape yang memberitahu terminal untuk menghapus layar seluruhnya.
	fmt.Print("\033[2J")
	// "\033[H" adalah urutan escape ANSI lain yang digunakan untuk memindahkan kursor ke posisi kiri atas layar (baris 1, kolom 1).
	// 1) \033 adalah escape character (ESC) yang sama seperti sebelumnya.
	// 2) [H memberitahu terminal untuk memindahkan kursor ke posisi home (kiri atas).
	fmt.Print("\033[H")
}

func menu() {
	var choice int
	clearscreen()
	gambarMenu()
	fmt.Println("Welcome to Perpustakaan!")
	for choice != 3 {
		fmt.Println("Anda seorang:")
		fmt.Println("1. Admin")
		fmt.Println("2. User")
		fmt.Println("3. Exit")
		fmt.Print("Pilih menu : ")
		fmt.Scan(&choice)
		clearscreen()
		switch choice {
		case 1:
			MenuAdmin()
			clearscreen()
		case 2:
			MenuUser()
			clearscreen()
		case 3:
			return
		default:
			fmt.Println("Menu tidak tersedia, silahkan menu yang tersedia")
		}
	}
}

func gambarMenu() {
	bookAscii := `
             .--.           .---.        .-.
         .---|--|   .-.     | A |  .---. |~|    .--.
      .--|===|Ch|---|_|--.__| S |--|:::| |~|-==-|==|---.
      |%%|NT2|oc|===| |~~|%%| C |--|   |_|~|CATS|  |___|-.
      |  |   |ah|===| |==|  | I |  |:::|=| |    |GB|---|=|
      |  |   |ol|   |_|__|  | I |__|   | | |    |  |___| |
      |~~|===|--|===|~|~~|%%|~~~|--|:::|=|~|----|==|---|=|
      ^--^---'--^---^-^--^--^---'--^---^-^-^-==-^--^---^-'
`
	textAscii := `
	__                __                
	(_  _| _  _  _ |_ |  \ _ |_ _  _  _  
	__)(-|(_||||(_||_ |__/(_||_(_|| )(_) 
					 _/  
	 __    __                            
	|  \. |__)_ _ _     _|_ _ |  _  _  _ 
	|__/| |  (-| |_)|_|_)|_(_||((_|(_|| )
	             |                       
	`
	fmt.Println(bookAscii)
	fmt.Println(textAscii)
}

func MenuAdmin() {
	var choiceAdmin int
	fmt.Println("Welcome Admin!")
	for choiceAdmin != 8 {
		fmt.Println("Menu Admin:")
		fmt.Println("1. Input data buku")
		fmt.Println("2. Cetak data buku")
		fmt.Println("3. Sorting data buku")
		fmt.Println("4. Cari data buku")
		fmt.Println("5. Tambahkan eksemplar buku")
		fmt.Println("6. Edit data buku")
		fmt.Println("7. Hapus data buku")
		fmt.Println("8. Exit")
		fmt.Print("Pilih menu : ")
		fmt.Scan(&choiceAdmin)
		clearscreen()
		switch choiceAdmin {
		case 1:
			inputBuku(&buku, &i)
		case 2:
			if i == 0 {
				fmt.Println("Data buku masih kosong")
				fmt.Println("Input Data Buku terlebih dahulu dengan memilih menu 1")
				fmt.Println(" ")
			} else {
				cetakBuku(buku, i)
				fmt.Println(" ")
			}
		case 3:
			if i == 0 {
				fmt.Println("Data buku masih kosong")
				fmt.Println("Input Data Buku terlebih dahulu dengan memilih menu 1")
				fmt.Println(" ")
			} else {
				sortingBuku(&buku, i)
			}
		case 4:
			if i == 0 {
				fmt.Println("Data buku masih kosong")
				fmt.Println("Input Data Buku terlebih dahulu dengan memilih menu 1")
				fmt.Println(" ")
			} else {
				cariBuku(buku, i)
			}
		case 5:
			if i == 0 {
				fmt.Println("Data buku masih kosong")
				fmt.Println("Input Data Buku terlebih dahulu dengan memilih menu 1")
				fmt.Println(" ")
			} else {
				tambahkanEksemplar(&buku, i)
			}
		case 6:
			editBuku(&buku, i, index)
		case 7:
			hapusBuku(&buku, &i, index)
		case 8:
			return
		default:
			fmt.Println("Menu tidak tersedia, silahkan pilih menu yang tersedia")
		}
	}
}
