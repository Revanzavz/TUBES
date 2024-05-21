package main

import (
	"fmt"
)

func sortingBuku(buku *tabBuku, i int) {
	var choiceSort int
	for choiceSort != 3 {
		fmt.Println("Menu Sorting:")
		fmt.Println("1. Sorting berdasarkan judul")
		fmt.Println("2. Sorting berdasarkan tahun terbit")
		fmt.Println("3. Exit")
		fmt.Print("Pilih menu : ")
		fmt.Scan(&choiceSort)
		fmt.Print("\033[2J")
		clearscreen()
		switch choiceSort {
		case 1:
			for j := 0; j < i; j++ {
				for k := j + 1; k < i; k++ {
					if (*buku)[j].judul > (*buku)[k].judul {
						temp := (*buku)[j]
						(*buku)[j] = (*buku)[k]
						(*buku)[k] = temp
					}
				}
			}
		case 2:
			for j := 0; j < i; j++ {
				for k := j + 1; k < i; k++ {
					if (*buku)[j].tahunTerbit > (*buku)[k].tahunTerbit {
						temp := (*buku)[j]
						(*buku)[j] = (*buku)[k]
						(*buku)[k] = temp
					}
				}
			}
		case 3:
			return
		default:
			fmt.Println("Menu tidak tersedia, silahkan pilih menu yang tersedia")
		}
	
	}
}