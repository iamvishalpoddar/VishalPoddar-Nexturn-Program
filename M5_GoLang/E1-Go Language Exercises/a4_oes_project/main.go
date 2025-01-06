package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Question struct {
	Question string
	Options  [4]string
	Answer   int
}

var questionBank = []Question{
	{
		Question: "What is the capital of France?",
		Options:  [4]string{"Berlin", "Madrid", "Paris", "Rome"},
		Answer:   3,
	},
	{
		Question: "Which programming language is known as Go?",
		Options:  [4]string{"Python", "Java", "Golang", "Ruby"},
		Answer:   3,
	},
	{
		Question: "What is 5 + 7?",
		Options:  [4]string{"10", "11", "12", "13"},
		Answer:   3,
	},
}

func TakeQuiz() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	score := 0
	questionTimer := 100 * time.Second

	for i, q := range questionBank {
		fmt.Printf("\nQuestion %d: %s\n", i+1, q.Question)
		for j, option := range q.Options {
			fmt.Printf("%d. %s\n", j+1, option)
		}
		fmt.Println("Enter your answer (1-4) or type 'exit' to quit:")

		answerChan := make(chan string)

		go func() {
			if scanner.Scan() {
				answerChan <- scanner.Text()
			}
		}()

		select {
		case input := <-answerChan:
			if strings.ToLower(input) == "exit" {
				return score, nil
			}

			answer, err := strconv.Atoi(input)
			if err != nil || answer < 1 || answer > 4 {
				fmt.Println("Invalid input. Please enter a number between 1 and 4.")
				continue
			}

			if answer == q.Answer {
				fmt.Println("Correct!")
				score++
			} else {
				fmt.Println("Wrong answer.")
			}

		case <-time.After(questionTimer):
			fmt.Println("Time's up for this question.")
		}
	}

	return score, nil
}

func DisplayPerformance(score int) {
	totalQuestions := len(questionBank)
	fmt.Printf("\nYour final score: %d/%d\n", score, totalQuestions)

	percentage := float64(score) / float64(totalQuestions) * 100
	switch {
	case percentage >= 80:
		fmt.Println("Performance: Excellent")
	case percentage >= 50:
		fmt.Println("Performance: Good")
	default:
		fmt.Println("Performance: Needs Improvement")
	}
}

func main() {
	fmt.Println("Welcome to the Online Examination System")
	score, err := TakeQuiz()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	DisplayPerformance(score)
}
