package leetcode

import (
	"encoding/csv"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

func TestMaxArea1(t *testing.T) {
	assert.Equal(t,
		49,
		maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}),
	)
}

func TestMaxArea2(t *testing.T) {
	assert.Equal(t,
		1,
		maxArea([]int{1, 1}),
	)
}

func TestMaxArea3(t *testing.T) {
	ints, err := readIntsFromFile()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t,
		705634720,
		maxArea(ints),
	)
}

func readIntsFromFile() ([]int, error) {
	file, err := os.Open("./ints.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	intsFromFile, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	var ints []int
	for _, intFromFile := range intsFromFile {
		var num int
		num, err = strconv.Atoi(intFromFile[0])
		if err != nil {
			return nil, err
		}
		ints = append(ints, num)
	}
	return ints, nil
}
