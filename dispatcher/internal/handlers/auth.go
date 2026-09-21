package handlers

import (
	"github.com/mohmmd00/Pulsar/dispatcher/internal/service"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterHandler(ctx *gin.Context) {

	var req RegisterRequest
	bindErr := ctx.ShouldBindJSON(&req)
	if bindErr != nil {
		ctx.JSON(400, gin.H{"error": bindErr.Error()})
		return
	}

	//service layer then repository layer
	newUser, regUsrErr := service.RegisterUser(req.Email, req.Password, s.DatabasePool, ctx.Request.Context())

	if regUsrErr != nil {
		ctx.JSON(500, gin.H{"error": regUsrErr.Error()})
		return
	}

	ctx.JSON(201, gin.H{"message": "operationSuccessful", "id": newUser.ID})

}


