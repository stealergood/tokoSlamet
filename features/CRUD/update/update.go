package update

import (
	"fmt"
	"toko-slamet/features/CRUD/create"
)

func Update(index, stock int) {
	newStock := create.Model[index].Stock + stock
	create.Model[index].Stock = newStock
	fmt.Println("Stock Berhasil Di Update")
}