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

func (refl *Reflection) ReflectPaths(sabr0 *sabr.SABR0, pathsS[][]float64, pathsSigma[][]float64) ([][]float64, [][]float64) {
	var numberSimus = sabr0.SimuConfig.NumberSimus
	var timeSteps = sabr0.SimuConfig.TimeSteps
	//var volvol = sabr0.VolVol
	
	for iSimu := 0; iSimu < numberSimus; iSimu++ {
		for iTime := 1; iTime < timeSteps; iTime++ {
			pathsS[iSimu][iTime] = .0
			pathsSigma[iSimu][iTime] = .0
		}
	}
	return pathsS, pathsSigma
}