package main

import (
	"fmt"
	"plot"
	"refl"
	"sabr"
	"cumu"
	"strconv"
)

func main() {
	simuConfig := sabr.SimulationConfig{1000, 1000, 2}
	var S0 = 1.0
	var sigma0 = .3
	var volvol = .085
	var correl = .0
	var order = 3

	sabr0 := sabr.SABR0{S0, sigma0, volvol, correl, simuConfig}

	var pathsS, pathsSigma = sabr0.GeneratePaths()
	var times = simuConfig.GetTimes()
	distrS := cumu.NewDistr(pathsS, order)
	distrSigma := cumu.NewDistr(pathsSigma, order)	

	reflection0 := refl.Reflection{S0, sigma0, volvol}
	var pathsSMirrored, pathsSigmaMirrored = reflection0.MirrorPaths(sabr0, pathsS, pathsSigma)
	distrSMirrored := cumu.NewDistr(pathsSMirrored, order)
	distrSigmaMirrored := cumu.NewDistr(pathsSigmaMirrored, order)

	reflection1 := refl.Reflection{-0.5, .3, volvol}
	var pathsSReflected, pathsSigmaReflected = reflection1.ReflectPaths(sabr0, pathsS, pathsSigma)
	distrSReflected := cumu.NewDistr(pathsSReflected, order)
	distrSigmaReflected := cumu.NewDistr(pathsSigmaReflected, order)

	for k := 1; k <= order; k++ {
		var cumulantsS = distrS.Cumulant(k)
		var cumulantsSigma = distrSigma.Cumulant(k)
		var cumulantsSMirrored = distrSMirrored.Cumulant(k)
		var cumulantsSigmaMirrored = distrSigmaMirrored.Cumulant(k)
		var cumulantsSReflected = distrSReflected.Cumulant(k)
		var cumulantsSigmaReflected = distrSigmaReflected.Cumulant(k)

		plot.SaveToDataFile(times, cumulantsS, cumulantsSMirrored, strconv.Itoa(k) + "_momentSMirrored.data")
		plot.SaveToDataFile(times, cumulantsSigma, cumulantsSigmaMirrored, strconv.Itoa(k) + "_momentSigmaMirrored.data")
		plot.SaveToDataFile(times, cumulantsS, cumulantsSReflected, strconv.Itoa(k) + "_momentSReflected.data")
		plot.SaveToDataFile(times, cumulantsSigma, cumulantsSigmaReflected, strconv.Itoa(k) + "_momentSigmaReflected.data")
	}

	distrIn := reflection1.CalcDistrIn(sabr0, pathsS, pathsSigma)
	distrTouched := reflection1.CalcDistrTouch(sabr0, pathsS, pathsSigma) 
	plot.SaveToDataFile(times, distrIn, distrTouched, "absor_distr.data")

	fmt.Println("Computation finished.")
}
