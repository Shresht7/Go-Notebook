package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"time"
)

func main() {
	//	Seed the Random Number Generator to obtain random values
	rand.Seed(time.Now().UTC().UnixNano())

	//	Get Arguments
	arguments := parseArguments()

	//	Match dice argument with dX format
	matched, _ := regexp.Match("d(\\d+)", []byte(*arguments.dice))
	if matched {

		//	Roll the die
		rolls := rollDice(arguments.dice, arguments.numRoll)
		printDice(rolls)

		//	Calculate Sum
		if *arguments.sum == true {
			diceSum := sumDice(rolls)
			fmt.Printf("The sum of the dice was %d\n", diceSum)
		}

		//	Roll with Advantage
		if *arguments.advantage == true {
			roll := rollWithAdvantage(rolls)
			fmt.Printf("The roll with advantage was %d\n", roll)
		}

		//	Roll with Disadvantage
		if *arguments.disadvantage == true {
			roll := rollWithDisadvantage(rolls)
			fmt.Printf("The roll with disadvantage was %d\n", roll)

		}

		//	Error matching dX
	} else {
		log.Fatal("Improper format for dice. dX where X is an integer")
	}
}

//	---------------
//	PARSE ARGUMENTS
//	---------------

type Arguments struct {
	dice         *string
	numRoll      *int
	sum          *bool
	advantage    *bool
	disadvantage *bool
}

//	Parse Command Line Arguments and return the flags
func parseArguments() Arguments {

	dice := flag.String("d", "d6", "The type of dice to roll. Format: dX where X is an integer. Default: d6")
	numRoll := flag.Int("n", 1, "The number of die to roll. Default 1")
	sum := flag.Bool("s", false, "Get the sum of all the dice rolls")
	advantage := flag.Bool("adv", false, "Roll the dice with advantage")
	disadvantage := flag.Bool("disadv", false, "Roll the dice with disadvantage")

	flag.Parse()

	return Arguments{dice, numRoll, sum, advantage, disadvantage}
}

//	------------
//	DICE HELPERS
//	------------

func rollDice(dice *string, times *int) []int {
	var rolls []int
	diceSides := (*dice)[1:]
	d, err := strconv.Atoi(diceSides)
	if err != nil {
		fmt.Println(err)
		return rolls
	}

	for i := 0; i < *times; i++ {
		rolls = append(rolls, rand.Intn(d)+1)
	}
	return rolls
}

func printDice(rolls []int) {
	for i, dice := range rolls {
		fmt.Printf("Roll %d was %d\n", i+1, dice)
	}
}

func sumDice(rolls []int) int {
	sum := 0
	for _, dice := range rolls {
		sum += dice
	}
	return sum
}

func rollWithAdvantage(rolls []int) int {
	sort.Ints(rolls)
	return rolls[len(rolls)-1]
}

func rollWithDisadvantage(rolls []int) int {
	sort.Ints(rolls)
	return rolls[0]
}
