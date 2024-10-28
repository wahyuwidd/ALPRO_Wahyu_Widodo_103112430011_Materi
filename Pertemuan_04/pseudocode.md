program InvestasiSaham
kamus
    type saham <
        harga_beli           : real
        harga_jual           : real
        jumlah_saham         : real
        total_investasi_awal : real
        total_penjualan      : real
        keuntungan_kotor     : real
        biaya_transaksi      : real
        pajak_keuntungan     : real
        keuntungan_bersih    : real
    >
    s : saham
algoritma
    input(s.harga_beli, s.harga_jual, s.jumlah_saham)
    s.total_investasi_awal = s.harga_beli * s.jumlah_saham
    s.total_penjualan = s.harga_jual * s.jumlah_saham
    s.keuntungan_kotor = s.total_penjualan - s.total_investasi_awal
    s.biaya_transaksi = s.total_penjualan * (0.2 / 100)
    if s.keuntungan_kotor > 0 then
        s.pajak_keuntungan = s.keuntungan_kotor * (10.0 / 100)
    endif
    s.keuntungan_bersih = s.keuntungan_kotor - s.biaya_transaksi - s.pajak_keuntungan
    output(s.total_investasi_awal, s.total_penjualan, s.keuntungan_kotor, s.keuntungan_bersih)
endprogram
