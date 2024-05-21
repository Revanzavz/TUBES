package main

	import (
	"fmt"
	"strings"
)

func cetakBuku(buku tabBuku, i int) {
	fmt.Println("Data Buku:")
	maxNoUrut := len("No Urut")
	maxNamaBuku := len("Judul Buku")
	maxNamaPengarang := len("Pengarang")
	maxNomorISBN := len("Nomor ISBN")
	maxJumlahEksemplar := len("Jumlah Eksemplar")
	maxTahunTerbit := len("Tahun Terbit")

	// Perhitungan lebar kolom berdasarkan data buku
	for j := 0; j < i; j++ {
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
	for j := 0; j < i; j++ {
		fmt.Printf("| %-*d | %-*s | %-*s | %-*s | %-*d | %-*d |\n",
			maxNoUrut, j+1,
			maxNamaBuku, strings.ReplaceAll(buku[j].judul, "_", " "),
			maxNamaPengarang, strings.ReplaceAll(buku[j].pengarang, "_", " "),
			maxNomorISBN, buku[j].nomorISBN,
			maxJumlahEksemplar, buku[j].jumlahEksemplar,
			maxTahunTerbit, buku[j].tahunTerbit)
		fmt.Println(strings.Repeat("─", maxNoUrut+maxNamaBuku+maxNamaPengarang+maxNomorISBN+maxJumlahEksemplar+maxTahunTerbit+19))
	}
	fmt.Println(" ")
}