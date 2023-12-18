package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/rodrigodosanjosoliveira/golang-crud-gin/config"
	"github.com/rodrigodosanjosoliveira/golang-crud-gin/controller"
	"github.com/rodrigodosanjosoliveira/golang-crud-gin/model"
	"github.com/rodrigodosanjosoliveira/golang-crud-gin/repository"
	"github.com/rodrigodosanjosoliveira/golang-crud-gin/router"
	"github.com/rodrigodosanjosoliveira/golang-crud-gin/service"
	"net/http"

	"github.com/rodrigodosanjosoliveira/golang-crud-gin/helper"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Started Server!")
	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	// Repository
	tagsRepository := repository.NewTagsRepositoryImpl(db)

	// Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	// Controller
	tagsController := controller.NewTagsController(tagsService)

	// Router
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
