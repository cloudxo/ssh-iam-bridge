package strarray

import (
	"bufio"
	"io"
	"os"
	"sort"
)

func Contains(set []string, test string) bool {
	for _, member := range set {
		if test == member {
			return true
		}
	}

	return false
}

func sum(values []int) (total int) {
	for _, v := range values {
		total += v
	}
	return
}

func ContainsAll(set []string, tests []string) bool {
	total := len(tests)
	found := make([]int, total)

	for _, s := range set {
		for idx, check := range tests {
			if s == check {
				found[idx] = 1
				if sum(found) == total {
					return true
				}
			}
		}
	}

	return false
}

// Compares array1 against a2 and returns the values in array1 that are not present in a2.
func Diff(tests, set []string) []string {

	missing := make([]string, 0)

	for _, test := range tests {
		if Contains(set, test) {
			continue
		}

		missing = append(missing, test)
	}

	return missing
}

func Unique(list []string) []string {

	sort.Strings(list)

	result := []string{}

	var last string

	for _, s := range list {
		if s == last {
			continue
		}
		last = s
		result = append(result, s)
	}

	return result
}

func Filter(ss []string, filterfn func(s string) bool) []string {

	result := []string{}

	for _, s := range ss {
		if !filterfn(s) {
			continue
		}
		result = append(result, s)
	}

	return result
}

func WriteFile(filename string, stringList ...[]string) error {

	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	for _, set := range stringList {
		for _, line := range set {
			if _, err := file.WriteString(line); err != nil {
				return err
			}
		}
	}

	return nil
}

func ReadFile(filename string) ([]string, error) {

	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	var line string
	var lines []string

	for {
		if line, err = reader.ReadString('\n'); err != nil {
			break
		}

		lines = append(lines, line)
	}

	if err != io.EOF {
		return nil, err
	}

	return lines, nil
}
