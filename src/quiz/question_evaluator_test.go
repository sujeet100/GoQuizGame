package quiz

import "testing"

func TestEvaluateQuestion(t *testing.T) {

	q := QuizQuestion{"1+2", "3"}
	result := Evaluate(q, "3")
	if result != true {
		t.Errorf("Provided answer: %s, does not match with expected answer: %s", "3", q.ExpectedAnswer)
	}

}

func TestEvaluateQuestionWrongAnswer(t *testing.T) {

	q := QuizQuestion{"1+2", "3"}
	result := Evaluate(q, "5")
	if result != false {
		t.Errorf("Provided answer: %s, should not match with expected answer: %s", "5", q.ExpectedAnswer)
	}

}
