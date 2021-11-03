package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

type scores struct {
	correct    int
	wrong      int
	totalQuest int
}

func readCsvFile(filepath string) [][]string {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Unable to read input file "+filepath+" ", err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filepath, err)
	}

	return records
}

func main() {
	filename := flag.String("filename", "problems.csv", "Questions and answers are inside csv file.")
	flag.Parse()
	records := readCsvFile(*filename)
	score := scores{correct: 0, wrong: 0, totalQuest: len(records)}
	for _, array := range records {
		fmt.Printf("%s: ", array[0])
		reader := bufio.NewReader(os.Stdin)
		answer, _ := reader.ReadString('\n')
		answer = answer[:len(answer)-2] // WHY ???
		if array[1] == answer {
			score.correct += 1
		} else {
			score.wrong += 1
		}
	}
	fmt.Println("Total Score")
	fmt.Println("---------------")
	fmt.Println("Total Question: ", score.totalQuest)
	fmt.Println("Total correct answers: ", score.correct)
	fmt.Println("Total wrong answers: ", score.wrong)
}
