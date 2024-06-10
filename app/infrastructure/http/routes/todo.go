package routes

import (
	"app/interface/controllers"

	"github.com/gin-gonic/gin"
)

func ToDoRoutes(r *gin.Engine, tc *controllers.ToDoController) {
	r.GET("/todos", func(c *gin.Context) {
		tc.GetAll(c)
	})

	r.GET("/todos/:id", func(c *gin.Context) {
		tc.GetByID(c)
	})

	r.POST("/todos", func(c *gin.Context) {
		tc.Create(c)
	})

	r.PUT("/todos/:id", func(c *gin.Context) {
		tc.Update(c)
	})

	r.DELETE("/todos/:id", func(c *gin.Context) {
		tc.Delete(c)
	})
}
