package main

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

type Locations struct {
	Column1 []int
	Column2 []int
}

func parsefile(scanner *bufio.Scanner) Locations {

	var locations Locations

	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line)

		if len(columns) == 2 {
			col1, err1 := strconv.Atoi(columns[0])
			col2, err2 := strconv.Atoi(columns[1])

			if err1 != nil || err2 != nil {
				log.Fatal("Error converting the string to an integer")
			}

			locations.Column1 = append(locations.Column1, col1)
			locations.Column2 = append(locations.Column2, col2)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return locations
}
