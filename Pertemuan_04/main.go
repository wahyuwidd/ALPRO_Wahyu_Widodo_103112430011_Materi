package main

import "fmt"

type saham struct {
	harga_beli           float64
	harga_jual           float64
	jumlah_saham         float64
	total_investasi_awal float64
	total_penjualan      float64
	keuntungan_kotor     float64
	biaya_transaksi      float64
	pajak_keuntungan     float64
	keuntungan_bersih    float64
}

func main() {
	var s saham

	//input
	fmt.Println("Informasi Investasi Saham:")
	fmt.Print("Harga beli: Rp ")
	fmt.Scan(&s.harga_beli)
	fmt.Print("Harga jual: Rp ")
	fmt.Scan(&s.harga_jual)
	fmt.Print("Jumlah saham: ")
	fmt.Scan(&s.jumlah_saham)

	//rumus menghitung keuntungan bersih
	s.total_investasi_awal = (s.harga_beli * s.jumlah_saham)
	s.total_penjualan = (s.harga_jual * s.jumlah_saham)
	s.keuntungan_kotor = (s.total_penjualan - s.total_investasi_awal)
	s.biaya_transaksi = (s.total_penjualan * (0.2 / 100))
	if s.keuntungan_kotor > 0 {
		s.pajak_keuntungan = (s.keuntungan_kotor * (10.0 / 100))
	}
	s.keuntungan_bersih = (s.keuntungan_kotor - s.biaya_transaksi - s.pajak_keuntungan)

	//output
	fmt.Println("Total Investasi Awal : ", int(s.total_investasi_awal))
	fmt.Println("Total Penjualan : ", int(s.total_penjualan))
	fmt.Println("Keuntungan Kotor : ", int(s.keuntungan_kotor))
	fmt.Println("Biaya Transaksi : ", int(s.biaya_transaksi))
	fmt.Println("Pajak Keuntungan : ", int(s.pajak_keuntungan))
	fmt.Println("Keuntungan Bersih : ", int(s.keuntungan_bersih))
}
