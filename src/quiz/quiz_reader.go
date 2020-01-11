package quiz

import (
	"os"
	"encoding/csv"
	"io"
	"strings"
	)

func ReadQuiz(filePath string) []QuizQuestion {
	var quizQuestions []QuizQuestion
	file, _ := os.Open(filePath)
	r := csv.NewReader(file)
	for {
		line, e := r.Read()
		if e == io.EOF {
			break
		} else {
			quizQuestions = append(
				quizQuestions,
				QuizQuestion{
					strings.TrimSpace(line[0]),
					strings.TrimSpace(line[1])})
		}

	}
	return quizQuestions
}
