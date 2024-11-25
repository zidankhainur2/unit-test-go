package main

import (
	"fmt"
)

func CalculateBMI(weight, height float64) (float64, string) {
	if height <= 0 || weight <= 0 {
		return 0, "Invalid input"
	}

	bmi := weight / (height * height)
	switch {
	case bmi < 18.5:
		return bmi, "Kekurangan Berat Badan"
	case bmi < 24.9:
		return bmi, "Normal"
	case bmi < 29.9:
		return bmi, "Kelebihan Berat Badan"
	}
	return bmi, "Obesitas"
}

func main() {
	var weight, heightCm float64

	fmt.Print("Masukkan berat badan Anda (kg): ")
	if _, err := fmt.Scanln(&weight); err != nil || weight <= 0 {
		fmt.Println("Input berat badan tidak valid. Harap masukkan angka positif.")
		return
	}

	fmt.Print("Masukkan tinggi badan Anda (cm): ")
	if _, err := fmt.Scanln(&heightCm); err != nil || heightCm <= 0 {
		fmt.Println("Input tinggi badan tidak valid. Harap masukkan angka positif.")
		return
	}

	height := heightCm / 100
	bmi, category := CalculateBMI(weight, height)

	fmt.Printf("\nBMI Anda: %.2f\nKategori: %s\n", bmi, category)
}
