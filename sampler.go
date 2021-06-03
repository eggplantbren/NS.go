package main

// Sampling options
const num_particles int = 100;
const mcmc_steps int = 1000;

type Sampler struct {
    particles [num_particles]Particle;
    logls     [num_particles]float64;
}


func make_sampler() Sampler {
    var result Sampler;
    for i:=0; i<num_particles; i++ {
        result.particles[i] = Generate()
        result.logls[i] = log_likelihood(&result.particles[i])
    }
    return result;
}
