package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
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

//exitProgram exit this program when errors occured
func exitProgram(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

//startQuiz start the program with user input
func startQuiz(quiz []Quiz) {
	//timerFlag is helper cli for timer
	timerFlag := flag.Int("time", 30, "The timer for answering this questions")	

	//flag.Parse() makes user can input the time in cli (-time={our_time})
	flag.Parse()

	//timer is the timer for this quiz
	timer := time.NewTimer(time.Duration(*timerFlag) * time.Second)

	var input string
	var correctAnswer = 0
	inputChannel := make(chan string)

	//goroutine for input answer
	//this anonymous goroutine prevent us from scanf bug with timer
	go func ()  {
		fmt.Scanf("%s\n", &input)
		inputChannel <- input
	}()

	//for loop until time runs out
	for _, item := range quiz {
		fmt.Printf("%s = ", item.questions)

		select {
		//if timer expires, timer will sent to C
		case <-timer.C:
			fmt.Printf("\nCorrect answer = %d", correctAnswer)
			return
		
		//if in goroutine user has input some answers, do this
		case input := <- inputChannel:
			if input == item.answer {
				correctAnswer++
			}
		}
	}

	
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

	startQuiz(allQuiz)
	//user input section and correctAnswer section

}
