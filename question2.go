package main

import "fmt"

func question2() {
	var total int
	fmt.Print("Total Belanja seorang customer: Rp. ")
	fmt.Scan(&total)

	var pay int
	fmt.Print("Pembeli membayar: Rp. ")
	fmt.Scan(&pay)

	// get the total kembalian
	kembalian := pay - total

	// get the total with rounding
	roundKembalian := kembalian - (kembalian % 100)

	// initialize totalpecahan (for counting how many pecahan we made)
	var totalPecahan []int

	pecahan := []string{"100.000", "50.000", "20.000", "10.000", "5.000", "2.000", "1.000", "500", "200", "100"}

	var result string
	// result condition
	if pay < total {
		result = "False, kurang bayar"
	} else {
		result = "Kembalian yang harus diberikan kasir: Rp. " + fmt.Sprint(kembalian) + ", dibulatkan menjadi Rp. " + fmt.Sprint(roundKembalian)
		totalPecahan = pecahanUang(roundKembalian)
	}

	fmt.Println(result)

	// print the pecahan, if have any kembalian
	if totalPecahan != nil {
		for i := 0; i < len(pecahan); i++ {
			if totalPecahan[i] != 0 {
				bentuk := "lembar"

				if i > 6 {
					bentuk = "koin"
				}

				fmt.Println(fmt.Sprint(totalPecahan[i]) + " " + bentuk + " " + pecahan[i])
			}
		}
	}

}

// count the total of kembalian with pecahan
func pecahanUang(roundKembalian int) []int {
	sumPecahan := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for roundKembalian > 0 {
		if roundKembalian >= 100000 {
			roundKembalian -= 100000
			sumPecahan[0] = sumPecahan[0] + 1
			continue
		}
		if roundKembalian >= 50000 {
			roundKembalian -= 50000
			sumPecahan[1] = sumPecahan[1] + 1
			continue
		}
		if roundKembalian >= 20000 {
			roundKembalian -= 20000
			sumPecahan[2] = sumPecahan[2] + 1
			continue
		}
		if roundKembalian >= 10000 {
			roundKembalian -= 10000
			sumPecahan[3] = sumPecahan[3] + 1
			continue
		}
		if roundKembalian >= 5000 {
			roundKembalian -= 5000
			sumPecahan[4] = sumPecahan[4] + 1
			continue
		}
		if roundKembalian >= 2000 {
			roundKembalian -= 2000
			sumPecahan[5] = sumPecahan[5] + 1
			continue
		}
		if roundKembalian >= 1000 {
			roundKembalian -= 1000
			sumPecahan[6] = sumPecahan[6] + 1
			continue
		}
		if roundKembalian >= 500 {
			roundKembalian -= 500
			sumPecahan[7] = sumPecahan[7] + 1
			continue
		}
		if roundKembalian >= 200 {
			roundKembalian -= 200
			sumPecahan[8] = sumPecahan[8] + 1
			continue
		}
		if roundKembalian >= 100 {
			roundKembalian -= 100
			sumPecahan[9] = sumPecahan[9] + 1
			continue
		}
	}
	return sumPecahan
}
