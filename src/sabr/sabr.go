package sabr

import(
	"math"
	"math/rand"
)

type SimulationConfig struct {
	NumberSimus, TimeSteps int
	Time float64
}

func (simulationConfig *SimulationConfig) GetTimes() (times []float64) {
	var timeSteps = simulationConfig.TimeSteps
	var dt = simulationConfig.Time / float64(timeSteps)
	times = make([]float64, timeSteps)

	for iTime := 0; iTime < timeSteps; iTime++ {
		times[iTime] = float64(iTime) * dt
	}

	return times
}

type SABR0 struct {
	S0, Sigma0, VolVol, Correl float64
	SimuConfig SimulationConfig
}	

func (sabr0 *SABR0) GeneratePaths() (pathsS [][]float64, pathsSigma [][]float64) {
	var numberSimus = sabr0.SimuConfig.NumberSimus
	var timeSteps = sabr0.SimuConfig.TimeSteps
	pathsS = make([][]float64, numberSimus)
	pathsSigma = make([][]float64, numberSimus)
	var dt = sabr0.SimuConfig.Time / float64(timeSteps)
	var dt_sqrt = math.Sqrt(dt)
	var dWS, dWSigma = sabr0.generateDWt()

	for iSimu := 0; iSimu < numberSimus; iSimu++ {
		pathsS[iSimu] = make([]float64, timeSteps)
		pathsSigma[iSimu] = make([]float64, timeSteps)

		// initialize
		pathsS[iSimu][0] = sabr0.S0
		pathsSigma[iSimu][0] = sabr0.Sigma0

		for iTime := 1; iTime < timeSteps; iTime++ {
			pathsS[iSimu][iTime] += pathsS[iSimu][iTime - 1] * (1.0 + pathsSigma[iSimu][iTime - 1] * dt_sqrt * dWS[iSimu][iTime])
			pathsSigma[iSimu][iTime] = pathsSigma[iSimu][iTime - 1] * (1.0 + sabr0.VolVol * dt_sqrt * dWSigma[iSimu][iTime])
		}
	}

	return pathsS, pathsSigma
}

func (sabr0 *SABR0) generateDWt() ([][]float64, [][]float64) {
	var numberSimus = sabr0.SimuConfig.NumberSimus
	var timeSteps = sabr0.SimuConfig.TimeSteps
	var dWS = make([][]float64, numberSimus)
	var dWSigma = make([][]float64, numberSimus)
	
	for iSimu := 0; iSimu < numberSimus; iSimu++ {
		dWS[iSimu] = make([]float64, timeSteps)
		dWSigma[iSimu] = make([]float64, timeSteps)

		for iTime := 0; iTime < timeSteps; iTime++ {
			dWS[iSimu][iTime] = rand.NormFloat64()
			dWSigma[iSimu][iTime] = rand.NormFloat64()
		}
	}
	return dWS, dWSigma
}

func (sabr0 *SABR0) CalcMean(pathsS [][]float64, pathsSigma [][]float64) ([]float64, []float64) {
	var timeSteps = sabr0.SimuConfig.TimeSteps
	var numberSimus = sabr0.SimuConfig.NumberSimus
	var pathsSMean = make([]float64, timeSteps)
	var pathsSigmaMean = make([]float64, timeSteps)
	for iTime := 0; iTime < timeSteps; iTime++ {
		for iSimu:= 0; iSimu < numberSimus; iSimu++ {
			pathsSMean[iTime] += pathsS[iSimu][iTime]
			pathsSigmaMean[iTime] += pathsSigma[iSimu][iTime]
		}
		pathsSMean[iTime] /= float64(numberSimus)
		pathsSigmaMean[iTime] /= float64(numberSimus)
	}
	return pathsSMean, pathsSigmaMean
}	