package main

import "fmt"

func main() {
	const (
		hargaAwal float32 = 370000
		discont float32 = 15
	)

	hargaDiscont := (discont / 100) * hargaAwal
	hargaAkhir := hargaAwal - hargaDiscont

	fmt.Println("Total harga = Rp.", hargaAkhir);

}