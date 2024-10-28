program DigitTerurut
kamus
    bilangan : integer
    d1, d2, d3 : integer
algoritma
    input(bilangan)
    d1 = bilangan / 100
	d2 = bilangan / 10 % 10
	d3 = bilangan % 10
    output(d1 < d2 && d2 < d3 || d1 > d2 && d2 > d3)
endprogram