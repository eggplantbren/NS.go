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
var c_u float64 = -0.5*math.Log(2*math.Pi*u*u)
var c_v float64 = -0.5*math.Log(2*math.Pi*v*v)
var log_hundred float64 = math.Log(100.0);
var shift float64 = 0.0

// A particle
type Particle struct {
    xs [size]float64
}

// Generate a particle from the prior
func Generate() Particle {
    result := Particle {}
    for i:=0; i<size; i++ {
        result.xs[i] = rng.Float64()
    }
    return result
}

// Evaluate log likelihood
func log_likelihood(particle *Particle) float64 {
    f := 0.0
    g := 0.0
    for i:=0; i<size; i++ {
        f += c_u - 0.5*tau_u*math.Pow(particle.xs[i] - 0.5, 2);
        g += c_v - 0.5*tau_v*math.Pow(particle.xs[i] - 0.5 - shift, 2);
    }
    return logsumexp([2]float64{f, log_hundred + g});
}
