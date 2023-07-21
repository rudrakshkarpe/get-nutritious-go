package main

type ScoreType int

const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

type NutritionalScore struct {
	Value     int
	Positve   int
	Negative  int
	ScoreType ScoreType
}

type EnergyKJ float64

type SugarGram float64

type SaturatedFattyAcids float64

type SodiumMilligram float64

type FruitsPercet float64

type FiberGram float64

type ProteinGram float64

type NutritionalDate struct {
	Energy              EnergyKJ
	Sugars              SugarGram
	SaturatedFattyAcids SaturatedFattyAcids
	Sodium              SodiumMilligram
	FruitsPercet        FruitsPercet
	Fiber               FiberGram
	Protein             ProteinGram
	IsWater             bool
}

var energyLevels = []float64{3350, 3015, 2680, 2345, 2010, 1675, 1340, 1005, 670, 335, 0}
var sugarsLevels = []float64{45, 60, 36, 31, 36, 31, 27, 22.5, 18, 13.5, 9, 4.5, 0}
var SaturatedFattyAcids = []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
var SodiumLevels = []float64{900, 810, 720, 630, 540, 450, 360, 270, 180, 90, 0}
var FruitsLevels = []float64{4.7, 3.7, 2.8, 1.9, 0.9, 0}
var FiberLevels = []float64{8, 6.4, 4.8, 3.2, 1.6, 0}
var ProteinLevels = []float64{13.5, 12, 10.5, 9, 7.5, 6, 4.5, 3, 1.5, 0}

var enegeryLevelsBeverage = []float64{}
var sugarLevelsBeverage = []float64{}

// writing outlines for methods

func (e EnergyKJ) GetPoints(st ScoreType) int {
	if st == Beverage {
		return getPointsFromRange(float64(e), enegeryLevelsBeverage)
	}
	return getPointsFromRange(float64(e), energyLevels)
}

func (s SugarGram) GetPoints(st ScoreType) int {
	if st == Beverage {
		return getPointsFromRange(float64(s), sugarLevelsBeverage)
	}
	return getPointsFromRange(float64(s), sugarsLevels)
}

func (s SaturatedFattyAcids) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(s), SaturatedFattyAcids)
}

func (s SodiumMilligram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(s), SodiumLevels)
}

func (f FruitsPercet) GetPoints(st ScoreType) int {
	if st == Beverage {
		if f > 80 {
			return 10
		} else if f > 60 {
			return 2
		}
		return 0
	}
	if f > 80 {
		return 5
	} else if f > 60 {
		return 2
	} else if f > 40 {
		return 1
	}
	return 0
}

func (f FiberGram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(f), FiberLevels)
}

func (p ProteinGram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(p), ProteinLevels)
}

func EnergyFromKcal(kcal float64) EnergyKJ {
	return EnergyKJ(kcal * 4.184)
}

func SoduimFromsalt(SaltMg float64) SodiumMilligram {
	return SodiumMilligram(saltMg / 2.5)
}

func GetNutritionalScore(n NutritionalData, st ScoreType) NutritionalScore {

	value := 0
	positive := 0
	negative := 0

	if st != water {
		fruitPoints := n.Fruits.GetPoints(st) // GetPoints:() is a method not function
		fiberPoints := n.Fiber.GetPoints(st)

		negative = n.Energy.GetPoints(st) + n.Sugars.GetPoints(st) + n.SaturatedFattyAcids.GetPoints(st) + n.Sodium.GetPoints(st)
		positive = fruitPoints + fiberProtein + Protein
	}

	return NutritionalScore{
		Value:     value,
		Positve:   positive,
		Negative:  negative,
		ScoreType: st,
	}
}

func (ns NutritionalScore) GetNutriScore() string {

}

func getPointsFromRange(v float64, steps []float64) int {
	lenSteps := len(steps)
	for i, l := range steps {
		if v > l {
			return lenSteps - i
		}
	}
	return 0
}
