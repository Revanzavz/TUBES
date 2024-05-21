package main

import (
	// "bufio"
	"fmt"

	"time"

	// "github.com/inancgumus/screen"
)

const NMAX = 1000

type dataBuku struct {
	judul, pengarang                        string
	nomorISBN, jumlahEksemplar, tahunTerbit int
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
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
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




func tambahkanEksemplar(buku *tabBuku, i int) {
	fmt.Println("Masukkan judul buku yang ingin ditambahkan eksemplarnya: ")
	var keyword string
	fmt.Scan(&keyword)
	for j := 0; j < i; j++ {
		if buku[j].judul == keyword {
			fmt.Println("Buku ditemukan!")
			fmt.Println("Judul buku: ", buku[j].judul)
			fmt.Println("Pengarang: ", buku[j].pengarang)
			fmt.Println("Nomor ISBN: ", buku[j].nomorISBN)
			fmt.Println("Jumlah eksemplar: ", buku[j].jumlahEksemplar)
			fmt.Println("Tahun terbit: ", buku[j].tahunTerbit)
			fmt.Println(" ")

			var tambah int
			fmt.Println("Masukkan jumlah eksemplar yang ingin ditambahkan: ")
			fmt.Scan(&tambah)
			buku[j].jumlahEksemplar += tambah
			fmt.Println("Jumlah eksemplar berhasil ditambahkan")
			fmt.Println("Jumlah eksemplar saat ini: ", buku[j].jumlahEksemplar)
			fmt.Println(" ")
		} else {
			fmt.Println("Buku tidak ditemukan")
			fmt.Println(" ")
		}
	}
}








func kembalikanBuku(buku *tabBuku, i int, riwayat *tabPeminjaman, jmlP int) {
	var keyword string
	var hari, bulan, tahun int
	var found bool
	fmt.Print("Masukkan judul buku yang ingin dikembalikan: ")
	fmt.Scan(&keyword)
	for j := 0; j < i; j++ {
		if (*buku)[j].judul == keyword {
			found = true
			buku[j].jumlahEksemplar++
			fmt.Println("Buku ditemukan!")
			fmt.Println("Judul buku: ", (*buku)[j].judul)
			fmt.Println("Pengarang: ", (*buku)[j].pengarang)
			fmt.Println("Nomor ISBN: ", (*buku)[j].nomorISBN)
			fmt.Println("Jumlah eksemplar: ", (*buku)[j].jumlahEksemplar)
			fmt.Println("Tahun terbit: ", (*buku)[j].tahunTerbit)
			fmt.Println(" ")

			// Input tanggal pengembalian
			fmt.Print("Masukkan tanggal pengembalian (dd mm yyyy): ")
			fmt.Scan(&hari, &bulan, &tahun)
			for hari > 31 || bulan > 12 {
				fmt.Println("Tanggal tidak valid")
				fmt.Print("Masukkan tanggal pengembalian (dd mm yyyy): ")
				fmt.Scan(&hari, &bulan, &tahun)
				for hari > 31 || bulan > 12 {
					fmt.Println("Tanggal tidak valid")
					fmt.Print("Masukkan tanggal pengembalian (dd mm yyyy): ")
					fmt.Scan(&hari, &bulan, &tahun)
				}
			}

			tanggalPengembalian := time.Date(tahun, time.Month(bulan), hari, 0, 0, 0, 0, time.UTC)

			// Cari riwayat peminjaman
			for k := 0; k < jmlP; k++ {
				if (*riwayat)[k].judul == keyword {
					tanggalPeminjaman := (*riwayat)[k].tanggalPeminjaman
					durasi := tanggalPengembalian.Sub(tanggalPeminjaman).Hours() / 24
					fmt.Println("Durasi Peminjaman: ", durasi, "hari")
					if durasi >= 7 {
						fmt.Println("Anda terlambat mengembalikan buku selama", durasi, "hari")
						fmt.Println("Denda yang harus dibayar sebesar Rp", (durasi-7)*1000)
					}
				}
			}
		}
	}
	if !found {
		fmt.Println("Buku tidak ditemukan")
		fmt.Println(" ")
	}
}

func hapusBuku(buku *tabBuku, i *int, index int) {
	var keyword string
	var found bool
	fmt.Print("Masukkan judul buku yang ingin dihapus: ")
	fmt.Scan(&keyword)
	for j := 0; j < *i; j++ {
		if (*buku)[j].judul == keyword {
			index = j
			found = true
		}
	}
	if found {
		for k := index; k < *i; k++ {
			(*buku)[k] = (*buku)[k+1]
		}
		*i--
		fmt.Println("Buku berhasil dihapus")
		fmt.Println(" ")
	} else {
		fmt.Println("Buku tidak ditemukan")
		fmt.Println(" ")
	}
}




