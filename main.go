package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("hello world")
	file,err := os.Open("./measurements.txt")

	if err != nil{
		panic("couldn't open file")
	}

	scanner := bufio.NewScanner(file)

	var stations []station 

	for scanner.Scan(){
		line := strings.Split(scanner.Text(),";")

		if !stationExists(stations,line) {
			stations = append(
				stations, 
				station{
					name: line[0],
					max: 20.0,
					min:20.0,
					count: 10,
					sum: 500,
	
				},
			)
		} else {

		}

		
	}
}

func stationExists(stations []station,line []string) bool{
	for _,station := range stations{
		if station.name == line[0] {
			return true
		}
	}

	return false
}

type station struct{
	name string
	max float32
	min float32
	sum int
	count int
}