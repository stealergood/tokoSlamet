package menu

import (
	"bufio"
	"fmt"
	"os"
	"toko-slamet/features/CRUD/create"
	"toko-slamet/features/CRUD/delete"
	"toko-slamet/features/CRUD/update"
)

func menuUpdate() {
	var pilihan, stock int
	fmt.Println("=== Menu Update Stock ===")
	for i, item := range create.Model {
		fmt.Printf("%d. %-16s stock = %-2d\n", i+1, item.Nama, item.Stock)
	}
	fmt.Println("=========================")
	fmt.Print("Pilih menu barang(0 untuk kembali): ")
	fmt.Scan(&pilihan)
	if pilihan == 0 {
		fmt.Printf("\x1bc")
		MainMenu()
	} else if pilihan > 0 && pilihan <= len(create.Model) {
		fmt.Print("Jumlah Penambahan Stock: ")
		fmt.Scan(&stock)
		fmt.Printf("\x1bc")
		update.Update(pilihan-1, stock)
		menuUpdate()
	} else {
		fmt.Printf("\x1bc")
		fmt.Println("Pilihan Tidak Valid")
		menuUpdate()
	}
}

func menuDelete() {
	var pilihan int
	fmt.Println("=== Menu Delete ===")
	for i, item := range create.Model {
		fmt.Printf("%d. %-16s\n", i+1, item.Nama)
	}
	fmt.Println("===================")
	fmt.Print("Pilih angka barang (0 untuk batal) >>> ")
	fmt.Scan(&pilihan)
	fmt.Printf("\x1bc")
	if pilihan == 0 {
		MainMenu()
	} else if pilihan > 0 && pilihan <= len(create.Model){
		delete.Delete(pilihan-1)
		menuDelete()
	} else {
		fmt.Println("Pilihan Tidak Valid")
		menuDelete()
	}
}

func menuDisplay() {
	var back int
	fmt.Println("=================================")
	fmt.Println("         DATABASE BARANG         ")
	fmt.Println("=================================")
	if len(create.Model) == 0 {
		fmt.Println("Data Belum Ditambahkan")
	} else {
		for i, item := range create.Model {
			fmt.Printf("%d. %-16s stock = %-2d\n", i+1, item.Nama, item.Stock)
		}
	}
	fmt.Println("=================================")
	fmt.Print("Ketik 0 Untuk Kembali >>> ")
	fmt.Scan(&back)
	fmt.Printf("\x1bc")
	if back == 0{
		MainMenu()
	} else {
		fmt.Println("Pilihan Tidak Valid")
		menuDisplay()
	}
}

func menuAdd() {
	var nama string
	var stock int
	fmt.Println("=== Menu Add Barang ===")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	fmt.Print("Masukkan Nama Barang: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		nama = scanner.Text()
	}
	fmt.Print("Masukkan Stock Barang: ")
	fmt.Scan(&stock)
	fmt.Printf("\x1bc")
	create.Add(nama, stock)
	MainMenu()
}

func MainMenu() {
	var menu int
	fmt.Println("======== Menu Utama ========")
	fmt.Println("1. Tambahkan Barang")
	fmt.Println("2. Update Stock Barang")
	fmt.Println("3. Hapus Barang")
	fmt.Println("4. Tampilkan Stock Barang")
	fmt.Println("0. Keluar")
	fmt.Println("============================")
	fmt.Print("Pilihan >>> ")
	fmt.Scan(&menu)

	switch menu {
	case 1:
		fmt.Printf("\x1bc")
		menuAdd()
	case 2:
		fmt.Printf("\x1bc")
		menuUpdate()
	case 3:
		fmt.Printf("\x1bc")
		menuDelete()
	case 4:
		fmt.Printf("\x1bc")
		menuDisplay()
	case 0:
		fmt.Printf("\x1bc")
		fmt.Println("Keluar...")
	default:
		fmt.Printf("\x1bc")
		fmt.Println("Keluar...")
	}
}