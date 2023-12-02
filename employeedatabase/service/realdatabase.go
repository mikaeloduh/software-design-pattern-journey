package service

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RealDatabase struct {
	filename string
	index    map[int64]int64
}

func NewRealDatabase(filename string) *RealDatabase {
	return &RealDatabase{filename: filename}
}

func (d *RealDatabase) Init() error {
	index, err := d.buildIndex()
	if err != nil {
		return fmt.Errorf("error building index: %s", err)
	}

	fmt.Println("INDEX:", index)

	d.index = index

	return nil
}

func (d *RealDatabase) GetEmployeeById(targetId int) (*RealEmployee, error) {
	lineNumber := int64(targetId)
	line, err := d.readSpecificLine(lineNumber)
	if err != nil {
		return nil, fmt.Errorf("readSpecificLine err: %s", err)
	}

	// Process each line until finding the specified person ID
	fields := strings.Fields(line)

	// Parse fields
	id, _ := strconv.Atoi(fields[0])
	age, _ := strconv.Atoi(fields[2])
	name := fields[1]

	// Parse subordinate IDs if available
	var subordinateIDs []int
	if len(fields) > 3 {
		subordinateIDStrings := strings.Split(fields[3], ",")
		for _, sid := range subordinateIDStrings {
			sidInt, _ := strconv.Atoi(sid)
			subordinateIDs = append(subordinateIDs, sidInt)
		}
	}

	return NewRealEmployee(id, name, age, subordinateIDs), nil
}

func (d *RealDatabase) buildIndex() (map[int64]int64, error) {
	file, err := os.Open(d.filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	index := make(map[int64]int64)
	var offset int64
	var lineNumber int64
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		index[lineNumber] = offset
		offset += int64(len(scanner.Bytes()) + 1) // Add 1 for the newline character
		lineNumber += 1
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return index, nil
}

func (d *RealDatabase) readSpecificLine(lineNumber int64) (string, error) {
	offset, found := d.index[lineNumber]
	if !found {
		return "", fmt.Errorf("line number %d not found", lineNumber)
	}

	file, err := os.Open(d.filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	file.Seek(offset, 0)

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return scanner.Text(), nil
	}

	return "", fmt.Errorf("error reading line number %d", lineNumber)
}
