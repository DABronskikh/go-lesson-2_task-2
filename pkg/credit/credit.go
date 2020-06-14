package credit

import (
	"math"
)

func Calculate(amount int64, term int64, percent int64) (monthly int64, over int64, total int64) {
	percentMonth := Round(float64(percent)/12.00/100.00, 4)
	percentYear := math.Pow(1+percentMonth, float64(term))
	annuity := Round((percentMonth*percentYear)/(percentYear-1), 6)

	monthly = int64(annuity * float64(amount))
	total = monthly * term
	over = total - amount

	return monthly, over, total
}

func Round(x float64, y int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(y))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}
