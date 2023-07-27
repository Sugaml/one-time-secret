package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) heathCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "onetime secret is healthy/Running.")
}
