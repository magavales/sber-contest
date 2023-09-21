package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"todo-list/pkg/database"
	"todo-list/pkg/model"
)

func (h *Handler) createTask(ctx *gin.Context) {
	var (
		db   database.Database
		task model.Task
		err  error
	)
	err = task.DecodeJSON(ctx.Request.Body)
	if err != nil {
		return
	}

	db.Connect()
	err = db.Access.CreateTask(db.Pool, task)
	if err != nil {
		return
	}
}

func (h *Handler) updateTask(ctx *gin.Context) {
	var (
		db   database.Database
		task model.Task
		err  error
	)
	err = task.DecodeJSON(ctx.Request.Body)
	if err != nil {
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	task.ID = int64(id)

	db.Connect()
	err = db.Access.UpdateTask(db.Pool, task)
	if err != nil {
		return
	}
}