package web

import (
	"fmt"
	"htmxQuizApp/server/models"
	"htmxQuizApp/server/repo"
	"htmxQuizApp/server/types"
	"strings"
)

func StartQuizSession(quizId int) types.Session {
	quizQuestions := repo.GetQuizQuestions(quizId)
	session := types.Session{
		Quiz_id:           quizId,
		Question_id:       quizQuestions[0].Id,
		Questions_correct: 0,
	}
	return session
}

func GetAllQuizzes() string {
	contentStringBuilder := []string{"<div><ul>"}

	quizzes := repo.GetAllQuizzes()

	for i, quiz := range quizzes {
		quizHTML := fmt.Sprintf(`<li><a href="#" hx-get="http://localhost:3000/quiz/%d" hx-target="#quiz" hx-swap="innerHTML">%d: %s</a></li>`, quiz.Id, i+1, quiz.Name)

		contentStringBuilder = append(contentStringBuilder, quizHTML)
	}

	contentStringBuilder = append(contentStringBuilder, "</div></ul>")

	return strings.Join(contentStringBuilder, "")
}

func HandleQuizAnswer(answerId int, session *types.Session) string {
	correct := repo.CheckQuestionAnswer(answerId)
	if correct {
		session.Questions_correct++
	}

	questions := repo.GetQuizQuestions(session.Quiz_id)
	if session.Question_id < questions[len(questions)-1].Id {
		nextQuestion := BuildQuestion(session.Question_id + 1)
		session.Questions_attempted++
		session.Question_id++
		return nextQuestion
	}

	session.QuizComplete = true
	session.QuizRes = float64(session.Questions_correct / session.Questions_attempted)
	return fmt.Sprintf(`<h2>Results %d/%d</h2><a hx-get="http://localhost:3000/quizzes" hx-target="#quiz" hx-swap="innerHTML"><button>Back to Quizzes</button></a>`, session.Questions_correct, len(questions))
}

func BuildQuestion(questionId int) string {
	contentStringBuilder := []string{"<div>"}

	contentStringBuilder = append(contentStringBuilder, fmt.Sprintf("<h2>%s</h2>", getQuestionTitle(questionId)))

	contentStringBuilder = append(contentStringBuilder, getQuestionOptions(questionId)...)

	contentStringBuilder = append(contentStringBuilder, "</div>")

	return strings.Join(contentStringBuilder, "")
}

func getQuestionTitle(questionId int) string {
	question := models.Quiz_question{}
	row := repo.QueryDBForRow("SELECT question FROM Quiz_Question WHERE id = ?", questionId)
	row.Scan(&question.Question)
	return question.Question
}

func getQuestionOptions(quizQuestion int) []string {
	options := repo.GetQuizQuestionOptions(quizQuestion)
	optionsHTML := []string{}
	for i, option := range options {
		optionHTML := fmt.Sprintf(`<p><a hx-get="http://localhost:3000/answer/%d" hx-target="#quiz" hx-swap="innerHTML"><button>%d: %s</button></a></p>`, option.Id, i+1, option.Option)
		optionsHTML = append(optionsHTML, optionHTML)
	}
	return optionsHTML
}
