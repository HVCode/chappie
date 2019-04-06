package activation

import (
	"math"
)

//Activation ...
type Activation interface {
	Run(float64) float64
}

//HLT is short for hard limit transfer
type HLT struct{}

//LT is short for linear transfer
type LT struct{}

//LST is short for log sigmoid transfer
type LST struct{}

//HTST is short for hiperbolic tangent sigmoid transfer
type HTST struct{}

//Run ...
func (h HLT) Run(element float64) float64 {
	if element < 0 {
		return 0
	}

	return 1
}

//Run ...
func (h LT) Run(element float64) float64 {
	return element
}

//Run ...
func (h LST) Run(element float64) float64 {
	return 1 / (1 + math.Pow(math.E, -element))
}

//Run ...
func (h HTST) Run(element float64) float64 {
	return (math.Pow(math.E, element) - math.Pow(math.E, -element)) /
		(math.Pow(math.E, element) + math.Pow(math.E, -element))
}
