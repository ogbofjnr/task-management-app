package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/ogbofjnr/maze/cronjobs"
	"github.com/ogbofjnr/maze/database/repositories"
	"github.com/ogbofjnr/maze/handlers"
	"github.com/ogbofjnr/maze/pkg/config_manager"
	"github.com/ogbofjnr/maze/pkg/cron"
	databasa "github.com/ogbofjnr/maze/pkg/db"
	logger2 "github.com/ogbofjnr/maze/pkg/logger"
	mailer2 "github.com/ogbofjnr/maze/pkg/mailer"
	n "github.com/ogbofjnr/maze/pkg/notificator"
	"github.com/ogbofjnr/maze/pkg/validator"
	"github.com/ogbofjnr/maze/routes"
	"github.com/ogbofjnr/maze/services"
	"github.com/sendgrid/sendgrid-go"
	"log"
	"net/http"
	"time"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	appConfig := config_manager.GetConfig("app.conf")

	logger := logger2.InitLogger(appConfig.GetString("app.logLevel"))
	defer logger.Sync()

	db := databasa.InitDB()
	defer db.Close()

	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository, db)
	appValidator := validator.NewValidator()
	userHandler := handlers.NewUserHandler(appValidator, userService, logger)

	taskRepository := repositories.NewTaskRepository()
	taskService := services.NewTaskService(taskRepository, db)
	taskHandler := handlers.NewTaskHandler(appValidator, taskService, userService, logger)

	mailClient := sendgrid.NewSendClient(appConfig.GetString("mail.sendgridApiKey"))
	mailer := mailer2.NewMailer(mailClient, logger)

	notificator := n.NewNotificator(logger)
	taskReminder := cronjobs.NewTaskReminder(db, userRepository, taskRepository, logger, notificator, mailer)
	cron.InitCron(taskReminder.Run)
	go notificator.Run()
	notificationHandler := handlers.NewNotificationHandler(logger, notificator)

	r := routes.InitRouter(userHandler, taskHandler)
	r.HandleFunc("/ws", notificationHandler.Connect)

	http.Handle("/", r)
	srv := &http.Server{
		Handler:      r,
		Addr:         appConfig.GetString("server.port"),
		WriteTimeout: appConfig.GetDuration("server.writeTimeout") * time.Second,
		ReadTimeout:  appConfig.GetDuration("server.readTimeout") * time.Second,
	}

	fmt.Printf("start server at %s", appConfig.GetString("server.port"))
	log.Fatal(srv.ListenAndServe())
}
