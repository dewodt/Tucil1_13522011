# Tucil1_13522011

## Deskripsi Singkat

Cyberpunk 2077 Breach Protocol adalah minigame meretas pada permainan video Cyberpunk 2077. Minigame ini merupakan simulasi peretasan jaringan local dari ICE (Intrusion Countermeasures Electronics) pada permainan Cyberpunk 2077. Program ini melakukan algoritma brute force untuk mencari path dengan reward yang optimal.

Beberapa istilah dalam permainan ini:

1. Token – terdiri dari dua karakter alfanumerik seperti E9, BD, dan 55.
2. Matriks – terdiri atas token-token yang akan dipilih untuk menyusun urutan kode.
3. Sekuens – sebuah rangkaian token (dua atau lebih) yang harus dicocokkan.
4. Buffer – jumlah maksimal token yang dapat disusun secara sekuensial.

Beberapa aturan permainan ini:

1. Pemain bergerak dengan pola horizontal, vertikal, horizontal, vertikal (bergantian) hingga semua sekuens berhasil dicocokkan atau buffer penuh
2. Pemain memulai dengan memilih satu token pada posisi baris paling atas dari matriks.
3. Sekuens dicocokkan pada token-token yang berada di buffer.
4. Satu token pada buffer dapat digunakan pada lebih dari satu sekuens.
5. Setiap sekuens memiliki bobot hadiah atau reward yang variatif.
6. Sekuens memiliki panjang minimal berupa dua token.

## Syarat Menjalankan

- Go version 1.21.6
- Linux

## Cara Mengkompilasi (build)

1. Clone repository ini lalu buka folder dan pastikan Anda berada di root directory project ini

```bash
git clone https://github.com/dewodt/Tucil1_13522011.git
```

2. Kunjungi `/src` karena program utama `main.go` ada di directory ini

```bash
cd src
```

3. Compile program go ke folder `/bin`

```bash
go build -o ../bin main.go
```

4. Kunjungi directory `/bin` karena hasil kompilasi ada di sini

```bash
cd ../bin
```

5. Jalankan program binary executable di folder `/bin`

```bash
./main
```

## Cara Menjalankan (run)

1. Clone repository ini lalu buka folder dan pastikan Anda berada di root directory project ini

```bash
git clone https://github.com/dewodt/Tucil1_13522011.git
```

2. Kunjungi `/src` karena program utama `main.go` ada di directory ini

```bash
cd src
```

3. Jalankan program utama `main.go`

```bash
go run main.go
```

## Identitas Pembuat

|   NIM    |        Nama         |
| :------: | :-----------------: |
| 13522011 | Dewantoro Triatmojo |
