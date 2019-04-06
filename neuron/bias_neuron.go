package neuron

import (
	"math/rand"
	"time"
)

//BiasNeuron ...
type BiasNeuron struct {
	kind      int8 //must be initialize
	input     int8
	Width     float64
	Threshold int8
}

//RandomizeWidth ...
func (n *BiasNeuron) RandomizeWidth() {

	rand.Seed(time.Now().UnixNano())

	n.Width = rand.Float64()
}

//GetWidht ...
func (n *BiasNeuron) GetWidht() float64 {
	return n.Width
}
