package main

// Sampling options
const num_particles int = 100;
const mcmc_steps int = 1000;

type Sampler struct {
    particles [num_particles]Particle;
    logls     [num_particles]float64;
}
