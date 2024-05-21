package main

import (
	// "bufio"
	"fmt"

	"time"

	// "github.com/inancgumus/screen"
)

const NMAX = 1000

type dataBuku struct {
	judul, pengarang, nomorISBN                       string
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
	fmt.Print("\033[2J")
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
		// fmt.Print("\033[2J")
		// fmt.Print("\033[H")
		clearscreen()
		switch choice {
		case 1:
			MenuAdmin()
			clearscreen()
		case 2:
			MenuUser()
			fmt.Print("\033[2J")
			clearscreen()
		case 3:
			return
		default:
			fmt.Println("Menu tidak tersedia, silahkax` menu yang tersedia")
		}
	}
}

func gambarMenu() {

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
		fmt.Print("\033[2J")
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
				cetakBuku(buku, i)
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


















