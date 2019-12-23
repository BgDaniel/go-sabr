package sabr

import(
	"math"
	"math/rand"
)

type SimulationConfig struct {
	NumberSimus, TimeSteps int
	Time float64
}

func NewSimulationConfig(numberSimus int, timeSteps int, time float64) *SimulationConfig {
	simuConfig := SimulationConfig{numberSimus, timeSteps, time}
	return &simuConfig
}

type SABR0 struct {
	S0, Sigma0, VolVol, Correl float64
	SimuConfig SimulationConfig
}

func NewSABR0(S0 float64, sigma0 float64, volvol float64, correl float64, simuConfig SimulationConfig) *SABR0 {
	sabr0 := SABR0{S0, sigma0, volvol, correl, simuConfig}
	return &sabr0
}	

func (sabr0 *SABR0) GeneratePaths() ([][]float64, [][]float64) {
	var numberSimus = sabr0.SimuConfig.NumberSimus
	var timeSteps = sabr0.SimuConfig.TimeSteps
	var pathsS = make([][]float64, numberSimus)
	var pathsSigma = make([][]float64, numberSimus)
	var dt = sabr0.SimuConfig.Time / float64(timeSteps)
	var dt_sqrt = math.Sqrt(dt)
	var dWS, dWSigma = sabr0.generateDWt()

	for iSimu := 0; iSimu < numberSimus; iSimu++ {
		pathsS[iSimu] = make([]float64, timeSteps)
		pathsSigma[iSimu] = make([]float64, timeSteps)

		// initialize
		pathsS[iSimu][0] = sabr0.S0
		pathsSigma[iSimu][0] = sabr0.VolVol

		for iTime := 1; iTime < timeSteps; iTime++ {
			pathsS[iSimu][iTime] += pathsSigma[iSimu][iTime - 1] * dt_sqrt * dWS[iSimu][iTime]
			pathsSigma[iSimu][iTime] = pathsSigma[iSimu][iTime - 1] * sabr0.VolVol * dt_sqrt * dWSigma[iSimu][iTime]
		}
	}

	return pathsS, pathsSigma
}

func (sabr0 *SABR0) generateDWt() ([][]float64, [][]float64){
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

func main() {
	simuConfig := SimulationConfig{1000, 500, 1.0}
	var S0 = 1.0
	var sigma0 = 0.3
	var volvol = 0.01
	var correl = .0

	sabr0 := SABR0{S0, sigma0, volvol, correl, simuConfig}
	sabr0.GeneratePaths()

}


	