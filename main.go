package main

import (
	"fmt"
)

func main() {
	ns := GetNutritionalScore(NutritionalData{
		Energy:              EnergyFromKcal(),
		Sugars:              SugarGram(),
		SaturatedFattyAcids: SaturatedFattyAcids(),
		Sodium:              SodiumMilligram(),
		Friuits:             FruitsPercet(),
		Fiber:               FiberGram(),
		Proteins:            ProteinGram(),
	}, Food)

	fmt.Printf("Nutritional Score: %d\n", ns.Value)

}
