package main

import "fmt"

func main() {
	// fmt.Println("")
	// fmt.Println("LATIHAN 1 – Hello World + Variable")
	// fmt.Println("Hello, Go!")

	// var nama = "Fikri"
	// var umur = 23
	// fmt.Println("Nama saya", nama, ", umur saya", umur)
	// fmt.Println("")

	// fmt.Println("LATIHAN 2 – Tipe Data Number & Operasi Matematika")
	// var a = 10
	// var b = 3
	// var penjumlahan = a + b
	// var pengurangan = a - b
	// var perkalian = a * b
	// var pembagian = a / b
	// fmt.Println("Hasil penjumlahan:", penjumlahan)
	// fmt.Println("Hasil pengurangan:", pengurangan)
	// fmt.Println("Hasil perkalian:", perkalian)
	// fmt.Println("Hasil pembagian:", pembagian)
	// fmt.Println("")

	// fmt.Println("LATIHAN 3 – Boolean & Operasi Perbandingan")
	// var nilai = 75
	// var lulus = nilai >= 75
	// var tidakLulus = nilai < 75
	// fmt.Println("Nilai:", nilai)
	// fmt.Println("Lulus:", lulus)
	// fmt.Println("Tidak Lulus:", tidakLulus)
	// fmt.Println("")

	// fmt.Println("LATIHAN 4 – Operasi Boolean (AND, OR, NOT)")
	// var hadir = true
	// var nilaiAkhir = 80
	// var kelulusan = hadir && nilaiAkhir >= 75
	// var remedial = !kelulusan || nilaiAkhir < 75
	// fmt.Println("Hadir:", hadir)
	// fmt.Println("Nilai Akhir:", nilaiAkhir)
	// fmt.Println("Kelulusan:", kelulusan)
	// fmt.Println("Remedial:", remedial)
	// fmt.Println("")

	// fmt.Println("LATIHAN 5 – Constant & Type Declaration")
	// const phi = 3.14
	// const r = 7
	// var luasLingkaran = phi * r * r
	// fmt.Println("phi:", phi)
	// fmt.Println("jari-jari (r):", r)
	// fmt.Println("Luas Lingkaran dengan jari-jari", r, "adalah:", luasLingkaran)
	// type Umur int
	// var usia Umur = 23
	// fmt.Println("Usia:", usia)
	// fmt.Println("")

	// fmt.Println("LATIHAN 6 – Konversi Tipe Data")
	// var panjang int16 = 10
	// var lebar int16 = 3
	// var panjangPersegi = float32(panjang)
	// var lebarPersegi = float32(lebar)
	// var luasPersegiPanjang = panjangPersegi * lebarPersegi
	// fmt.Println("Panjang:", panjang)
	// fmt.Println("Lebar:", lebar)
	// fmt.Println("Luas Persegi Panjang:", luasPersegiPanjang)
	// fmt.Println("")

	// 	🔥 LATIHAN BONUS (WAJIB BIAR NAIK LEVEL)
	// ❓ Pertanyaan Refleksi (JANGAN DISKIP):

	// Jawab sendiri (boleh tulis di kertas):
	// 1. Kenapa Go melarang 2 main() dalam 1 folder?
	// 2. Kenapa hasil pembagian integer tidak ada koma?
	// 3. Kapan harus pakai const, kapan var?

	// Jawaban singkat:
	// 1. lupa
	// 2. belum paham
	// 3. pake const kalo nilainya enggak bakal berubah, pake var kalo nilainya bisa berubah

	// var names = [...]string{
	// 	"Fikri",
	// 	"Andi",
	// 	"Siti",
	// 	"Budi",
	// 	"Rina",
	// 	"Lina",
	// 	"Agus",
	// 	"",
	// 	"",
	// }

	// var values = [...]int{
	// 	90,
	// 	85,
	// 	78,
	// 	88,
	// 	92,
	// 	80,
	// 	75,
	// }

	// fmt.Println("")
	// namesSlice := names[2:5]
	// valuesSlice := values[2:5]
	// fmt.Println("Names Slice:", namesSlice)
	// fmt.Println("Values Slice:", valuesSlice)
	// fmt.Println("")
	// namesSlice[0] = "Joko"
	// valuesSlice[0] = 100
	// fmt.Println("Names Slice:", namesSlice)
	// fmt.Println("Values Slice:", valuesSlice)
	// fmt.Println("")
	// fmt.Println("Names Array:", names)
	// fmt.Println("Values Array:", values)
	// fmt.Println("")

	// namesSlice2 := namesSlice[0:4]
	// valuesSlice2 := valuesSlice[0:4]
	// fmt.Println("Names Slice 2:", namesSlice2)
	// fmt.Println("Values Slice 2:", valuesSlice2)
	// fmt.Println("")
	// namesSlice2[1] = "Lala"
	// valuesSlice2[1] = 60
	// fmt.Println("Names Slice 2:", namesSlice2)
	// fmt.Println("Values Slice 2:", valuesSlice2)
	// fmt.Println("")
	// fmt.Println("Names Slice:", namesSlice)
	// fmt.Println("Values Slice:", valuesSlice)
	// fmt.Println("")
	// fmt.Println(len(namesSlice), cap(namesSlice))
	// fmt.Println(len(valuesSlice), cap(valuesSlice))
	// fmt.Println(len(namesSlice2), cap(namesSlice2))
	// fmt.Println(len(valuesSlice2), cap(valuesSlice2))
	// fmt.Println("")
	// fmt.Println("Names Array:", names)
	// fmt.Println("Values Array:", values)
	// fmt.Println("")

	// namesSlice3 := append(namesSlice2, "Dodo")
	// valuesSlice3 := append(valuesSlice2, 55)
	// fmt.Println("Names Slice 3:", namesSlice3)
	// fmt.Println("Values Slice 3:", valuesSlice3)
	// fmt.Println("")
	// fmt.Println("Names Slice 2:", namesSlice2)
	// fmt.Println("Values Slice 2:", valuesSlice2)
	// fmt.Println("")
	// fmt.Println("Names Array:", names)
	// fmt.Println("Values Array:", values)
	// fmt.Println(len(names))
	// fmt.Println(cap(names))

	// namesSlice4 := make([]string, 0, 5)
	// valuesSlice4 := make([]int, 0, 5)
	// namesSlice4 = append(namesSlice4, "Asep", "Ucup", "Otong")
	// valuesSlice4 = append(valuesSlice4, 70, 80, 90)
	// fmt.Println("Names Slice 4:", namesSlice4)
	// fmt.Println("Values Slice 4:", valuesSlice4)
	// fmt.Println(len(namesSlice4), cap(namesSlice4))
	// fmt.Println(len(valuesSlice4), cap(valuesSlice4))
	// fmt.Println("")

	// namesSlice[0] = "Joko"
	// valuesSlice[0] = 100
	// fmt.Println("Names Slice:", namesSlice)
	// fmt.Println("Values Slice:", valuesSlice)

	// namesSlice2 := namesSlice[0:3]
	// valuesSlice2 := valuesSlice[0:3]
	// fmt.Println("Names Slice 2:", namesSlice2)
	// fmt.Println("Values Slice 2:", valuesSlice2)

	// namesSlice2[1] = "Lala"
	// valuesSlice2[1] = 60
	// fmt.Println("Names Slice 2:", namesSlice2)
	// fmt.Println("Values Slice 2:", valuesSlice2)
	// fmt.Println("Names Slice:", namesSlice)
	// fmt.Println("Values Slice:", valuesSlice)
	// fmt.Println(len(namesSlice), cap(namesSlice))
	// fmt.Println(len(valuesSlice), cap(valuesSlice))
	// fmt.Println(len(namesSlice2), cap(namesSlice2))
	// fmt.Println(len(valuesSlice2), cap(valuesSlice2))

	// nilai := []int{80, 75, 90, 60, 85}
	// fmt.Println("")
	// fmt.Println("Nilai siswa:", nilai)
	// fmt.Println("Nilai pertama:", nilai[0])
	// fmt.Println("Nilai terakhir:", nilai[4])
	// fmt.Println("Jumlah nilai:", nilai[0]+nilai[1]+nilai[2]+nilai[3]+nilai[4])
	// fmt.Println("")

	// siswa := make([]string, 0, 5)
	// siswa = append(siswa, "Andi", "Budi", "Citra")
	// fmt.Println("Daftar siswa:", siswa)
	// fmt.Println("Jumlah siswa:", len(siswa))
	// fmt.Println("Kapasitas slice siswa:", cap(siswa))
	// fmt.Println("")

	// Nilai := map[string]int{
	// 	"math": 80,
	// 	"english":  75,
	// 	"science":  90,
	// }
	// fmt.Println("Nilai map siswa:", Nilai)
	// fmt.Println("Nilai math:", Nilai["math"])
	// fmt.Println("Nilai english:", Nilai["english"])
	// fmt.Println("Nilai science:", Nilai["science"])
	// fmt.Println("")

	// nilai := 78
	// if skor := nilai; skor >= 75 {
	// 	fmt.Println("LULUS")
	// } else {
	// 	fmt.Println("TIDAK LULUS")
	// }

	// nilai := 4
	// switch nilai {
	// case 1:
	// 	fmt.Println("Sangat Buruk")
	// case 2:
	// 	fmt.Println("Buruk")
	// case 3:
	// 	fmt.Println("Cukup")
	// case 4:
	// 	fmt.Println("Baik")
	// case 5:
	// 	fmt.Println("Sangat Baik")
	// default:
	// 	fmt.Println("Nilai tidak valid")
	// }

	names := []string{"anto", "dewi", "nanda", "junior"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
}
