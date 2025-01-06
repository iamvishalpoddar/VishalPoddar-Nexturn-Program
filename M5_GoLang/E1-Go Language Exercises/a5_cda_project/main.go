package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CityClimate struct {
	Name         string
	Temperature  float64
	Rainfall     float64
}
var climateData = []CityClimate{
	{Name: "Mumbai", Temperature: 30.5, Rainfall: 2000.0},
	{Name: "Delhi", Temperature: 35.0, Rainfall: 800.0},
	{Name: "Chennai", Temperature: 32.0, Rainfall: 1200.0},
	{Name: "Kolkata", Temperature: 29.5, Rainfall: 1500.0},
	{Name: "Bangalore", Temperature: 25.0, Rainfall: 1000.0},
}
func FindHighestTemperature(data []CityClimate) CityClimate {
	highest := data[0]
	for _, city := range data {
		if city.Temperature > highest.Temperature {
			highest = city
		}
	}
	return highest
}
func FindLowestTemperature(data []CityClimate) CityClimate {
	lowest := data[0]
	for _, city := range data {
		if city.Temperature < lowest.Temperature {
			lowest = city
		}
	}
	return lowest
}
func CalculateAverageRainfall(data []CityClimate) float64 {
	totalRainfall := 0.0
	for _, city := range data {
		totalRainfall += city.Rainfall
	}
	return totalRainfall / float64(len(data))
}
func FilterCitiesByRainfall(data []CityClimate, threshold float64) []CityClimate {
	var filteredCities []CityClimate
	for _, city := range data {
		if city.Rainfall > threshold {
			filteredCities = append(filteredCities, city)
		}
	}
	return filteredCities
}
func SearchCityByName(data []CityClimate, name string) (*CityClimate, error) {
	for _, city := range data {
		if strings.EqualFold(city.Name, name) {
			return &city, nil
		}
	}
	return nil, errors.New("city not found")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	highest := FindHighestTemperature(climateData)
	fmt.Printf("City with the highest temperature: %s (%.2f°C)\n", highest.Name, highest.Temperature)
	lowest := FindLowestTemperature(climateData)
	fmt.Printf("City with the lowest temperature: %s (%.2f°C)\n", lowest.Name, lowest.Temperature)
	averageRainfall := CalculateAverageRainfall(climateData)
	fmt.Printf("Average rainfall across all cities: %.2f mm\n", averageRainfall)

	fmt.Println("\nEnter a rainfall threshold to filter cities:")
	if scanner.Scan() {
		thresholdInput := scanner.Text()
		threshold, err := strconv.ParseFloat(thresholdInput, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			return
		}
		filteredCities := FilterCitiesByRainfall(climateData, threshold)
		if len(filteredCities) == 0 {
			fmt.Println("No cities found with rainfall above the threshold.")
		} else {
			fmt.Println("Cities with rainfall above the threshold:")
			for _, city := range filteredCities {
				fmt.Printf("%s: %.2f mm\n", city.Name, city.Rainfall)
			}
		}
	}

	fmt.Println("\nEnter a city name to search:")
	if scanner.Scan() {
		cityName := scanner.Text()
		city, err := SearchCityByName(climateData, cityName)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("City found: %s, Temperature: %.2f°C, Rainfall: %.2f mm\n", city.Name, city.Temperature, city.Rainfall)
		}
	}
}
