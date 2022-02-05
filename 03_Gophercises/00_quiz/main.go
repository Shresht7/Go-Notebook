package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Score struct {
	correct int8
	total   int8
}

func main() {
	//	Instantiate scorecard
	scorecard := Score{}

	//	Read CSV file
	data := readCSVFile()

	//	Play Quiz
	for _, v := range data {

		//	Prompt question
		fmt.Print(v[0], ": ")

		//	Take input
		var res string
		fmt.Scanln(&res)

		//	Parse answer and user response as int
		response, errRes := strconv.Atoi(res)
		answer, errAns := strconv.Atoi(v[1])

		//	If no errors while conversion
		if errRes == nil && errAns == nil {
			//	...and user response matches answer
			if response == answer {
				scorecard.correct++ //	Increment score
			}
		}

		//	Increment total questions attempted
		scorecard.total++
	}

	//	Show scorecard
	fmt.Println("\nScored ", scorecard.correct, " out of ", scorecard.total)
}

func readCSVFile() [][]string {
	//	Read CSV file
	file, err := os.ReadFile("./problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	//	Parse CSV
	csvReader := csv.NewReader(strings.NewReader(string(file)))

	var data [][]string
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, record)
	}

	return data
}