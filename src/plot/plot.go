package plot

import (
	"fmt"
	"os"
)

func SaveToDataFile(times []float64, samples []float64, samplesReflected []float64, file string) {
	dataFile, error := os.Create(file)
	if error != nil {
		fmt.Println(error)
		return
	}

	var header = "# \t time \t original \t transformed \n"

	l, error := dataFile.WriteString(header)

	fmt.Println(l, "bytes written successfully")

	if error != nil {
		fmt.Println(error)
		dataFile.Close()
		return
	}

	for iTime := 0; iTime < len(samples); iTime++ {
		var line = fmt.Sprintf("%.2f \t %.4f \t %.4f \n", times[iTime], samples[iTime], samplesReflected[iTime])
		l, error := dataFile.WriteString(line)
		fmt.Println(l, "bytes written successfully")
		if error != nil {
			fmt.Println(error)
			dataFile.Close()
			return
		}
	}

	fmt.Println("Data successfully written to " + file + "!")
	dataFile.Close()
}
