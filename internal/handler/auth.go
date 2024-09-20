package handler

import (
	"github.com/VeeRomanoff/todo-app/domain"
	"github.com/VeeRomanoff/todo-app/internal/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var reqBody domain.User

	// куда мы хотим распарсить тело запроса
	if err := ctx.BindJSON(&reqBody); err != nil {
		helpers.NewResponseMessage(ctx, http.StatusBadRequest, err.Error())
	}
}

func (h *Handler) signIn(ctx *gin.Context) {

}
