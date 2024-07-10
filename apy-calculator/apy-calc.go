package main

import (
	"fmt"
	"math"
)

func apyCalculator()  error {
	var r, initialInvestment float64 // Input: (r) APY user & initialInvestment user
	var n float64 = 365 // Jumlah hari dalam setahun

	// Inputan dari user via keyboard 
	fmt.Printf("\n==============================================================\n")
	fmt.Printf("\nMasukkan total uang yang akan di investasikan: $")
	fmt.Scan(&initialInvestment)
	fmt.Printf("Masukkan APY Tahunan anda (dalam persen): %%")
	fmt.Scan(&r)

	// Rumus perhitungan daily APY dengan input dari user
	convertPercentage  := r / 100
	dailyRate := math.Pow(1+convertPercentage, 1.0/n) - 1 
	totalDailyAPY := math.Pow(1+dailyRate,n) - 1

	// Rumus perhitungan total investasi terhadap APY dengan input dari user
	totalInitialInvestment := (totalDailyAPY * initialInvestment) + initialInvestment

	fmt.Printf("\n==============================================================\n")
	fmt.Printf("\nAPY harian anda (dalam persen): %.2f%%\nAPY harian diatas bisa berubah-ubah tergantung suku bunga majemuk \n", totalDailyAPY*100)
	fmt.Printf("\n==============================================================\n")
	fmt.Printf("\nTotal hasil investasi dalam setahun adalah: $%.2f\n",totalInitialInvestment)
	fmt.Printf("\n==============================================================\n")

	// Rumus penghitungan total investasi yang berkembang setiap bulan
	for hari := 30; hari <= int(n); hari += 30 {
		totalInvestPerMonth := initialInvestment * math.Pow(1+dailyRate, float64(hari))
		fmt.Printf("Total pertumbuhan investasi hari ke: %d sebesar= $%.2f \n",hari, totalInvestPerMonth)
	}
	return nil
}

func main() {
	apyCalculator()
	
}
