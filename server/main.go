package main

import (
	"fmt"
	"net/http"
	"strconv"

	"htmxQuizApp/server/types"
	"htmxQuizApp/server/web"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var session types.Session

/* func quizSessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.URL)

		answer, _ := strconv.Atoi(c.Params.ByName("number"))
		fmt.Printf("%d\n", answer)
		if answer == 1 {
			fmt.Printf("Correct\n")
		}

		c.Set("Session", session)
		c.Next()
	}
} */

func getQuizzes(c *gin.Context) {
	c.SetCookie("cookie", "value", 3600, "/", "", false, true)
	c.String(http.StatusOK, web.GetAllQuizzes())
}

func answerQuestion(c *gin.Context) {
	answerId, _ := strconv.Atoi(c.Params.ByName("number"))

	res := web.HandleQuizAnswer(answerId, &session)

	if session.QuizComplete {
		c.SetCookie("quiz1Res", fmt.Sprintf("%f", session.QuizRes), 3600, "/", "", false, true)
	}

	c.String(http.StatusOK, res)
}

func startQuiz(c *gin.Context) {
	quizId, _ := strconv.Atoi(c.Params.ByName("number"))

	session = web.StartQuizSession(quizId)
	c.String(http.StatusOK, web.BuildQuestion(session.Question_id))
}

func main() {
	r := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:8000"},                   // Allow all origins or specify your domain
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowed HTTP methods
		AllowHeaders:     []string{"Content-Type", "Authorization", "HX-Request", "HX-Trigger", "HX-Trigger-Name", "HX-Target", "HX-Prompt", "HX-Location", "HX-Current-Url"},
		ExposeHeaders:    []string{"HX-Trigger", "HX-Redirect"}, // Expose any htmx headers if needed
		AllowCredentials: true,                                  // Allow credentials (cookies, etc.)
	}

	r.Use(cors.New(config))

	r.GET("/quizzes", getQuizzes)
	r.GET("/quiz/:number", startQuiz)
	r.GET("/answer/:number", answerQuestion)
	r.Run(":3000")
}
