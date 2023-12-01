package service

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type RealDatabase struct {
	File   string
	reader io.Reader
}

func NewRealDatabase() *RealDatabase {
	input := `id name age subordinateIds
1 waterball 25 
2 fixiabis 15 1,3
3 fong 7 1
4 cc 18 1,2,3
5 peterchen 3 1,4
6 handsomeboy 22 1`

	reader := strings.NewReader(input)

	return &RealDatabase{reader: reader}
}

func (d *RealDatabase) GetEmployeeById(target int) (*RealEmployee, error) {
	// Create a function to load person data on-demand
	scanner := bufio.NewScanner(d.reader)

	// Skip the header line
	if !scanner.Scan() {
		return nil, fmt.Errorf("failed to read header")
	}

	// Process each line until finding the specified person ID
	for lineNumber := 1; scanner.Scan(); lineNumber++ {
		line := scanner.Text()
		fields := strings.Fields(line)

		// Parse id
		id, _ := strconv.Atoi(fields[0])

		// Check if the current line corresponds to the specified person ID
		if id == target {
			// Parse age
			age, _ := strconv.Atoi(fields[2])

			// Parse subordinate IDs if available
			var subordinateIDs []int
			if len(fields) > 3 {
				subordinateIDStrings := strings.Split(fields[3], ",")
				for _, sid := range subordinateIDStrings {
					sidInt, _ := strconv.Atoi(sid)
					subordinateIDs = append(subordinateIDs, sidInt)
				}
			}

			return &RealEmployee{
				Id:             id,
				Name:           fields[1],
				Age:            age,
				SubordinateIds: subordinateIDs,
			}, nil
		}
	}

	// If the specified person ID is not found
	return nil, fmt.Errorf("person with ID %d not found", target)
}
