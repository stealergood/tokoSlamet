package delete

import (
	"fmt"
	"toko-slamet/features/CRUD/create"
)

func Delete(index int) {
	create.Model = append(create.Model[:index], create.Model[index+1:]...)
	fmt.Println("data Berhasil Dihapus")
}