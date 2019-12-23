package refl

import(
	"sabr"
)

type Reflection struct {
	x, y float64
}

func NewReflection(x float64, y float64) *Reflection {
	refl := Reflection{x, y}
	return &refl
}

func (refl *Reflection) reflectPaths(simuConfig *sabr.SimulationConfig, pathsS[][]float64, pathsSigma[][]float64) ([][]float64, [][]float64) {
	var numberSimus = simuConfig.NumberSimus
	var timeSteps = simuConfig.TimeSteps
	for iSimu := 0; iSimu < numberSimus; iSimu++ {
		for iTime := 1; iTime < timeSteps; iTime++ {
			pathsS[iSimu][iTime] = .0
			pathsSigma[iSimu][iTime] = .0
		}
	}
	return pathsS, pathsSigma
}