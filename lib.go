package main

import (
	"log"
	"os"
	"runtime/pprof"
)

type station struct {
	name  string
	max   float64
	min   float64
	sum   float64
	count float64
	mean  float64
}

func (station *station) calculateMin(num float64) {
	station.min = min(station.min, num)
}

func (station *station) calculateMax(num float64) {
	station.min = min(station.min, num)
}

func (station *station) calculateMean(num float64) {
	station.sum += num
	station.count++
	station.mean = station.sum / station.count
}

func profile(){
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()
}
