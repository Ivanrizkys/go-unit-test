package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ! unit test

// ? aturan menjalankan unit test
// * saat akan menjalankan unit test kita harus masuk dulu ke directory package dari function yang akan di test
// * namun jika kita ingin menjalankan semua unit test yang ada di dalam sebuah package maka bisa melakukannya dengan perintah go test ./...

// ? menjalankan unit test
// * go test
// ini akan menjalankan semua function unit test yang kita buat di dalam sebuah package
// * go test -v
// menjalankan semua unit test yang kita buat tetapi juga sambil memberikan informasi function mana yang di eksekusi
// * go test -v -run Testnamafunction
// menjalankan unit test spesific sesuai nama function nya

// ? running uni test di package dan package di dalamnya
// * go test ./...
// ini akan me running function unit test yang ada di dalam package dan pacage di dalamnya
// * go test -v ./..
// sambil memberikan informasi function yang akan di jalankan

// ? aturan penamaan file test
// * harus menggunakan akhiran file _test

// ? Aturan penamaan function unit test
// * namanya harus diawali dengan Test
// * tidak ada aturan nama belakang function harus sama dengan function yang akan di testing
// * harus memiliki parameter (t *testing.T)
// * tidak boleh memiliki return value

// ? menggagalkan seubah unit test
// * t.Fail()
// jika menemui ini maka unit test akan dianggap gagal namun tetap mengeksekusi kode selajutnya di dalam function unit test yang gagal tadi
// * t.FailNow()
// jika menemui ini maka unit test akan dianggap gagal dan eksekusi kode selanjutnya tidak akan dijalankan / masuk ke function unit test yang lain (kalau ada)
// * t.Error()
// perilakunya sama seperti fail tapi bisa sambil memberi informasi error sesuai yang kita berikan
// * t.Fatal
// sama seperti FailNow tapi bisa memberikan informasi error sesuai yang kita berikan

// ? skip unit test
// * di golang kita bisa membatalkan unit test dengan menggunakan t.Skip(message)

// ? function TestMain
// * jika terdapat function ini di dalam unit test maka function ini akan di eksekusi pertama kali dan paling terakhir sebuah unit test
// * penamaanya harus dengan nama TestMain
// * bisa dilakukan untuk perintah before dan after
// * menggunakan type *testing.M

// ? subtest
// * di function unit test kita bisa membuat function lagi yang akan menjadi sub test dari sebuah init test
// * caranya adalah dengan menggunakan t.Run(namaSubtes, function subtest)

// ? table test
// * digunakan untuk membuat unit test agar mencegah perulangan kode yang akan kita buat

// ? package testify

// ? asertion
// * ini adalah pengecekan yang dilakukan dengan package testify dan jika terjadi error maka akan memanggil t.fail()

// ? require
// * pengecekan yang dilakukan dengan package testify dan jika terjadi error maka akan memanggil t.FailNow()

// ! Benchmark

// ? aturan penamaan file bechmark
// * untuk aturan penamaannya sendiri sama dengan unit test yang mana harus memiliki akhiran _test

// ? aturan pembuatan function benchmark
// * harus diawali dengan kata Benchmark
// * contohnya BenchmarkHelloIvan()
// * harus memiliki parameter (b *testing.B)
// * testing.B sendiri memiliki kesamaan dengan testing.T dan memiliki method seperti Fail(), FailNow(), dan sebagainya

// ? menjalankan function benchmark
// * go test -v -bench=.
// akan menjalankan seluruh benchmark di dalam module, tetapi function unit test juga akan dijalankan
// * go test -v -run=NotMathUnitTest -bench=.
// NotMathUnitTest bisa kita ini dengan nama function yang tidak ada dalam unit test di dalam module
// ini akan menjalankan seluruh benchmark yang ada
// * go test -v run=NotMathUnitTest -bench=BechmarkTest
// BenchmarkTest bisa kita isi dengan function benchmark yang akan kita jalankan
// ini akan menjalankan function benchmark tertentu
// * go test -v -bench=. ./..
// menjalankan benchmark di root project

// ? sub benchmark
// * kita juga bisa melakukan sub benchmark seperti di unit test
// * caranya juga sama dan kita hanya perlu memanggil method b.run()

// ? table benchmark
// * kita juga bisa membuat table benchmark
// * caranya juga sama seperti yang ada di unit test

func TestMain(m *testing.M) {
	// * before unit test
	fmt.Println("Before unit test")

	// * run semua unit test
	m.Run()

	// * after unit test
	fmt.Println("After unit test")
}

func TestHelloIvan(t *testing.T) {
	result := HelloWord("Ivan")
	if result != "Hallo Ivan" {
		// ! unit test failed
		panic("Result is not Hallo Ivan")
	}
}
func TestHelloAis(t *testing.T) {
	result := HelloWord("Ais")
	if result != "Hai Ais" {
		// ! unit test failed
		t.Error("Result not Hai Ais")
	}
	// * kode program akan dijalankan
	fmt.Println("Kode Program dijalankan")
}
func TestHelloHilmi(t *testing.T) {
	result := HelloWord("Hilmi")
	if result != "Hai Hilmi" {
		// ! unit test failed
		t.Fatal("Result is not Hallo Hilmi")
	}
	// * kode program tidak akan dijalankan
	fmt.Print("Kode program dijalankan")
}

// ? skip unit test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cannot running in windows")
	}
	// * kode di bawah tidak akan di eksekusi
	fmt.Println("Kode program di jalankan")
}

// ? asertion
func TestHelloSari(t *testing.T) {
	result := HelloWord("Sari")
	// ! jika parameter dua dan tiga tidak sama maka akan error
	assert.Equal(t, "Hai Sari", result, "Must be equal")
	// * kode program akan dijalankan
	fmt.Print("Kode program dijalankan")
}

// ? require
func TestHelloFahmi(t *testing.T) {
	result := HelloWord("Fahmi")
	// ! jika parameter dua dan tiga tidak sama maka akan error
	require.Equal(t, "Hai Fahmi", result, "Must be equal")
	// * kode program tidak akan dijalankan
	fmt.Print("Kode program dijalankan")
}

// ? subtest
func TestSubTest(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		result := HelloWord("Sari")
		assert.Equal(t, "Hallo Sari", result, "Must be equal")
		// * kode program akan dijalankan
		fmt.Print("Kode program dijalankan")
	})
	t.Run("test2", func(t *testing.T) {
		result := HelloWord("Sari")
		assert.Equal(t, "Hai Sari", result, "Must be equal")
		// * kode program akan dijalankan
		fmt.Print("Kode program dijalankan")
	})
}

// ? table test
func TestTableTest(t *testing.T) {
	test := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "ivan",
			request:  "Ivan",
			expected: "Hallo Ivan",
		},
		{
			name:     "wahyu",
			request:  "Wahyu",
			expected: "Hallo Wahyu",
		},
		{
			name:     "sania",
			request:  "Sania",
			expected: "Hallo Sania",
		},
	}

	for _, value := range test {
		t.Run(value.name, func(t *testing.T) {
			result := HelloWord(value.request)
			require.Equal(t, value.expected, result, "Must be equal")
		})
	}
}

// ? benchmark
func BenchmarkHelloIvan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWord("Ivan")
	}
}

func BenchmarkHelloAis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWord("Ais")
	}
}
