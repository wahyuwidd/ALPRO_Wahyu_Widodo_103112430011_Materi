program FaktorBilangan
kamus
    i, N, d : integer
	s : boolean
algoritma
    input(N)
    for i = 1 to N do
        d = i
        s = N%d == 0
        output(d,s)
    endfor
endprogram