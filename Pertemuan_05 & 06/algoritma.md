Soal9. Faktor Bilangan

Soal:
Buatlah program dalam bahasa Go untuk menampilkan faktor dari suatu bilangan. Faktor adalah bilangan yang habis membagi suatu bilangan. Contoh: Faktor dari 15 adalah 1, 3, 5 dan 15. Faktor dari 24 adalah 1, 2, 3, 4, 6, 8, 12, dan 24.

Masukan terdiri dari sebuah bilangan bulat positif N.
Keluaran terdiri dari kumpulan pasangan nilai bulat d dan boolean s yang dipisahkan oleh newline. Di mana d adalah bilangan yang mungkin menjadi faktor dari N, dan s menyatakan apakah d adalah faktor dari N.

Jawaban:
1. Masukkan sebuah bilangan bulat positif N sebagai input
2. Gunakan for loop untuk menjalankan perulangan dari d = 1 hingga d = N
3. Pada setiap iterasi, periksa apakah N habis dibagi oleh d:
    - Jika N habis dibagi d (sisa pembagian N dengan d adalah 0), tetapkan s = true, yang berarti d adalah faktor dari N.
    - Jika tidak, tetapkan s = false, yang berarti d bukan faktor dari N.
4. Tampilkan nilai d dan s sebagai keluaran.
5. langi langkah tersebut hingga d mencapai N.
6. Akhiri program setelah semua nilai d sudah mencapai N