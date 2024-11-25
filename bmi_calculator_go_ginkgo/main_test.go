package main

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("CalculateBMI", func() {
	ginkgo.It("should return the correct BMI and category for normal weight", func() {
		bmi, category := CalculateBMI(60, 1.7)
		gomega.Expect(bmi).To(gomega.BeNumerically("~", 20.76, 0.01))
		gomega.Expect(category).To(gomega.Equal("Normal"))
	})

	ginkgo.It("should return the correct BMI and category for underweight", func() {
		bmi, category := CalculateBMI(50, 1.7)
		gomega.Expect(bmi).To(gomega.BeNumerically("~", 17.30, 0.01))
		gomega.Expect(category).To(gomega.Equal("Kekurangan Berat Badan"))
	})

	ginkgo.It("should return the correct BMI and category for overweight", func() {
		bmi, category := CalculateBMI(80, 1.7)
		gomega.Expect(bmi).To(gomega.BeNumerically("~", 27.68, 0.01))
		gomega.Expect(category).To(gomega.Equal("Kelebihan Berat Badan"))
	})

	ginkgo.It("should return the correct BMI and category for obesity", func() {
		bmi, category := CalculateBMI(100, 1.7)
		gomega.Expect(bmi).To(gomega.BeNumerically("~", 34.60, 0.01))
		gomega.Expect(category).To(gomega.Equal("Obesitas"))
	})

	ginkgo.It("should return Invalid input for zero weight", func() {
		bmi, category := CalculateBMI(0, 1.7)
		gomega.Expect(bmi).To(gomega.Equal(0.0))
		gomega.Expect(category).To(gomega.Equal("Invalid input"))
	})

	ginkgo.It("should return Invalid input for zero height", func() {
		bmi, category := CalculateBMI(60, 0)
		gomega.Expect(bmi).To(gomega.Equal(0.0))
		gomega.Expect(category).To(gomega.Equal("Invalid input"))
	})
})
