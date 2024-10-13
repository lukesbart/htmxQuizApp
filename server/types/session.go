package types

type Session struct {
	Quiz_id             int
	Question_id         int
	Questions_attempted int
	Questions_correct   int
	QuizRes             float64
	QuizComplete        bool
}
