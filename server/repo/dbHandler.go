package repo

import (
	"database/sql"
	"fmt"
	"htmxQuizApp/server/models"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const file string = "repo/quiz.db"

func QueryDBForRow(query string, args ...any) *sql.Row {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}

	return db.QueryRow(query, args...)
}

func QueryDBForRowsMultiple(query string, args ...any) (*sql.Rows, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}

	return db.Query(query, args...)
}

func GetAllQuizzes() []models.Quiz {
	rows, err := QueryDBForRowsMultiple("SELECT * FROM Quiz")

	if err != nil {
		fmt.Println(err)
		return []models.Quiz{}
	}

	quizzes := []models.Quiz{}
	for rows.Next() {
		quiz := models.Quiz{}
		rows.Scan(&quiz.Id, &quiz.Name)
		quizzes = append(quizzes, quiz)
	}

	return quizzes
}

func GetQuizQuestions(quizId int) []models.Quiz_question {
	rows, err := QueryDBForRowsMultiple("SELECT * FROM Quiz_Question WHERE quiz_id = ?", quizId)

	if err != nil {
		fmt.Println(err)
		return []models.Quiz_question{}
	}

	questions := []models.Quiz_question{}
	for rows.Next() {
		question := models.Quiz_question{}
		rows.Scan(&question.Id, &question.Quiz_id, &question.Question)
		questions = append(questions, question)
	}
	return questions
}

func GetQuizQuestionOptions(quizQuestionId int) []models.QuizQuestionOption {
	rows, err := QueryDBForRowsMultiple("SELECT id, quiz_question_id, option, correct FROM Quiz_Question_Option WHERE quiz_question_id = ?", quizQuestionId)

	if err != nil {
		fmt.Println(err)
		return []models.QuizQuestionOption{}
	}

	options := []models.QuizQuestionOption{}
	for rows.Next() {
		option := models.QuizQuestionOption{}
		rows.Scan(&option.Id, &option.QuizQuestionId, &option.Option, &option.Correct)
		options = append(options, option)
	}
	return options
}

func CheckQuestionAnswer(optionId int) bool {
	row := QueryDBForRow("SELECT correct FROM Quiz_Question_Option WHERE id = ?", optionId)
	option := models.QuizQuestionOption{}
	fmt.Printf("OPTId: %d\n", optionId)
	row.Scan(&option.Correct)
	fmt.Println(option.Correct)
	return option.Correct
}
