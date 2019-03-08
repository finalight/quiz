package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Question structure
type Question struct {
	QuestionText string
	Answer       int
}

// Quiz structure
type Quiz struct {
	Questions []Question
	score     int
}

func main() {
	quizOverall, err := readFromCSV("problems.csv")

	if err != nil {
		fmt.Println("Error in reading file")
		return
	}
	for index := range quizOverall.Questions {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("What is the answer for %s?: ", quizOverall.Questions[index].QuestionText)
		input, _ := reader.ReadString('\n')

		answer, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("Why have error?")
		}

		if answer == quizOverall.Questions[index].Answer {
			fmt.Println("You are correct")
			quizOverall.score++
		} else {
			fmt.Printf("Wrong answer? The correct answer is %v\n", quizOverall.Questions[index].Answer)
		}
	}

	fmt.Printf("The score is %v\n", quizOverall.score)

}

func readFromCSV(location string) (*Quiz, error) {
	quizOverall := Quiz{}
	file, err := os.Open(location)

	if err != nil {
		fmt.Println("error in opening this file")
	}

	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()

	if err != nil {
		fmt.Println("error in reading this file")
	}

	for index := range lines {
		question := Question{}
		question.QuestionText = lines[index][0]
		answer, err := strconv.Atoi(lines[index][1])
		if err != nil {
			return nil, err
		}

		question.Answer = answer
		quizOverall.Questions = append(quizOverall.Questions, question)
	}

	return &quizOverall, nil

}
