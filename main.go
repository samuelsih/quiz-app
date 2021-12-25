package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

//Quiz is a struct to hold questions and answers
type Quiz struct {
	questions string
	answer    string
}


//readFile read the csv file, and return in to the main
func readFile(filename string) ([][]string) {
	//Open opens the named file for reading. If successful, methods on
	//the returned file can be used for reading
	csvFile, err := os.Open(filename)

	if err != nil {
		exitProgram(fmt.Sprintf("Cannot read file %s", filename))
	}

	defer csvFile.Close()

	// NewReader returns a new Reader that reads from csvFile.
	file := csv.NewReader(csvFile)

	if _, err := file.Read(); err != nil {
		exitProgram(err.Error())
	}

	//ReadAll() eads all the remaining records from the reader. 
	//Each record is a slice of fields.
	records, err := file.ReadAll()

	if err != nil {
		exitProgram(err.Error())
	}

	return records
}

func exitProgram(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	var allQuiz []Quiz
	var quiz Quiz

	//get the content file in readFile()
	records := readFile("quiz_problems.csv")


	//assign all quiz_problems.csv to slice of quiz in variable allQuiz
	for _, record := range records {
		quiz.questions = record[0]
		quiz.answer = strings.TrimSpace(record[1])

		allQuiz = append(allQuiz, quiz)
	}

	//user input section and correctAnswer section
	var input string
	var correctAnswer = 0

	for _, item := range allQuiz {
		fmt.Printf("%s = ", item.questions)
		fmt.Scanf("%s\n", &input)

		if input == item.answer {
			correctAnswer++
		}
	}

	fmt.Println("Correct answer = " + strconv.Itoa(correctAnswer))
}
