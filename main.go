package main

import (
	"bufio"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

func main() {

	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

    //parseMeasurements("text-files/dummy.txt")
    parseMeasurementsV2("text-files/dummy.txt")
}

func parseMeasurements(path string){
	file,err := os.Open(path)

	if err != nil{
		panic("couldn't open file")
	}

	scanner := bufio.NewScanner(file)

	var stations []station 

	for scanner.Scan() {
		
		index := strings.Index(scanner.Text(),";")
		name := scanner.Text()[:index]
		
		value,err := strconv.ParseFloat(scanner.Text()[index+1:],32)

		existing_station := stationExists(stations,name)
			
		if err != nil{
			panic("Couldn't parse float")
		}

		if existing_station == nil {
			stations = append(
				stations, 
				station{
					name: name,
					max: value,
					min: value,
					count: 1,
					sum: value,
	
				},
			) 
		} else {
			existing_station.calculateMax(value)
			existing_station.calculateMin(value)
			existing_station.calculateMean(value)
		}
	}	
}

func parseMeasurementsV2(path string){
	file,err := os.Open(path)

	if err != nil{
		panic("couldn't open file")
	}

	scanner := bufio.NewScanner(file)

	stations := make(map[string]station)

	for scanner.Scan() {
		
		index := strings.Index(scanner.Text(),";")
		name := scanner.Text()[:index]
		
		value,err := strconv.ParseFloat(scanner.Text()[index+1:],32)

		if err != nil{
			panic("Couldn't parse float")
		}

		if existing_station, ok := stations[name]; ok {
			existing_station.calculateMax(value)
			existing_station.calculateMin(value)
			existing_station.calculateMean(value)
			stations[name] = existing_station
		} else{
			stations[name] = station{
				name: name,
				max: value,
				min: value,
				count: 1,
				sum: value,

			}
		}
	}	
}

func stationExists(stations []station,name string) *station{
	//TODO find a faster method to check stations
	for i,station := range stations{
		if station.name == name {
			return &stations[i]
		}
	}

	return nil
}

func calculate(station *station,num float64){
	station.min = min(station.min,num)
	station.max = max(station.max,num)
	station.sum += num
	station.count ++
	station.mean = station.sum / station.count
}

