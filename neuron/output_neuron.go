package neuron

import (
	"math/rand"
	"time"
)

//OutputNeuron ...
type OutputNeuron struct {
	kind      int8 //must be initialize
	width     float64
	Threshold int8
}

//RandomizeWidth ...
func (n *OutputNeuron) RandomizeWidth() {

	rand.Seed(time.Now().UnixNano())
	n.width = rand.Float64()
}

//GetWidht ...
func (n *OutputNeuron) GetWidht() float64 {
	return n.width
}
