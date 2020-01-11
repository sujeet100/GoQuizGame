package main

import (
	"fmt"
	"quiz"
	"bufio"
	"os"
	"strings"
	"time"
)

func main() {

	quizResult.Answers = make(map[quiz.QuizQuestion]bool)
	welcome()
	startTimer()
	ch := make(chan bool)
	go isTimeOver(ch)
	askQuestions(ch)
	printResult()
}

func isTimeOver(ch chan bool) {
	time.Sleep(1 * time.Second)
	isOver := timeLapsed() >= float64(quizResult.TimeLimitInSeconds)
	if isOver {
		ch <- isOver
		close(ch)
	} else {
		isTimeOver(ch)
	}
}

func askQuestions(timeOver chan bool) {
	questions := quiz.ReadQuiz("problems.csv")
	for qno, question := range questions {
		select {
		case <-timeOver:
			break;
		default:
			askQuestion(qno+1, question)
		}

	}
}

func startTimer() {
	quizResult.StartTime = time.Now()
}

func timeLapsed() float64 {
	return time.Now().Sub(quizResult.StartTime).Seconds()
}

func welcome() {
	fmt.Println("Welcome to quiz")
}

func askQuestion(qno int, question quiz.QuizQuestion) {
	fmt.Printf("Question %d: %s = ", qno, question.Question)
	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	isRight := quiz.Evaluate(question, strings.TrimSpace(answer))
	quizResult.Answers[question] = isRight
}

func printResult() {
	fmt.Printf("Total Questions: %d", len(quizResult.Answers))
	fmt.Printf("Total Right Answers: %d", quizResult.totalRightAnswers())
}

type QuizResult struct {
	Answers            map[quiz.QuizQuestion]bool
	TimeLimitInSeconds int
	StartTime          time.Time
}

var quizResult = &QuizResult{TimeLimitInSeconds: 10}

func (quizResult *QuizResult) totalRightAnswers() int {
	count := 0
	for _, v := range quizResult.Answers {
		if v {
			count += 1
		}
	}
	return count;
}
