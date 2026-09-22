package handlers

import (
	"uuid"

	"github.com/mohmmd00/Pulsar/dispatcher/internal/service"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateNotificationHandler(ctx *gin.Context) {
	var req CreateNotificationRequest
	bindErr := ctx.ShouldBindJSON(&req)
	if bindErr != nil {
		ctx.JSON(400, gin.H{"error": bindErr.Error()})
		return
	}

	userIDValue, exists := ctx.Get("user_id") //ctx doesnt care of variables so you have to be sure that the value you are getting is uuid !

	if !exists {
		ctx.JSON(500, gin.H{"error": "Unauthorized !"})
		return
	}
	userID, ok := userIDValue.(uuid.UUID) //type assertions
	if !ok {
		ctx.JSON(400, gin.H{"error": "user_id has unexpected type !"})
		return
	}

	//service layer
	newNote, crtNoteErr := service.SubmitNotification(userID, req.Channel, req.Recipient, req.Message, s.DatabasePool, ctx.Request.Context())
	if crtNoteErr != nil {
		ctx.JSON(400, gin.H{"error": crtNoteErr.Error()})
		return
	}

	ctx.JSON(201, gin.H{"message": "Operation successful", "notification": newNote})

}
