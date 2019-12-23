package main

import(
	"sabr"
)

func main() {
	simuConfig := sabr.SimulationConfig{1000, 500, 1.0}
	var S0 = 1.0
	var sigma0 = 0.3
	var volvol = 0.01
	var correl = .0

	sabr0 := sabr.SABR0{S0, sigma0, volvol, correl, simuConfig}
	sabr0.GeneratePaths()

}


	