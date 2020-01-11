package quiz

import (
	"testing"
	"reflect"
)

func TestReadQuiz(t *testing.T) {

	quiz := ReadQuiz("testQuiz.csv")

	expectedQuiz := []QuizQuestion{QuizQuestion{"1+2", "3"}, QuizQuestion{"4+4 is what, sir?", "8"}}

	if !reflect.DeepEqual(quiz, expectedQuiz) {
		t.Errorf("Quiz questions do not match expcted: %s, actual: %s", expectedQuiz, quiz)
	}

}
