package neuron

import (
	"fmt"
)

//InputLayer and the rest of opcions establish
//what layer the current neuron belongs to
const (
	InputLayer = iota
	HideLayer
	OutputLayer
	Bias
)

//Neuron ...
type Neuron interface {
	RandomizeWidth()
	GetWidht() float64
}

//NFactory (NeuronFactory) selects the correct kinf of neuron
func NFactory(kind rune) (Neuron, error) {
	switch kind {
	case InputLayer:
		return new(InputNeuron), nil
	case HideLayer:
		return new(HideNeuron), nil
	case OutputLayer:
		return new(OutputNeuron), nil
	case Bias:
		return new(BiasNeuron), nil
	default:
		return nil, fmt.Errorf(fmt.Sprintf("Option %d not recognized\n", kind))
	}

}
