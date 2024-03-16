package create

import (
	"fmt"
	"toko-slamet/features/models"
)

var Model []models.Barang

func Add(nama string, stock int) {
	stocking := models.Barang{
		Nama: nama,
		Stock: stock,
	}
	Model = append(Model, stocking)
	fmt.Println("Barang Berhasil Ditambahkan")
}