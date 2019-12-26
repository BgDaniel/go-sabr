package cumu

import(
	"math"
	"fmt"
)

type Distr struct {
	Samples [][]float64
	Order int
	Moments [][]float64
	NumberSimus int
	TimeSteps int
}

func NewDistr(samples [][]float64, order int) *Distr {
	var numberSimus = len(samples)
	var timeSteps = len(samples[0])
	var moments = make([][]float64, order + 1)

	for m := 1; m <= order; m++ {
		moments[m] = make([]float64, timeSteps)
		for iTime := 0; iTime < timeSteps; iTime++ {
			for iSimu := 0; iSimu < numberSimus; iSimu++ {
				moments[m][iTime] += math.Pow(samples[iSimu][iTime], float64(m))
			}
			moments[m][iTime] /= float64(numberSimus)
		}
	}

	distr := Distr{samples, order, moments, numberSimus, timeSteps}
	return &distr
}

func (distr *Distr) Cumulant(order int) (cumulant []float64) {
	switch order {
	case 1:
		return distr.FirstCumulant()
	case 2:
		return distr.SecondCumulant()
	case 3:
		return distr.ThirdCumulant()
	case 4:
		return distr.FourthCumulant()
	case 5:
		return distr.FifthCumulant()
	default:
		var e = InvalidArgumentError{"Invalid argument: order must be samller or equal to 5!"}
		e.Error()
	}
	return nil
}

type InvalidArgumentError struct {
	ErrorMessage string
}

func (e *InvalidArgumentError ) Error() string {
    return fmt.Sprintf(e.ErrorMessage)
}

func (distr *Distr) FirstCumulant() (cumulant []float64) {
	return distr.Moments[1]
}

func (distr *Distr) SecondCumulant() (cumulant []float64) {
	cumulant = make([]float64, distr.TimeSteps)
	var m1 = distr.Moments[1]
	var m2 = distr.Moments[2]

	for iTime := 0; iTime < distr.TimeSteps; iTime++ {
		cumulant[iTime] = m2[iTime] - Pow(m1[iTime], 2)
	}
	return cumulant
}

func (distr *Distr) ThirdCumulant() (cumulant []float64) {
	cumulant = make([]float64, distr.TimeSteps)
	var m1 = distr.Moments[1]
	var m2 = distr.Moments[2]
	var m3 = distr.Moments[3]

	for iTime := 0; iTime < distr.TimeSteps; iTime++ {
		cumulant[iTime] = m3[iTime] - 3.0 * m2[iTime] * m1[iTime] + 2.0 * Pow(m1[iTime], 3)
	}
	return cumulant
}

func (distr *Distr) FourthCumulant() (cumulant []float64) {
	cumulant = make([]float64, distr.TimeSteps)
	var m1 = distr.Moments[1]
	var m2 = distr.Moments[2]
	var m3 = distr.Moments[3]
	var m4 = distr.Moments[4]

	for iTime := 0; iTime < distr.TimeSteps; iTime++ {
		cumulant[iTime] = m4[iTime] - 4.0 * m3[iTime] * m1[iTime] - 3.0 * Pow(m2[iTime], 2) +
			12.0 * m2[iTime] * Pow(m1[iTime], 2) - 6.0 * Pow(m1[iTime], 4)
	}
	return cumulant
}

func (distr *Distr) FifthCumulant() (cumulant []float64) {
	cumulant = make([]float64, distr.TimeSteps)
	var m1 = distr.Moments[1]
	var m2 = distr.Moments[2]
	var m3 = distr.Moments[3]
	var m4 = distr.Moments[4]
	var m5 = distr.Moments[5]

	for iTime := 0; iTime < distr.TimeSteps; iTime++ {
		cumulant[iTime] = m5[iTime] + 5.0 * m1[iTime] * (6.0 * Pow(m2[iTime], 2) - m4[iTime]) -
			10.0 * m3[iTime] * m2[iTime] + 20.0 * m3[iTime] * Pow(m1[iTime], 2) - 60.0 * m2[iTime] * Pow(m2[iTime], 3) +
			24.0 * Pow(m1[iTime], 5)
	}
	return cumulant
}

func Pow(x float64, n int) (Pow float64){
	if n == 0 {
		return 1.0
	} else if n == 1 {
		return x
	} else {
		var Pow = x
		for i := 2; i <= n; i++ {
			Pow *= x
		}
		return Pow 
	}
}