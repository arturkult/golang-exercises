package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.Open("./problems.csv")
	check(err)

	scanner := bufio.NewReader(dat)
	reader := bufio.NewScanner(os.Stdin)
	correctAnswersCount := 0
	allAnswers := 0

	data := csv.NewReader(scanner)
	data.FieldsPerRecord = 2

	for {
		line, readErr := data.Read()
		if line == nil && readErr == io.EOF {
			break
		}
		question := line[0]
		correctAnswer := line[1]

		fmt.Println(question)

		ch := make(chan int)
		var answer string
		go func() {
			reader.Scan()
			answer = reader.Text()
			ch <- 1
		}()

		select {
		case <-ch:
			if strings.EqualFold(answer, correctAnswer) {
				correctAnswersCount += 1
			}
		case <-time.After(30 * time.Second):
			fmt.Println("Time out, next question")
		}
		allAnswers += 1
	}

	fmt.Printf("Number of correct answers: %d/%d", correctAnswersCount, allAnswers)
}
