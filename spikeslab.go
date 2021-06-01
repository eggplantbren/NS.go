package nested_sampling

import (
    "math"
    rng "math/rand"
)

// Number of dimensions
const size int = 20;

// Various constants
const u float64 = 0.1;
const v float64 = 0.01;
var tau_u float64 = math.Pow(u, -2)
var tau_v float64 = math.Pow(v, -2)
var log_hundred float64 = math.Log(100.0); // C++ would let that be a constexpr!


// A particle
type Particle struct {
    xs [size]float64
}

// Generate a particle from the prior
func generate() Particle {
    result := Particle {}
    for i:=0; i<size; i++ {
        result.xs[i] = rng.Float64()
    }
    return result
}
