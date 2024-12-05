package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("locations.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	locations := parsefile(scanner)

	column2Map := makeMap(locations.Column2)
	fmt.Println(findSimilarity(locations.Column1, column2Map))
}

func makeMap(arr []int) map[int]int {
	locationMap := make(map[int]int)

	for _, val := range arr {
		_, exists := locationMap[val]
		if !exists {
			locationMap[val] = 0
		}
		locationMap[val]++
	}
	return locationMap
}

func findSimilarity(arr []int, locationMap map[int]int) int {
	totalScore := 0
	for _, val := range arr {
		similarityScore, exists := locationMap[val]
		if exists {
			totalScore = totalScore + (similarityScore * val)
		}
	}
	return totalScore
}
