package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

//Quiz is a struct to hold questions and answers
type Quiz struct {
	questions string
	answer    string
}


//readFile read the csv file, and return in to the main
func readFile(filename string) ([][]string, error) {
	//Open opens the named file for reading. If successful, methods on
	//the returned file can be used for reading
	csvFile, err := os.Open(filename)

	if err != nil {
		return [][]string{}, err
	}

	defer csvFile.Close()

	// NewReader returns a new Reader that reads from csvFile.
	file := csv.NewReader(csvFile)

	if _, err := file.Read(); err != nil {
		return [][]string{}, err
	}

	//ReadAll() eads all the remaining records from the reader. 
	//Each record is a slice of fields.
	records, err := file.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

func main() {
	var allQuiz []Quiz
	var quiz Quiz

	//get the content file in readFile()
	records, err := readFile("quiz_problems.csv")

	if err != nil {
		log.Fatal(err)
	}

	//assign all quiz_problems.csv to slice of quiz in variable allQuiz
	for _, record := range records {
		quiz.questions = record[0]
		quiz.answer = record[1]

		allQuiz = append(allQuiz, quiz)
	}

	for _, item := range allQuiz {
		fmt.Printf("%s = %s\n", item.questions, item.answer)
	}

}
