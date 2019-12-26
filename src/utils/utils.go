package utils

func Binom(n int, k int) (binom int) {
	if n == k || k == 1 {
		return 1
	} else {
		return int(fac(n) / (fac(k) * fac(n-k)))
	}
}

func fac(n int) (fac int) {
	if n == 0 || n == 1 {
		return 1
	} else {
		var fac = 1
		for i := 2; i <= n; i++ {
			fac *= i
		}
		return fac
	}
}

func MinMax(samples []float64) (min float64, max float64) {
	min = samples[0]
	max = samples[0]

	for i := 1; i < len(samples); i++ {
		if min > samples[i] {
			min = samples[i]
		}
		if max < samples[i] {
			max = samples[i]
		}
	}

	return min, max
}

func CDF(samples []float64, stepWidth float64, min float64, max float64) (x []float64, cdf []float64) {

	var numberSteps = int((max - min) / stepWidth)
	var numberSimus = len(samples)
	cdf = make([]float64, numberSteps)
	x = make([]float64, numberSteps)

	for iStep := 0; iStep < numberSteps; iStep++ {
		x[iStep] = float64(iStep) * stepWidth
		for iSimu := 0; iSimu < numberSimus; iSimu++ {
			if samples[iSimu] <= float64(iStep)*stepWidth {
				cdf[iStep] += 1.0
			}
		}
		cdf[iStep] /= float64(numberSimus)
	}

	return x, cdf
}
