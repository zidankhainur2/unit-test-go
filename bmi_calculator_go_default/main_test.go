package main

import (
	"math"
	"testing"
)

func TestCalculateBMI(t *testing.T) {
	tests := []struct {
		weight, height    float64
		expectedBMI       float64
		expectedCategory  string
	}{
		{50, 1.60, 19.53, "Normal"},
		{70, 1.75, 22.86, "Normal"},
		{80, 1.60, 31.25, "Obesitas"},
		{45, 1.70, 15.57, "Kekurangan Berat Badan"},
		{0, 1.70, 0, "Invalid input"},
		{70, 0, 0, "Invalid input"},
	}

	for _, tt := range tests {
		t.Run("BMI calculation", func(t *testing.T) {
			bmi, category := CalculateBMI(tt.weight, tt.height)

			bmi = math.Round(bmi*100) / 100

			if bmi != tt.expectedBMI {
				t.Errorf("expected BMI %.2f, got %.2f", tt.expectedBMI, bmi)
			}
			if category != tt.expectedCategory {
				t.Errorf("expected category %s, got %s", tt.expectedCategory, category)
			}
		})
	}
}
