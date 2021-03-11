package matematica

import (
	"fmt"
	"strconv"
)

// Media é repsonsável por calcular o que você já sabe :)
func Media(numros ...float64) float64 {
	total := 0.0
	for _, num := range numros {
		total += num
	}
	media := total / float64(len(numros))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada
}
