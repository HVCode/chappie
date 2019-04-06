package neuron

import (
	"math/rand"
	"time"
)

//HideNeuron ...
type HideNeuron struct {
	kind      int8 //must be initialize
	input     int8
	width     float64
	Threshold int8
}

//RandomizeWidth ...
func (n *HideNeuron) RandomizeWidth() {

	rand.Seed(time.Now().UnixNano())
	n.width = rand.Float64()
}

//GetWidht ...
func (n *HideNeuron) GetWidht() float64 {
	return n.width
}
