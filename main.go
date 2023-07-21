package main

import (
	"fmt"
)

func main() {
	ns := GetNutritionalScore(NutritionalData{
		Energy:              EnergyFromKcal(0),
		Sugars:              SugarGram(10),
		SaturatedFattyAcids: SaturatedFattyAcidsGram(2),
		Sodium:              SodiumMilligram(500),
		Fruits:              FruitsPercet(60),
		Fiber:               FiberGram(4),
		Protein:             ProteinGram(2),
	}, Food)

	fmt.Printf("Nutritional Score: %d\n", ns.Value)
	// optional to print NutriScore
	// fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())
}
