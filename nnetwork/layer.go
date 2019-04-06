package nnetwork

import (
	"fmt"

	"github.com/HVCode/chappie/activation"
	"github.com/HVCode/chappie/neuron"
)

//Layer ...
type Layer struct {
	Neurons            []neuron.Neuron
	ActivationFunction activation.Activation
}

//SetNeurons ...
func (l *Layer) SetNeurons(NeuronsInLayer int, KindOfLayer rune) {

	l.Neurons = make([]neuron.Neuron, NeuronsInLayer)

	for index := range l.Neurons {
		l.Neurons[index], _ = neuron.NFactory(KindOfLayer)
		l.Neurons[index].RandomizeWidth()
	}

}

//RandomizeWidths ...
func (l *Layer) RandomizeWidths() {

	for _, n := range l.Neurons {
		n.RandomizeWidth()
	}
}

//SendSignal is used to simulate dot product
func (l *Layer) SendSignal(input []float64) []float64 {

	var total = make([]float64, len(l.Neurons))
	var sum float64

	sumatorio := func(width float64) float64 {
		sum = 0
		for _, element := range input {
			sum += (width * element)
		}
		return sum
	}

	for index, element := range l.Neurons {
		total[index] = l.ActivationFunction.Run(sumatorio(element.GetWidht()))
	}
	fmt.Println(total)
	return total
}
