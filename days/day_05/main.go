package day_05

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// Run function of the daily challenge
func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input))
	}
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	return strconv.Itoa(findLowestLocation(input))
}

// findLowestLocation finds the lowest location number corresponding to any of the initial seeds
func findLowestLocation(input []string) int {
	seeds := getInitialSeeds(input)
	m := math.MaxInt
	for _, seed := range seeds {
		loc := getHumidity2Location(getTemperature2Humidity(getLight2Temperature(getWater2Light(getFertilizer2Water(getSoil2Fertilizer(getSeed2Soil(seed, input), input), input), input), input), input), input)
		if loc < m {
			m = loc
		}
	}
	return m
}

// getInitialSeeds find the initial seeds from the input
func getInitialSeeds(input []string) []int {
	splitter := regexp.MustCompile("\\s+")
	seeds := splitter.Split(strings.Split(input[0], ": ")[1], -1)
	var s []int
	for _, seed := range seeds {
		i, _ := strconv.Atoi(seed)
		s = append(s, i)
	}
	return s
}

// getSeed2Soil find the soil that corresponds to a given seed
func getSeed2Soil(id int, input []string) int {
	return getMappedValue(id, "seed-to-soil map:", input)
}

// getSoil2Fertilizer find the fertilizer that corresponds to a given soil
func getSoil2Fertilizer(id int, input []string) int {
	return getMappedValue(id, "soil-to-fertilizer map:", input)
}

// getFertilizer2Water find the water that corresponds to a given fertilizer
func getFertilizer2Water(id int, input []string) int {
	return getMappedValue(id, "fertilizer-to-water map:", input)
}

// getWater2Light find the light that corresponds to a given water
func getWater2Light(id int, input []string) int {
	return getMappedValue(id, "water-to-light map:", input)
}

// getLight2Temperature find the temperature that corresponds to a given light
func getLight2Temperature(id int, input []string) int {
	return getMappedValue(id, "light-to-temperature map:", input)
}

// getTemperature2Humidity find the humidity that corresponds to a given temperature
func getTemperature2Humidity(id int, input []string) int {
	return getMappedValue(id, "temperature-to-humidity map:", input)
}

// getHumidity2Location find the location that corresponds to a given humidity
func getHumidity2Location(id int, input []string) int {
	return getMappedValue(id, "humidity-to-location map:", input)
}

// getMappedValue find the mapped value that corresponds to a given source
func getMappedValue(val int, cat string, input []string) int {
	mappedValue := val
	for _, i := range getMappedValues(cat, input) {
		if val >= i[1] && val < i[1]+i[2] {
			return val - i[1] + i[0]
		}
	}
	return mappedValue
}

// getMappedValues reads the mapping ranges from the input
func getMappedValues(cat string, input []string) [][]int {
	ok := false
	var output [][]int
	for _, i := range input {
		if ok && len(i) > 0 {
			splitter := regexp.MustCompile("\\s+")
			row := splitter.Split(i, -1)
			target, _ := strconv.Atoi(row[0])
			source, _ := strconv.Atoi(row[1])
			offset, _ := strconv.Atoi(row[2])
			output = append(output, []int{target, source, offset})
		} else if i == cat {
			ok = true
		}
		if ok && i == "" {
			return output
		}
	}
	return output
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}
