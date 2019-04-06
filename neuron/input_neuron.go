package neuron

import (
	"math/rand"
	"time"
)

//InputNeuron ... no posee width
type InputNeuron struct {
	Kind      int8 //must be initialize
	Input     int8
	Width     float64
	Threshold int8
}

//RandomizeWidth ...
func (n *InputNeuron) RandomizeWidth() {

	rand.Seed(time.Now().UnixNano())
	n.Width = rand.Float64()
}

//GetWidht ...
func (n *InputNeuron) GetWidht() float64 {
	return n.Width
}
