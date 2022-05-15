package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	//	Get fileName
	fileName := flag.String("fileName", "./problems.csv", "Problem Set")
	timeLimit := flag.Int("time-limit", 20, "The Time Limit")
	shuffle := flag.Bool("shuffle", true, "Shuffle the problem set")
	flag.Parse()

	//	Instantiate scorecard
	correct := 0

	//	Read CSV file
	data := readCSVFile(*fileName, *shuffle)

	//	Create timer channel
	timer := time.NewTimer(time.Duration(*timeLimit * int(time.Second)))

	//	Play Quiz
	for _, v := range data {

		//	Prompt question
		fmt.Print(v[0], ": ")

		ansCh := make(chan string)
		go func() {
			//	Take input
			var res string
			fmt.Scanln(&res)
			ansCh <- res
		}()

		select {
		case <-timer.C:
			fmt.Println("\nScored ", correct, " out of ", len(data))
			return
		case res := <-ansCh:
			//	Parse answer and user response as int
			response, errRes := strconv.Atoi(res)
			answer, errAns := strconv.Atoi(v[1])

			//	If no errors while conversion
			if errRes == nil && errAns == nil {
				//	...and user response matches answer
				if response == answer {
					correct++ //	Increment score
				}
			}
		}
	}

	//	Show scorecard
	fmt.Println("\nScored ", correct, " out of ", len(data))
}

func readCSVFile(fileName string, shuffle bool) [][]string {
	//	Read CSV file
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Failed to open csv file: ", fileName)
		os.Exit(1)
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

	if shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })
	}

	return data
}
