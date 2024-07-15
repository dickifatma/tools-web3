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
	if _,err := fmt.Scan(&initialInvestment);err != nil || initialInvestment <= 0{
		return fmt.Errorf("Investasi awal tidak valid: %v", err)
	}
	fmt.Printf("Masukkan APY Tahunan anda (dalam persen): %%")
	if _, err := fmt.Scan(&r); err != nil || r <=0 {
		return fmt.Errorf("APY Tahunan tidak valid: %v", err)
	}

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
	if err := apyCalculator(); err != nil {
		fmt.Println("Error", err)
	}
	
}