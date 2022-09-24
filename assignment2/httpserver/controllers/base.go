package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rifqoi/hacktiv8-assignments/assignment2/httpserver/controllers/views"
)

func WriteJSONResponse(ctx *gin.Context, res *views.Response) {
	ctx.JSON(res.Status, &res)
}
