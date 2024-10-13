package models

type QuizQuestionOption struct {
	Id             int
	QuizQuestionId int
	Option         string
	Correct        bool
}
