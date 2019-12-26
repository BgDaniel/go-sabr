package main

import (
	"fmt"
	"plot"
	"refl"
	"sabr"
)

func main() {
	simuConfig := sabr.SimulationConfig{5000, 4000, 2}
	var S0 = 1.0
	var sigma0 = .3
	var volvol = .085
	var correl = .0
	//var order = 3

	sabr0 := sabr.SABR0{S0, sigma0, volvol, correl, simuConfig}

	var pathsS, pathsSigma = sabr0.GeneratePaths()
	var times = simuConfig.GetTimes()
	//distrS := cumu.NewDistr(pathsS, order)
	//distrSigma := cumu.NewDistr(pathsSigma, order)

	reflection0 := refl.Reflection{S0, sigma0, volvol}
	var pathsSMirrored, pathsSigmaMirrored = reflection0.MirrorPaths(sabr0, pathsS, pathsSigma)
	//distrSMirrored := cumu.NewDistr(pathsSMirrored, order)
	//distrSigmaMirrored := cumu.NewDistr(pathsSigmaMirrored, order)

	var x, cdfSTerminal, cdfSigmaTerminal = reflection0.CalcTerminalDistr(sabr0, pathsS, pathsSigma, .0001, .0, 5.0)
	var _, cdfSMirroredTerminal, cdfSigmaMirroredTerminal = reflection0.CalcTerminalDistr(sabr0, pathsSMirrored, pathsSigmaMirrored, .0001, .0, 5.0)
	plot.SaveToDataFile(x, cdfSTerminal, cdfSMirroredTerminal, "plot_data/terminal_distr/distr_S_T_mirrored.data")
	plot.SaveToDataFile(x, cdfSigmaTerminal, cdfSigmaMirroredTerminal, "plot_data/terminal_distr/distr_Sigma_T_mirrored.data")

	reflection1 := refl.Reflection{-0.5, .3, volvol}
	var pathsSReflected, pathsSigmaReflected = reflection1.ReflectPaths(sabr0, pathsS, pathsSigma)
	//distrSReflected := cumu.NewDistr(pathsSReflected, order)
	//distrSigmaReflected := cumu.NewDistr(pathsSigmaReflected, order)

	var _, cdfSReflectedTerminal, cdfSigmaReflectedTerminal = reflection1.CalcTerminalDistr(sabr0, pathsSReflected, pathsSigmaReflected, .0001, .0, 5.0)
	plot.SaveToDataFile(x, cdfSTerminal, cdfSReflectedTerminal, "plot_data/terminal_distr/distr_S_T_reflected.data")
	plot.SaveToDataFile(x, cdfSigmaTerminal, cdfSigmaReflectedTerminal, "plot_data/terminal_distr/distr_Sigma_T_reflected.data")

	/*
		for k := 1; k <= order; k++ {
			var cumulantsS = distrS.Cumulant(k)
			var cumulantsSigma = distrSigma.Cumulant(k)
			var cumulantsSMirrored = distrSMirrored.Cumulant(k)
			var cumulantsSigmaMirrored = distrSigmaMirrored.Cumulant(k)
			var cumulantsSReflected = distrSReflected.Cumulant(k)
			var cumulantsSigmaReflected = distrSigmaReflected.Cumulant(k)

			plot.SaveToDataFile(times, cumulantsS, cumulantsSMirrored, "plot_data/"+strconv.Itoa(k)+"_momentSMirrored.data")
			plot.SaveToDataFile(times, cumulantsSigma, cumulantsSigmaMirrored, "plot_data/"+strconv.Itoa(k)+"_momentSigmaMirrored.data")
			plot.SaveToDataFile(times, cumulantsS, cumulantsSReflected, "plot_data/"+strconv.Itoa(k)+"_momentSReflected.data")
			plot.SaveToDataFile(times, cumulantsSigma, cumulantsSigmaReflected, "plot_data/"+strconv.Itoa(k)+"_momentSigmaReflected.data")
		}
	*/

	distrIn := reflection1.CalcDistrIn(sabr0, pathsS, pathsSigma)
	distrTouched := reflection1.CalcDistrTouch(sabr0, pathsS, pathsSigma)
	plot.SaveToDataFile(times, distrIn, distrTouched, "plot_data/absor_distr/absor_distr.data")

	fmt.Println("Computation finished.")
}
