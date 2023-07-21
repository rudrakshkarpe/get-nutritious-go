package main

import (
	"fmt"
)

func main() {
	ns := GetNutritionalScore(NutritionalData{
		Energy:              EnergyFromKcal(10),
		Sugars:              SugarGram(20),
		SaturatedFattyAcids: SaturatedFattyAcidsGram(2),
		Sodium:              SodiumMilligram(100),
		Fruits:              FruitsPercet(30),
		Fiber:               FiberGram(42),
		Protein:             ProteinGram(12),
	}, Food)

	fmt.Printf("Nutritional Score: %d\n", ns.Value)
	// optional to print NutriScore
	fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())
}
