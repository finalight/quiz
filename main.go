package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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

func ask(question *Question, c chan int) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("What is the answer for %s?: ", question.QuestionText)
	input, _ := reader.ReadString('\n')

	answer, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Why have error?")
		c <- 0
		return
	}

	if answer == question.Answer {
		fmt.Println("You are correct")
		c <- 1
	} else {
		fmt.Printf("Wrong answer, the correct answer is %v\n", question.Answer)
		c <- 0
	}

}

func main() {
	quizOverall, err := readFromCSV("problems.csv")

	if err != nil {
		fmt.Println("Error in reading file")
		return
	}

	deadline := time.Second * 5
	c := make(chan int, 1)

	for _, question := range quizOverall.Questions {
		ctx, cancel := context.WithTimeout(context.Background(), deadline)
		go func() {
			ask(&question, c)
		}()

		defer cancel()
		select {
		case <-ctx.Done():
			fmt.Println()
			fmt.Println("Timeout, next question")
		case result := <-c:
			quizOverall.score += result
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
