package main

import (
	"app/domain/model"
	"app/infrastructure/database"
	repositories "app/infrastructure/database/repository"
	"app/infrastructure/http/routes"
	"app/interface/controllers"
	"app/usecase/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := database.GetDB()

	if err != nil {
		log.Fatalf(("Failed to connect to database: %v"), err)
	}

	err = db.AutoMigrate(&(model.Todo{}))
	if err != nil {
		log.Fatalf(("Failed to migrate database: %v"), err)
	}

	tr := repositories.NewToDoRepository(db)
	ts := services.NewToDoService(tr)
	tc := controllers.NewToDoController(ts)

	r := gin.Default()
	routes.ToDoRoutes(r, tc)

	r.Run(":8080")
}
