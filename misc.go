package nested_sampling

import "math"

// Will my arrays be copied? I hope not? Doesn't work with *
func logsumexp(xs []float64) float64 {

    // Find maximum value
    max := math.Inf(-1) // Minus infinity
    for _, x := range xs {
        if x >= max {
            max = x
        }
    }

    // Sum exp
    tot := 0.0
    for _, x := range xs {
        tot += math.Exp(x - max)
    }

    return max + math.Log(tot)
}
