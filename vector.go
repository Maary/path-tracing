package assitance

import (
	"math/rand"
)

//Vector ..
type Vector struct {
	X float64
	Y float64
	Z float64
}

//UnitVector normolized
var UnitVector = Vector{1, 1, 1}

//VectorInUnitSphere ..
func VectorInUnitSphere(rnd *rand.Rand) Vector {
	for {
		r := Vector{
			X: rand.Float64(),
			Y: rand.Float64(),
			Z: rand.Float64(),
		}
	}
	return r
}

func (v Vector) 
