package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
    parseMeasurements("dummy.txt")

}

func parseMeasurements(path string){
	file,err := os.Open(path)

	if err != nil{
		panic("couldn't open file")
	}

	scanner := bufio.NewScanner(file)

	var stations []station 

	for scanner.Scan() {
		
		line := strings.Split(scanner.Text(),";")
		existing_station := stationExists(stations,line)
		value,err := strconv.ParseFloat(line[1],32)
			
		if err != nil{
			panic("Couldn't parse float")
		}

		if existing_station == nil {
			stations = append(
				stations, 
				station{
					name: line[0],
					max: value,
					min: value,
					count: 1,
					sum: value,
	
				},
			) 
		} else {
			calculate(existing_station,value)
		}
	}

}

func stationExists(stations []station,line []string) *station{
	for i,station := range stations{
		if station.name == line[0] {
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
}

