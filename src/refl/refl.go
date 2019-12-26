package refl

import (
	"cumu"
	"sabr"
	"utils"
)

type Reflection struct {
	C, Y, VolVol float64
}

func NewReflection(c float64, y float64, volvol float64) *Reflection {
	refl := Reflection{c, y, volvol}
	return &refl
}

func (refl *Reflection) Level(x float64, y float64) (level float64) {
	return cumu.Pow(refl.VolVol*(x-refl.C), 2) + cumu.Pow(y, 2) - cumu.Pow(refl.Y, 2)
}

func (refl *Reflection) Reflect(x float64, y float64) (u float64, v float64) {
	return (cumu.Pow(refl.Y, 2)*(x-refl.C))/(cumu.Pow(refl.VolVol*(x-refl.C), 2)+cumu.Pow(y, 2)) + refl.C,
		(cumu.Pow(refl.Y, 2) * y) / (cumu.Pow(refl.VolVol*(x-refl.C), 2) + cumu.Pow(y, 2))
}

func (refl *Reflection) MirrorPaths(sabr0 sabr.SABR0, pathsS [][]float64, pathsSigma [][]float64) (pathsSReflected [][]float64, pathsSigmaReflected [][]float64) {
	var numberSimus = sabr0.SimuConfig.NumberSimus
	var timeSteps = sabr0.SimuConfig.TimeSteps
	pathsSReflected = make([][]float64, numberSimus)
	pathsSigmaReflected = make([][]float64, numberSimus)

	for iSimu := 0; iSimu < numberSimus; iSimu++ {
		pathsSReflected[iSimu] = make([]float64, timeSteps)
		pathsSigmaReflected[iSimu] = make([]float64, timeSteps)
		for iTime := 0; iTime < timeSteps; iTime++ {
			pathsSReflected[iSimu][iTime], pathsSigmaReflected[iSimu][iTime] = refl.Reflect(pathsS[iSimu][iTime], pathsSigma[iSimu][iTime])
		}
	}
	return pathsSReflected, pathsSigmaReflected
}

func (refl *Reflection) ReflectPaths(sabr0 sabr.SABR0, pathsS [][]float64, pathsSigma [][]float64) (pathsSReflected [][]float64, pathsSigmaReflected [][]float64) {
	var numberSimus = sabr0.SimuConfig.NumberSimus
	var timeSteps = sabr0.SimuConfig.TimeSteps
	pathsSReflected = make([][]float64, numberSimus)
	pathsSigmaReflected = make([][]float64, numberSimus)

	for iSimu := 0; iSimu < numberSimus; iSimu++ {
		pathsSReflected[iSimu] = make([]float64, timeSteps)
		pathsSigmaReflected[iSimu] = make([]float64, timeSteps)
		var initialLevel = refl.Level(pathsS[iSimu][0], pathsSigma[iSimu][0])
		var hasChangeSide = false

		for iTime := 0; iTime < timeSteps; iTime++ {
			if hasChangeSide {
				pathsSReflected[iSimu][iTime], pathsSigmaReflected[iSimu][iTime] = refl.Reflect(pathsS[iSimu][iTime], pathsSigma[iSimu][iTime])
			} else {
				var currentLevel = refl.Level(pathsS[iSimu][iTime], pathsSigma[iSimu][iTime])
				if initialLevel*currentLevel < .0 {
					hasChangeSide = true
					pathsSReflected[iSimu][iTime], pathsSigmaReflected[iSimu][iTime] = refl.Reflect(pathsS[iSimu][iTime], pathsSigma[iSimu][iTime])
				} else {
					pathsSReflected[iSimu][iTime], pathsSigmaReflected[iSimu][iTime] = pathsS[iSimu][iTime], pathsSigma[iSimu][iTime]
				}
			}
		}
	}
	return pathsSReflected, pathsSigmaReflected
}

func (refl *Reflection) CalcDistrIn(sabr0 sabr.SABR0, pathsS [][]float64, pathsSigma [][]float64) (distrIn []float64) {
	var numberSimus = sabr0.SimuConfig.NumberSimus
	var timeSteps = sabr0.SimuConfig.TimeSteps
	distrIn = make([]float64, timeSteps)

	for iSimu := 0; iSimu < numberSimus; iSimu++ {
		var initialLevel = refl.Level(pathsS[iSimu][0], pathsSigma[iSimu][0])

		for iTime := 0; iTime < timeSteps; iTime++ {
			var currentLevel = refl.Level(pathsS[iSimu][iTime], pathsSigma[iSimu][iTime])

			if initialLevel*currentLevel < .0 {
				distrIn[iTime] += 1.0
			}
		}
	}

	for iTime := 0; iTime < timeSteps; iTime++ {
		distrIn[iTime] *= 2.0 / float64(numberSimus)
	}

	return distrIn
}

func (refl *Reflection) CalcDistrTouch(sabr0 sabr.SABR0, pathsS [][]float64, pathsSigma [][]float64) (distrTouched []float64) {
	var numberSimus = sabr0.SimuConfig.NumberSimus
	var timeSteps = sabr0.SimuConfig.TimeSteps
	distrTouched = make([]float64, timeSteps)

	for iSimu := 0; iSimu < numberSimus; iSimu++ {
		var initialLevel = refl.Level(pathsS[iSimu][0], pathsSigma[iSimu][0])
		var hasTouched = false

		for iTime := 0; iTime < timeSteps; iTime++ {
			if hasTouched {
				distrTouched[iTime] += 1.0
			} else {
				var currentLevel = refl.Level(pathsS[iSimu][iTime], pathsSigma[iSimu][iTime])

				if initialLevel*currentLevel < .0 {
					hasTouched = true
					distrTouched[iTime] += 1.0
				}
			}

		}
	}

	for iTime := 0; iTime < timeSteps; iTime++ {
		distrTouched[iTime] /= float64(numberSimus)
	}

	return distrTouched
}

func (refl *Reflection) CalcTerminalDistr(sabr0 sabr.SABR0, pathsS [][]float64, pathsSigma [][]float64, stepWidth float64, min float64, max float64) (x []float64, distrS []float64, distrSigma []float64) {
	var numberSimus = sabr0.SimuConfig.NumberSimus
	var timeSteps = sabr0.SimuConfig.TimeSteps
	var terminalPathsS = make([]float64, numberSimus)
	var terminalPathsSigma = make([]float64, numberSimus)

	for iSimu := 0; iSimu < numberSimus; iSimu++ {
		terminalPathsS[iSimu] = pathsS[iSimu][timeSteps-1]
		terminalPathsSigma[iSimu] = pathsSigma[iSimu][timeSteps-1]
	}

	_, distrS = utils.CDF(terminalPathsS, stepWidth, min, max)
	x, distrSigma = utils.CDF(terminalPathsSigma, stepWidth, min, max)
	return x, distrS, distrSigma
}
