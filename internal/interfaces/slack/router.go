package slack

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	r := gin.Default()

	r.POST("/slack/command", func(ctx *gin.Context) {
		command := ctx.PostForm("command")
		userID := ctx.PostForm("user_id")

		var response string
		var err error

		switch command {
		case "/register":
			response, err = userHandler.HandleUserRegistration(userID)
		case "/checkin":
			response, err = attendanceHandler.HandleCheckIn(userID)
		case "/checkout":
			response, err = attendanceHandler.HandleCheckOut(userID)
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown command"})
			return
		}

	})
}
