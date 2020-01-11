package quiz

func Evaluate(question QuizQuestion, answer string) bool {
	return answer == question.ExpectedAnswer
}
