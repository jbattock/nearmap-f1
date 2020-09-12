package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	const layoutISO = "2006-01-02"

	var y []int
	var yhat []int

	// open file
	file, err := os.Open(os.Args[1])
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully Opened file")

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		// ignore comments
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "dates") {
			continue
		}

		split := strings.Split(line, "|")

		// check if date is a Thursday
		date, _ := time.Parse(layoutISO, split[0])

		if date.Weekday() == 4 {
			expected, _ := strconv.Atoi(split[1])
			split2 := strings.TrimSuffix(split[2], "\n")
			predicted, _ := strconv.Atoi(split2)
			y = append(y, expected)
			yhat = append(yhat, predicted)
		}
	}

	identified := 0
	total := 0
	realPositive := 0
	for i := 0; i < len(y); i++ {
		if yhat[i] == 1 {
			if y[i] == 1 {
				identified++
			}
			total++
		}
		if y[i] == 1 {
			realPositive++
		}
	}

	percision := float64(identified) / float64(total)
	recall := float64(identified) / float64(realPositive)
	// calculate f1
	f1 := 2 * (percision * recall) / (percision + recall)
	fmt.Println("f1 score: ", f1)
}
