package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	distance := find_distance(locations)
	fmt.Println(distance)

}

func bubblesort(arr []int) []int {

	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}

	}
	return arr
}

func find_distance(locations Locations) float64 {
	sorted1 := bubblesort(locations.Column1)
	sorted2 := bubblesort(locations.Column2)

	fmt.Println("Column1: ", sorted1)
	fmt.Println("Column2: ", sorted2)

	total_dist := 0.0

	for i := 0; i < len(sorted1); i++ {
		dist := math.Abs(float64(sorted1[i]) - float64(sorted2[i]))
		total_dist = total_dist + dist
		fmt.Println(total_dist)
	}

	return total_dist

}
