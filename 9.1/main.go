package main

import (
	"fmt"
)

// Testy wycinkow, czyli slice'ow, czyli dynamicznych tablic,
// ktore moga sie rozrastac w trakcie dzialania programu, w przeciwienstwie do zwyklych tablic,
// ktore maja staly rozmiar i nie moga sie zmieniac
func main() {

	var listagpuNVIDA = []string{"Nvidia GTX 1030", "Nvidia GTX 1060 3GB", "Nvidia GTX 1060 6GB", "Nvidia GTX 1070 8gb", "Nvidia GTX 1080 8GB"}

	var listagpuAMD = []string{"AMD RX 550 2GB", "AMD RX 560 4GB", "AMD RX 570 4GB", "AMD RX 580 8GB", "AMD RX 590 8GB"}

	fmt.Printf("Listy gpu NVIDA: %v\n", listagpuNVIDA)
	fmt.Printf("Listy gpu AMD: %v\n", listagpuAMD)

}
