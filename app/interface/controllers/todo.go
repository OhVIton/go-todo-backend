package controllers

import (
	"app/domain/model"
	"app/usecase/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ToDoController struct {
	service *services.ToDoService
}

func NewToDoController(service *services.ToDoService) *ToDoController {
	return &ToDoController{service: service}
}

func (tc *ToDoController) GetAll(c *gin.Context) {
	var todos []*model.Todo
	todos, err := tc.service.GetAll(c)

	if err != nil {
		log.Fatalf("Error todos: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (tc *ToDoController) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var todo *model.Todo
	todo, err := tc.service.GetByID(c, id)

	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record not found"})
		return
	}

	if err != nil {
		log.Fatalf("Error todos: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (tc *ToDoController) Create(c *gin.Context) {
	var todo model.Todo
	c.BindJSON(&todo)
	err := tc.service.Create(c, todo.Content)

	if err != nil {
		log.Fatalf("Error todos: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Data(http.StatusCreated, "application/json", []byte(""))
}

func (tc *ToDoController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var todo *model.Todo
	todo, err := tc.service.GetByID(c, id)

	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record not found"})
		return
	}

	if err != nil {
		log.Fatalf("Error todos: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = tc.service.Update(c.Request.Context(), todo)

	if err != nil {
		log.Fatalf("Error todos: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Data(http.StatusOK, "application/json", []byte(""))
}

func (tc *ToDoController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := tc.service.Delete(c, id)

	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record not found"})
		return
	}

	if err != nil {
		log.Fatalf("Error todos: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Data(http.StatusNoContent, "application/json", []byte(""))
}
