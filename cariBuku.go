package main

import (
	"fmt"
	"strings"
)

func cariBuku(buku tabBuku, i int) {
	var caridatabuku int
	var keyword string
	var keywordint int
	var found bool

	fmt.Println("1. Judul buku")
	fmt.Println("2. Pengarang")
	fmt.Println("3. Nomor ISBN")
	fmt.Println("4. Tahun terbit")
	fmt.Println("5. Exit")
	fmt.Print("Masukkan data buku yang ingin dicari: ")

	fmt.Scan(&caridatabuku)
	switch caridatabuku {
	case 1:
			//sequential search
		fmt.Print("Jika judul buku lebih dari 1 kata gunakan _ sebagai pengganti spasi\n")
		fmt.Print("Masukkan judul buku yang ingin dicari: ")
		fmt.Scan(&keyword)
		clearscreen()
		for j := 0; j < i; j++ {
			if buku[j].judul == keyword {
				printBuku(buku, j)
				found = true
			}
		}
		if !found {
			fmt.Println("Buku tidak ditemukan")
			fmt.Println(" ")
		}
	case 2:
			//sequential search
		fmt.Print("Jika nama pengarang lebih dari 1 kata gunakan _ sebagai pengganti spasi\n")
		fmt.Print("Masukkan pengarang buku yang ingin dicari: ")
		fmt.Scan(&keyword)
		clearscreen()
		for j := 0; j < i; j++ {
			if buku[j].pengarang == keyword {
				printBuku(buku, j)
				found = true
			}
		}
		if !found {
			fmt.Println("Buku tidak ditemukan")
			fmt.Println(" ")
		}
	case 3:
			//sequential search
		fmt.Print("Masukkan nomor ISBN buku yang ingin dicari: ")
		fmt.Scan(&keywordint)
		clearscreen()
		for j := 0; j < i; j++ {
			if buku[j].nomorISBN == keyword {
				printBuku(buku, j)
				found = true
			}
		}
		if !found {
			fmt.Println("Buku tidak ditemukan")
			fmt.Println(" ")
		}
	case 4:
			//sequential search
		fmt.Print("Masukkan tahun terbit buku yang ingin dicari: ")
		fmt.Scan(&keywordint)
		clearscreen()
		for j := 0; j < i; j++ {
			if buku[j].tahunTerbit == keywordint {
				printBuku(buku, j)
				found = true
			}
		}
		if !found {
			fmt.Println("Buku tidak ditemukan")
			fmt.Println(" ")
		}
	case 5:
		return
	default:
		clearscreen()
		fmt.Println("Menu tidak tersedia, silahkan pilih menu yang tersedia")
	}
}

func printBuku(buku tabBuku, j int) {
	maxNoUrut := len("No Urut")
	maxNamaBuku := len("Judul Buku")
	maxNamaPengarang := len("Pengarang")
	maxNomorISBN := len("Nomor ISBN")
	maxJumlahEksemplar := len("Jumlah Eksemplar")
	maxTahunTerbit := len("Tahun Terbit")

	widthNoUrut := len(fmt.Sprintf("%d", j+1))
	widthNamaBuku := len(strings.ReplaceAll(buku[j].judul, "_", " "))
	widthNamaPengarang := len(strings.ReplaceAll(buku[j].pengarang, "_", " "))
	widthNomorISBN := len(fmt.Sprintf(buku[j].nomorISBN))
	widthJumlahEksemplar := len(fmt.Sprintf("%d", buku[j].jumlahEksemplar))
	widthTahunTerbit := len(fmt.Sprintf("%d", buku[j].tahunTerbit))

	if widthNoUrut > maxNoUrut {
		maxNoUrut = widthNoUrut
	}
	if widthNamaBuku > maxNamaBuku {
		maxNamaBuku = widthNamaBuku
	}
	if widthNamaPengarang > maxNamaPengarang {
		maxNamaPengarang = widthNamaPengarang
	}
	if widthNomorISBN > maxNomorISBN {
		maxNomorISBN = widthNomorISBN
	}
	if widthJumlahEksemplar > maxJumlahEksemplar {
		maxJumlahEksemplar = widthJumlahEksemplar
	}
	if widthTahunTerbit > maxTahunTerbit {
		maxTahunTerbit = widthTahunTerbit
	}

	// Cetak header
	fmt.Println(strings.Repeat("━", maxNoUrut+maxNamaBuku+maxNamaPengarang+maxNomorISBN+maxJumlahEksemplar+maxTahunTerbit+19))
	fmt.Printf("| %-*s | %-*s | %-*s | %-*s | %-*s | %-*s |\n",
		maxNoUrut, "No Urut",
		maxNamaBuku, "Judul Buku",
		maxNamaPengarang, "Pengarang",
		maxNomorISBN, "Nomor ISBN",
		maxJumlahEksemplar, "Jumlah Eksemplar",
		maxTahunTerbit, "Tahun Terbit")
	fmt.Println(strings.Repeat("━", maxNoUrut+maxNamaBuku+maxNamaPengarang+maxNomorISBN+maxJumlahEksemplar+maxTahunTerbit+19))

	// Cetak data buku
	fmt.Printf("| %-*d | %-*s | %-*s | %-*s | %-*d | %-*d |\n",
		maxNoUrut, j+1,
		maxNamaBuku, strings.ReplaceAll(buku[j].judul, "_", " "),
		maxNamaPengarang, strings.ReplaceAll(buku[j].pengarang, "_", " "),
		maxNomorISBN, buku[j].nomorISBN,
		maxJumlahEksemplar, buku[j].jumlahEksemplar,
		maxTahunTerbit, buku[j].tahunTerbit)
	fmt.Println(strings.Repeat("─", maxNoUrut+maxNamaBuku+maxNamaPengarang+maxNomorISBN+maxJumlahEksemplar+maxTahunTerbit+19))
	fmt.Println(" ")
}
